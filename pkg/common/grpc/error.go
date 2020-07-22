package grpc_helper

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func expect_codes(err error, cds ...codes.Code) error {
	s, ok := status.FromError(err)
	if !ok {
		return err
	}

	scd := s.Code()
	for _, cd := range cds {
		if scd == cd {
			return nil
		}
	}

	return err
}

func ExpectCodes(err error, cds ...codes.Code) error {
	cds = append(cds, codes.OK)
	return expect_codes(err, cds...)
}

func ExpectCodesWithoutOK(err error, cds ...codes.Code) error {
	return expect_codes(err, cds...)
}

func ErrorWrapper(err error) error {
	s, ok := status.FromError(err)
	if !ok {
		return status.Errorf(codes.Internal, err.Error())
	}

	return s.Err()
}

type ErrorMapping map[error]codes.Code

type ErrorParser struct {
	em ErrorMapping
}

func (ep *ErrorParser) ParseError(err error) error {
	s, ok := status.FromError(err)
	if ok {
		return s.Err()
	}

	c, ok := ep.em[err]
	if !ok {
		c = codes.Internal
	}

	return status.Errorf(c, err.Error())
}

func NewErrorParser(em ErrorMapping) *ErrorParser {
	return &ErrorParser{em: em}
}
