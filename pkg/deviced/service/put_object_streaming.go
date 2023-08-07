package metathings_deviced_service

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	file_helper "github.com/nayotta/metathings/pkg/common/file"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	ss "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	pb "github.com/nayotta/metathings/proto/deviced"
)

func (self *MetathingsDevicedService) new_put_object_streaming_chunks_response(offsets []int64) *pb.PutObjectStreamingResponse {
	chunks := &pb.ObjectChunks{}
	for _, offset := range offsets {
		chunks.Chunks = append(chunks.Chunks, &pb.ObjectChunk{
			Offset: offset,
			Length: self.opt.Methods.PutObjectStreaming.ChunkSize,
		})
	}
	res := &pb.PutObjectStreamingResponse{
		Id:       id_helper.NewId(),
		Response: &pb.PutObjectStreamingResponse_Chunks{Chunks: chunks},
	}
	return res
}

func (self *MetathingsDevicedService) put_object_streaming_send_pull_request_loop(
	quit chan struct{},
	stm pb.DevicedService_PutObjectStreamingServer,
	fs file_helper.FileSyncer,
	sem chan struct{},
	errs chan error,
) {
	logger := self.get_logger()

	defer func() {
		close(quit)
		logger.Debugf("put object streaming send pull request loop exit")
	}()

	retry := 0

	for {
		select {
		case _, ok := <-sem:
			retry = 0
			if !ok {
				logger.Debugf("sem closed")
				return
			}
		case <-time.After(time.Duration(self.opt.Methods.PutObjectStreaming.PullRequestTimeout) * time.Second):
			logger.WithField("retry", retry).Debugf("pull request timeout, retry to send pull request")
			retry++
		}

		if retry >= self.opt.Methods.PutObjectStreaming.PullRequestRetry {
			err := ErrPutObjectStreamingTimeout
			logger.WithError(err).Warningf("pull request loop timeout")
			errs <- err
			return
		}

		offsets, err := fs.Next(self.opt.Methods.PutObjectStreaming.ChunkPerRequest)
		if err != nil {
			if err == file_helper.DONE {
				return
			}
			logger.WithError(err).Warningf("failed to get next offsets")
			errs <- err
			return
		}

		res := self.new_put_object_streaming_chunks_response(offsets)
		if err = stm.Send(res); err != nil {
			logger.WithError(err).Errorf("failed to send pull chunks request")
			errs <- err
			return
		}
	}
}

func (self *MetathingsDevicedService) put_object_streaming_recv_push_response_loop(
	quit chan struct{},
	stm pb.DevicedService_PutObjectStreamingServer,
	fs file_helper.FileSyncer,
	sem chan struct{},
	errs chan error,
) {
	logger := self.get_logger()

	defer func() {
		close(quit)
		logger.Debugf("put object streaming recv push response loop eixt")
	}()

	for {
		res, err := stm.Recv()
		if err != nil {
			logger.WithError(err).Errorf("failed to recv push chunks response")
			errs <- err
			return
		}

		chunks := res.GetChunks()
		if chunks == nil {
			logger.Warningf("empty chunks")
			continue
		}

		chks := chunks.GetChunks()
		for _, chk := range chks {
			size := int(chk.GetLength().GetValue())
			data := chk.GetData().GetValue()
			offset := chk.GetOffset().GetValue()
			if err = fs.Sync(offset, data, size); err != nil {
				if err == file_helper.DONE {
					return
				}
				logger.WithError(err).Warningf("failed to sync file")
				errs <- err
				return
			}
		}

		sem <- struct{}{}
	}
}

func (self *MetathingsDevicedService) PutObjectStreaming(stm pb.DevicedService_PutObjectStreamingServer) error {
	logger := self.get_logger().WithField("#method", "PutObjectStreaming")

	req, err := stm.Recv()
	if err != nil {
		logger.WithError(err).Errorf("failed to recv metadata request")
		return status.Errorf(codes.Internal, err.Error())
	}

	md_req := req.GetMetadata()
	if md_req == nil {
		logger.WithError(err).Errorf("metadata is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	obj := md_req.GetObject()
	if obj == nil {
		logger.WithError(err).Errorf("metadata.object is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	obj_s := parse_object(obj)

	sha1 := md_req.GetSha1()
	if sha1 == nil {
		logger.WithError(err).Errorf("metadata.sha1 is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	sha1_str := sha1.GetValue()

	fs, err := self.simple_storage.PutObjectAsync(obj_s, &ss.PutObjectAsyncOption{
		SHA1:      sha1_str,
		ChunkSize: self.opt.Methods.PutObjectStreaming.ChunkSize,
	})
	if err != nil {
		logger.WithError(err).Errorf("failed to put object async")
		return status.Errorf(codes.Internal, err.Error())
	}

	errs := make(chan error, 1)
	req_quit := make(chan struct{})
	res_quit := make(chan struct{})

	sem := make(chan struct{}, 1)
	defer close(sem)

	for i := 0; i < 1; i++ {
		sem <- struct{}{}
	}

	go self.put_object_streaming_send_pull_request_loop(req_quit, stm, fs, sem, errs)
	go self.put_object_streaming_recv_push_response_loop(res_quit, stm, fs, sem, errs)

	select {
	case <-req_quit:
		logger.Debugf("sync file")
	case <-res_quit:
		logger.Debugf("sync file")
	case err := <-errs:
		logger.WithError(err).Errorf("failed to sync file")
		if err == ErrPutObjectStreamingTimeout {
			return status.Errorf(codes.DeadlineExceeded, ErrPutObjectStreamingTimeout.Error())
		}
		return status.Errorf(codes.Internal, err.Error())
	case <-time.After(time.Duration(self.opt.Methods.PutObjectStreaming.Timeout) * time.Second):
		err = ErrPutObjectStreamingTimeout
		logger.WithError(err).Errorf("failed to sync file")
		return status.Errorf(codes.Internal, err.Error())
	}

	logger.Infof("put object streaming done")

	return nil
}
