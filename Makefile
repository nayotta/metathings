CLEAN_PATHS=./bin ./lib

clean:
	rm -rf $(CLEAN_PATHS)

gen_gopb:
	bash ./hack/proto/gen_gopb.sh

gen_pbset:
	bash ./hack/proto/gen_pbset.sh

precompile_ensure_outdir:
	mkdir -p bin

precompile_go_mod_vendor:
	go mod vendor

precompile_replace_casbin_server_proto:
	bash hack/casbin_server/replace_package_dir.sh

precompile: precompile_go_mod_vendor precompile_replace_casbin_server_proto precompile_ensure_outdir

compile_metathings: precompile
	go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o ./bin/metathings ./cmd/metathings/main.go

compile_metathingsd: precompile
	go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore" -o ./bin/metathingsd ./cmd/metathingsd/main.go
