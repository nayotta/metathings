RM=$(shell which rm)
CLEAN_PATHS=./bin


all: identity_proto echo_proto metathings_bin metathingsd_bin

clean:
	$(RM) -rf $(CLEAN_PATHS)

identity_proto:
	$(MAKE) -C pkg/proto/identity all

echo_proto:
	$(MAKE) -C pkg/proto/echo all

metathings_bin:
	$(MAKE) -C cmd/metathings all

metathingsd_bin:
	$(MAKE) -C cmd/metathingsd all
