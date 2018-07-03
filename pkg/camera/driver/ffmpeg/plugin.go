package main

import (
	"context"
	"os/exec"
	"sync"

	"github.com/cbroglie/mustache"
	"github.com/nayotta/viper"
	log "github.com/sirupsen/logrus"

	driver "github.com/nayotta/metathings/pkg/camera/driver"
	driver_helper "github.com/nayotta/metathings/pkg/common/driver"
	opt_helper "github.com/nayotta/metathings/pkg/common/option"
)

const (
	FFMPEG_TEMPLATE  = `ffmpeg -y -f v4l2 -pix_fmt yuv420p -video_size {{width}}x{{height}} -i {{device}} -r {{framerate}} -b:v {{bitrate}} -c:v h264_omx -zerocopy 1 -f flv {{url}}`
	FFMPEG_DEVICE    = "/dev/video0"
	FFMPEG_WIDTH     = 640
	FFMPEG_HEIGHT    = 480
	FFMPEG_BITRATE   = 500000 // 500 kbits
	FFMPEG_FRAMERATE = 24     // 24 Hz
)

type driverOption struct {
	Device    string
	Width     uint32
	Height    uint32
	Bitrate   uint32
	Framerate uint32
	Template  string
}

func defaultOptions() driverOption {
	return driverOption{
		Device:    FFMPEG_DEVICE,
		Width:     FFMPEG_WIDTH,
		Height:    FFMPEG_HEIGHT,
		Bitrate:   FFMPEG_BITRATE,
		Framerate: FFMPEG_FRAMERATE,
		Template:  FFMPEG_TEMPLATE,
	}
}

type ffmpegCameraDriver struct {
	mutex  *sync.Mutex
	state  driver.CameraState
	config driver.CameraConfig

	logger log.FieldLogger
	opt    driverOption

	defaultConfig driver.CameraConfig
	ffmpegCmd     *exec.Cmd
	cancel        context.CancelFunc

	state_notification_broadcast_channels map[chan driver.CameraState]interface{}
}

func (drv *ffmpegCameraDriver) GetStateNotificationChannel() chan driver.CameraState {
	ch := make(chan driver.CameraState)
	drv.state_notification_broadcast_channels[ch] = nil
	return ch
}

func (drv *ffmpegCameraDriver) CloseStateNotificationChannel(ch chan driver.CameraState) {
	if _, ok := drv.state_notification_broadcast_channels[ch]; ok {
		delete(drv.state_notification_broadcast_channels, ch)
	}
	close(ch)
}

func (drv *ffmpegCameraDriver) setState(s driver.CameraState) {
	drv.state = s
	for ch, _ := range drv.state_notification_broadcast_channels {
		ch <- s
	}
	drv.logger.WithField("state", s).Debugf("update state")
}

func (drv *ffmpegCameraDriver) show() driver.Camera {
	return driver.Camera{
		State:  drv.state,
		Config: drv.config,
	}
}

func (drv *ffmpegCameraDriver) readConfig() driver.CameraConfig {
	return driver.CameraConfig{
		Device:    drv.opt.Device,
		Width:     drv.opt.Width,
		Height:    drv.opt.Height,
		Bitrate:   drv.opt.Bitrate,
		Framerate: drv.opt.Framerate,
	}
}

func (drv *ffmpegCameraDriver) Init(opt opt_helper.Option) error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	drv.state = driver.STATE_STOP
	v, ok := opt.Get("driver").(*viper.Viper)
	if !ok {
		return driver_helper.ErrInitFail
	}

	logger, ok := opt.Get("logger").(log.FieldLogger)
	if !ok {
		return driver_helper.ErrInitFail
	}
	drv.logger = logger

	var o driverOption
	err := v.Unmarshal(&o)
	if err != nil {
		return err
	}

	if o.Template != "" {
		drv.opt.Template = o.Template
	}

	if o.Device != "" {
		drv.opt.Device = o.Device
	}

	if o.Width != 0 && o.Height != 0 {
		drv.opt.Width = o.Width
		drv.opt.Height = o.Height
	}

	if o.Bitrate != 0 {
		drv.opt.Bitrate = o.Bitrate
	}

	if o.Framerate != 0 {
		drv.opt.Framerate = o.Framerate
	}

	drv.logger.WithFields(log.Fields{
		"template":  drv.opt.Template,
		"device":    drv.opt.Device,
		"width":     drv.opt.Width,
		"height":    drv.opt.Height,
		"bitrate":   drv.opt.Bitrate,
		"framerate": drv.opt.Framerate,
	}).Debugf("ffmpeg default params")
	drv.logger.Debugf("camera plugin(ffmpeg) initialized")

	return nil
}

func (drv *ffmpegCameraDriver) Close() error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state != driver.STATE_STOP {
		err := drv.stop()
		if err != nil {
			return err
		}
	}
	return nil
}

func (drv *ffmpegCameraDriver) Show() (driver.Camera, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	return drv.show(), nil
}

func (drv *ffmpegCameraDriver) prepareConfig(cfg driver.CameraConfig) (driver.CameraConfig, error) {
	c := drv.readConfig()

	if cfg.Url == "" {
		return cfg, driver.ErrInvalidArgument
	}
	c.Url = cfg.Url

	if cfg.Device != "" {
		c.Device = cfg.Device
	}

	if cfg.Width != 0 {
		c.Width = cfg.Width
	}

	if cfg.Height != 0 {
		c.Height = cfg.Height
	}

	if cfg.Bitrate != 0 {
		c.Bitrate = cfg.Bitrate
	}

	if cfg.Framerate != 0 {
		c.Framerate = cfg.Framerate
	}

	return c, nil
}

func (drv *ffmpegCameraDriver) Start(cfg driver.CameraConfig) (driver.Camera, error) {
	var err error
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state != driver.STATE_STOP {
		return drv.show(), driver.ErrUnstartable
	}

	drv.config, err = drv.prepareConfig(cfg)
	if err != nil {
		return drv.show(), err
	}
	drv.setState(driver.STATE_STARTING)
	go func() {
		drv.mutex.Lock()
		defer drv.mutex.Unlock()

		drv.logger.WithFields(log.Fields{
			"url":       drv.config.Url,
			"device":    drv.config.Device,
			"width":     drv.config.Width,
			"height":    drv.config.Height,
			"framerate": drv.config.Framerate,
			"bitrate":   drv.config.Bitrate,
		}).Debugf("start camera")

		rdr_opts := map[string]interface{}{
			"url":       drv.config.Url,
			"device":    drv.config.Device,
			"width":     drv.config.Width,
			"height":    drv.config.Height,
			"framerate": drv.config.Framerate,
			"bitrate":   drv.config.Bitrate,
		}
		cmd_str, err := mustache.Render(drv.opt.Template, rdr_opts)
		if err != nil {
			drv.logger.WithError(err).Errorf("failed to render ffmpeg template")
			drv.setState(driver.STATE_STOP)
			return
		}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		drv.ffmpegCmd = exec.CommandContext(ctx, "/bin/bash", "-c", cmd_str)
		drv.cancel = cancel
		err = drv.ffmpegCmd.Start()
		if err != nil {
			drv.logger.WithError(err).Errorf("failed to start ffmpeg")
			drv.setState(driver.STATE_STOP)
			return
		}
		drv.logger.WithField("cmd", cmd_str).Debugf("start ffmpeg")
		go func() {
			err := drv.ffmpegCmd.Wait()

			drv.mutex.Lock()
			defer drv.mutex.Unlock()

			if err != nil && drv.state == driver.STATE_RUNNING {
				drv.reset()
				drv.setState(driver.STATE_STOP)
				drv.logger.WithError(err).Errorf("ffmpeg unexpected exit")
			}

			drv.logger.Debugf("ffmpeg exit")
		}()
		drv.setState(driver.STATE_RUNNING)
		drv.logger.Debugf("camera is running")
	}()
	drv.logger.Debugf("camera is starting")

	return drv.show(), nil
}

func (drv *ffmpegCameraDriver) reset() {
	drv.ffmpegCmd = nil
	drv.cancel = nil
	drv.config.Url = ""
}

func (drv *ffmpegCameraDriver) stop() error {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.cancel != nil {
		drv.cancel()
	}
	drv.reset()
	drv.setState(driver.STATE_STOP)
	drv.logger.Debugf("camera stopped")

	return nil
}

func (drv *ffmpegCameraDriver) Stop() (driver.Camera, error) {
	drv.mutex.Lock()
	defer drv.mutex.Unlock()

	if drv.state != driver.STATE_RUNNING {
		return drv.show(), driver.ErrUnstopable
	}

	drv.setState(driver.STATE_TERMINATING)
	go drv.stop()
	drv.logger.Debugf("camera is terminating")

	return driver.Camera{}, nil
}

var NewDriver driver_helper.NewDriverMethod = func(opt opt_helper.Option) (driver_helper.Driver, error) {
	return &ffmpegCameraDriver{
		mutex: &sync.Mutex{},
		state: driver.STATE_UNKNOWN,
		opt:   defaultOptions(),
		state_notification_broadcast_channels: make(map[chan driver.CameraState]interface{}),
	}, nil
}
