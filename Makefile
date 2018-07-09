RM=$(shell which rm)
CLEAN_PATHS=./bin ./lib

all: \
	protos \
	echo_plugins \
	switcher_drivers \
	switcher_plugins \
	motor_drivers \
	motor_plugins \
	camera_drivers \
	camera_plugins \
	servo_drivers \
	servo_plugins \
	metathings_bin \
	metathingsd_bin

clean:
	$(RM) -rf $(CLEAN_PATHS)

protos:
	$(MAKE) -C pkg/proto all

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

motor_drivers:
	$(MAKE) -C pkg/motor/driver all

motor_plugins:
	$(MAKE) -C pkg/motor/plugin all

camera_drivers:
	$(MAKE) -C pkg/camera/driver all

camera_plugins:
	$(MAKE) -C pkg/camera/plugin all

servo_drivers:
	$(MAKE) -C pkg/servo/driver all

servo_plugins:
	$(MAKE) -C pkg/servo/plugin all

build_docker_images:
	./script/metathings_build.sh
	./script/metathingsd_build.sh
	./script/docker_build_metathingsd.sh
