PROTOC := $(shell which protoc)

GRPC_FLAGS := -I. -I/usr/local/include \
	-I$(GOPATH)/src

PROTO_FILES := service.proto \
	domain.proto \
	user.proto \
	create_domain.proto \
	delete_domain.proto \
	patch_domain.proto \
	get_domain.proto \
	list_domain.proto \
	create_user.proto \
	delete_user.proto \
	patch_user.proto \
	get_user.proto \
	list_user.proto

all:
	$(PROTOC) $(GRPC_FLAGS) \
	--go_out=plugins=grpc:. \
	--govalidators_out=. \
	$(PROTO_FILES)
