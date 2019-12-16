package metathings_deviced_service

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	file_helper "github.com/nayotta/metathings/pkg/common/file"
	id_helper "github.com/nayotta/metathings/pkg/common/id"
	ss "github.com/nayotta/metathings/pkg/deviced/simple_storage"
	pb "github.com/nayotta/metathings/pkg/proto/deviced"
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
	fs *file_helper.FileSyncer,
	sem chan struct{},
) {
	logger := self.logger

	defer func() {
		close(quit)
		logger.Debugf("put object streaming send pull request loop exit")
	}()

	cnt := 0
	sc := make(chan struct{})
	defer close(sc)

_put_object_streaming_send_pull_request_loop:
	for {
		select {
		case <-sc:
			if cnt > 7 {
				logger.Warningf("recv push response timeout")
				break _put_object_streaming_send_pull_request_loop
			}
		case _, ok := <-sem:
			if !ok {
				logger.Debugf("sem closed")
				break _put_object_streaming_send_pull_request_loop
			}
			cnt = 0
		case <-time.After(3 * time.Second):
			cnt += 1
			sc <- struct{}{}
			continue _put_object_streaming_send_pull_request_loop
		}

		offsets, err := fs.Next(3)
		if err != nil {
			if err == file_helper.DONE {
				logger.Debugf("file sync done")
				return
			}
			logger.WithError(err).Warningf("failed to get next offsets")
			return
		}

		res := self.new_put_object_streaming_chunks_response(offsets)
		if err = stm.Send(res); err != nil {
			logger.WithError(err).Errorf("failed to send pull chunks request")
			return
		}
	}
}

func (self *MetathingsDevicedService) put_object_streaming_recv_push_response_loop(
	quit chan struct{},
	stm pb.DevicedService_PutObjectStreamingServer,
	fs *file_helper.FileSyncer,
	sem chan struct{},
) {
	logger := self.logger

	defer func() {
		close(quit)
		logger.Debugf("put object streaming recv push response loop eixt")
	}()

	for {
		res, err := stm.Recv()
		if err != nil {
			logger.WithError(err).Errorf("failed to rcev push chunks response")
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
					logger.Debugf("file sync done")
					return
				}
				logger.WithError(err).Warningf("failed to sync file")
				return
			}
		}

		sem <- struct{}{}
	}
}

func (self *MetathingsDevicedService) PutObjectStreaming(stm pb.DevicedService_PutObjectStreamingServer) error {
	req, err := stm.Recv()
	if err != nil {
		self.logger.WithError(err).Errorf("failed to recv metadata request")
		return status.Errorf(codes.Internal, err.Error())
	}

	md_req := req.GetMetadata()
	if md_req == nil {
		self.logger.WithError(err).Errorf("metadata is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}

	obj := md_req.GetObject()
	if obj == nil {
		self.logger.WithError(err).Errorf("metadata.object is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	obj_s := parse_object(obj)

	sha1 := md_req.GetSha1()
	if sha1 == nil {
		self.logger.WithError(err).Errorf("metadata.sha1 is empty")
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	sha1_str := sha1.GetValue()

	fs, err := self.simple_storage.PutObjectAsync(obj_s, &ss.PutObjectAsyncOption{
		SHA1:      sha1_str,
		ChunkSize: self.opt.Methods.PutObjectStreaming.ChunkSize,
	})
	if err != nil {
		self.logger.WithError(err).Errorf("failed to put object async")
		return status.Errorf(codes.Internal, err.Error())
	}

	req_quit := make(chan struct{})
	res_quit := make(chan struct{})

	sem := make(chan struct{})
	for i := 0; i < 3; i++ {
		sem <- struct{}{}
	}

	go self.put_object_streaming_send_pull_request_loop(req_quit, stm, fs, sem)
	go self.put_object_streaming_recv_push_response_loop(res_quit, stm, fs, sem)

	select {
	case <-req_quit:
	case <-res_quit:
	case <-time.After(600 * time.Second):
	}

	self.logger.Debugf("file synced")

	return nil
}
