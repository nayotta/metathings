RM=$(shell which rm)
CLEAN_PATHS=./bin


all: identity_proto \
	core_proto \
	core_agent_proto \
	echo_proto \
	metathings_bin \
	metathingsd_bin

clean:
	$(RM) -rf $(CLEAN_PATHS)

identity_proto:
	$(MAKE) -C pkg/proto/identity all

echo_proto:
	$(MAKE) -C pkg/proto/echo all

core_proto:
	$(MAKE) -C pkg/proto/core all

core_agent_proto:
	$(MAKE) -C pkg/proto/core_agent all

metathings_bin:
	$(MAKE) -C cmd/metathings all

metathingsd_bin:
	$(MAKE) -C cmd/metathingsd all
