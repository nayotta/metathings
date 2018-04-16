package grpc_helper

import (
	"errors"
	"strings"
)

var (
	InvalidFullMethodName = errors.New("invalid full method name")
)

type MethodDescription struct {
	Package string
	Service string
	Method  string
}

func ParseMethodDescription(fullMethodName string) (*MethodDescription, error) {
	pack_serv_meth := strings.Split(fullMethodName, "/")
	if len(pack_serv_meth) != 3 {
		return nil, InvalidFullMethodName
	}

	pack_serv := strings.SplitAfter(pack_serv_meth[1], ".")
	serv := pack_serv[len(pack_serv)-1]
	pack := pack_serv_meth[1][0 : len(pack_serv_meth[1])-len(serv)-1]
	meth := pack_serv_meth[2]

	return &MethodDescription{
		Package: pack,
		Service: serv,
		Method:  meth,
	}, nil
}
