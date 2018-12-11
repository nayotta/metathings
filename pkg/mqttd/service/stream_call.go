package metathingsmqttdservice

import (
	pb "github.com/nayotta/metathings/pkg/proto/mqttd"
)

// StreamCall StreamCall
func (serv *MetathingsMqttdService) StreamCall(stream pb.MqttdService_StreamCallServer) error {

	// TODO(zh) streamcall
	/*
		if conn, err := serv.cc.StreamCall(stream); err != nil {
			serv.logger.WithError(err).Errorf("failed to stream call")
			return status.Errorf(codes.Internal, err.Error())
		}

		serv.logger.Debugf("stream call")

		<-conn.Wait()
		serv.logger.Debugf("stream call closed")

		if err = conn.Err(); err != nil {
			serv.logger.WithError(err).Errorf("stream call error")
			return status.Errorf(codes.Internal, err.Error())
		}
	*/

	return nil
}
