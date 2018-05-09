RM=$(shell which rm)
CLEAN_PATHS=./bin ./lib


all: \
	state_proto \
	identity_proto \
	core_proto \
	core_agent_proto \
	echo_proto \
	echo_plugins \
	switcher_proto \
	switcher_drivers \
	switcher_plugins \
	metathings_bin \
	metathingsd_bin

clean:
	$(RM) -rf $(CLEAN_PATHS)

state_proto:
	$(MAKE) -C pkg/proto/common/state all

identity_proto:
	$(MAKE) -C pkg/proto/identity all

core_proto:
	$(MAKE) -C pkg/proto/core all

core_agent_proto:
	$(MAKE) -C pkg/proto/core_agent all

echo_proto:
	$(MAKE) -C pkg/proto/echo all

switcher_proto:
	$(MAKE) -C pkg/proto/switcher all

metathings_bin:
	$(MAKE) -C cmd/metathings all

metathingsd_bin:
	$(MAKE) -C cmd/metathingsd all

echo_plugins:
	$(MAKE) -C pkg/echo/plugin all

switcher_drivers:
	$(MAKE) -C pkg/switcher/driver all

switcher_plugins:
	$(MAKE) -C pkg/switcher/plugin all

build_docker_images:
	./script/metathings_build.sh
	./script/metathingsd_build.sh
	./script/docker_build_metathingsd.sh
