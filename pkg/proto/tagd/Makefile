PROTOC := $(shell which protoc)

GRPC_FLAGS := -I. -I/usr/local/include \
	-I$(GOPATH)/src

PROTO_FILES := $(shell ls *.proto)

all:
	$(PROTOC) $(GRPC_FLAGS) \
	--descriptor_set_out=service.protoset \
	--include_imports \
	--go_out=plugins=grpc:. \
	--govalidators_out=. \
	$(PROTO_FILES)
