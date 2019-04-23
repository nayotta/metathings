package metathings_component

import (
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"time"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type ModuleServiceLookuper interface {
	LookupUnaryCall()
	LookupStreamCall()
}

type ModuleServiceInitializer interface {
	InitModuleService(*Module) error
}

type ModuleOption struct {
	Config           string
	CredentialId     string
	CredentialSecret string
	Addresses        []string
}

type Module struct {
	krn    *Kernel
	tgt    interface{}
	srv    ModuleServer
	opt    *ModuleOption
	flags  *pflag.FlagSet
	logger log.FieldLogger
}

func (m *Module) init_flags() error {
	m.flags.StringVarP(&m.opt.Config, "config", "c", "", "Config file")
	m.flags.StringVar(&m.opt.CredentialId, "credential-id", "", "Module Credential Id")
	m.flags.StringVar(&m.opt.CredentialSecret, "credential-secret", "", "Module Credential Secret")
	m.flags.StringArrayVar(&m.opt.Addresses, "addresses", []string{"device:127.0.0.1:5002", "default:metathings.ai:21733"}, "Metathings Service Addresses")

	err := m.flags.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	return nil
}

func parse_service_endpoints(xs []string) (map[string]string, error) {
	ys := map[string]string{}

	for _, x := range xs {
		ts := strings.SplitN(x, ":", 2)
		if len(ts) != 2 {
			return nil, ErrBadServiceEndpoint
		}
		ys[ts[0]] = ts[1]
	}

	return ys, nil
}

func (m *Module) init_kernel() error {
	var err error

	opt := &NewKernelOption{}

	if m.opt.Config != "" {
		buf, err := ioutil.ReadFile(m.opt.Config)
		if err != nil {
			return err
		}
		opt.ConfigText = string(buf)
	} else {
		opt.Credential.Id = m.opt.CredentialId
		opt.Credential.Secret = m.opt.CredentialSecret
		opt.ServiceEndpoints, err = parse_service_endpoints(m.opt.Addresses)
		if err != nil {
			return err
		}
	}

	m.krn, err = NewKernel(opt)
	if err != nil {
		return err
	}

	return nil
}

func parse_scheme(x string) (string, string, error) {
	ts := strings.SplitN(x, "+", 2)
	if len(ts) == 1 {
		return ts[0], "grpc", nil
	}

	if len(ts) != 2 {
		return "", "", ErrBadScheme
	}

	if ts[1] == "" {
		ts[1] = "grpc"
	}

	return ts[0], ts[1], nil
}

func (m *Module) init_logger() error {
	var err error

	kc := m.Kernel().Config()
	m.logger, err = log_helper.NewLogger(kc.GetString("name"), kc.GetString("log.level"))
	if err != nil {
		return err
	}

	return nil
}

func (m *Module) init_server() error {
	var err error

	mdl_srv_initer, ok := m.tgt.(ModuleServiceInitializer)
	if ok {
		err = mdl_srv_initer.InitModuleService(m)
		if err != nil {
			return err
		}
	}

	protocol, adapter, err := parse_scheme(m.Kernel().Config().GetString("service.scheme"))
	if err != nil {
		return err
	}

	if protocol != "mtp" {
		return ErrBadScheme
	}

	m.srv, err = NewModuleServer(adapter, m)
	if err != nil {
		return err
	}

	return nil
}

func (m *Module) Kernel() *Kernel {
	return m.krn
}

func (m *Module) Target() interface{} {
	return m.tgt
}

func (m *Module) Logger() log.FieldLogger {
	return m.logger
}

func (m *Module) Init() error {
	var err error

	err = m.init_flags()
	if err != nil {
		return err
	}

	err = m.init_kernel()
	if err != nil {
		return err
	}

	err = m.init_logger()
	if err != nil {
		return err
	}

	err = m.init_server()
	if err != nil {
		return err
	}

	return nil
}

func (m *Module) HeartbeatLoop() {
	for {
		err := m.Kernel().Heartbeat()
		if err != nil {
			m.logger.WithError(err).Warningf("failed to heartbeat")
		}
		time.Sleep(time.Duration(m.Kernel().Config().GetInt("heartbeat.interval")) * time.Second)
	}
}

func (m *Module) Serve() error {
	go m.HeartbeatLoop()
	return m.srv.Serve()
}

func (m *Module) Stop() {
	m.srv.Stop()
}

func (m *Module) Launch() error {
	var err error

	if err = m.Init(); err != nil {
		return err
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		m.Stop()
	}()

	return m.Serve()
}

func NewModule(name string, target interface{}) (*Module, error) {
	return &Module{
		tgt:   target,
		opt:   &ModuleOption{},
		flags: pflag.NewFlagSet(name, pflag.ExitOnError),
	}, nil
}
