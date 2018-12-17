package metathingsdevicecloudmqttbridge

import "errors"

// ErrInvalidArgument ErrInvalidArgument
var ErrInvalidArgument = errors.New("invalid argument")

// ErrUnexpectedResponse ErrUnexpectedResponse
var ErrUnexpectedResponse = errors.New("unexpected response")

// ErrMqttMsgBlank ErrMqttMsgBlank
var ErrMqttMsgBlank = errors.New("publish blank msg")

// ErrMqttPubFailed ErrMqttPubFailed
var ErrMqttPubFailed = errors.New("publish msg failed")

// ErrMqttSubFailed ErrMqttSubFailed
var ErrMqttSubFailed = errors.New("subscribe failed")

// ErrMqttKeygenFailed ErrMqttKeygenFailed
var ErrMqttKeygenFailed = errors.New("keygen failed")

// ErrMqttUpKeygenFailed ErrMqttUpKeygenFailed
var ErrMqttUpKeygenFailed = errors.New("up keygen failed")

// ErrMqttStatusKeygenFailed ErrMqttStatusKeygenFailed
var ErrMqttStatusKeygenFailed = errors.New("status keygen failed")

// ErrMqttDownKeygenFailed ErrMqttDownKeygenFailed
var ErrMqttDownKeygenFailed = errors.New("down keygen failed")

//ErrMqttConnectFailed ErrMqttConnectFailed
var ErrMqttConnectFailed = errors.New("mqtt connect failed")

// ErrMqttRequestTimeout ErrMqttRequestTimeout
var ErrMqttRequestTimeout = errors.New("mqtt request timeout")

// ErrMqttStreamCallConfigError ErrMqttStreamCallConfigError
var ErrMqttStreamCallConfigError = errors.New("mqtt streamcall config error")

// ErrMqttStreamCallHearBeatTimeoutError ErrMqttStreamCallHearBeatTimeoutError
var ErrMqttStreamCallHearBeatTimeoutError = errors.New("mqtt streamcall heartbeat timeout")

// ErrMqttDisconnectedError ErrMqttDisconnectedError
var ErrMqttDisconnectedError = errors.New("mqtt connection disconnected")
