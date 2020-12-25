package metathings_component

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/stretchr/objx"

	log_helper "github.com/nayotta/metathings/pkg/common/log"
	version_helper "github.com/nayotta/metathings/pkg/common/version"
	deviced_pb "github.com/nayotta/metathings/proto/deviced"
)

type ModuleServiceLookuper interface {
	LookupUnaryCall()
	LookupStreamCall()
}

type ModuleServiceInitializer interface {
	InitModuleService(*Module) error
}

type ModuleOption struct {
	Config              string
	CredentialId        string
	CredentialSecret    string
	ServiceEndpoints    map[string]ServiceEndpoint
	TransportCredential TransportCredential
}

type Module struct {
	version_helper.Versioner

	name_once *sync.Once
	name      string

	tgt interface{}

	krn    *Kernel
	srv    ModuleServer
	opt    *ModuleOption
	args   []string
	flags  *pflag.FlagSet
	logger log.FieldLogger
	closed chan struct{}
}

func (m *Module) init_flags() error {
	m.flags.StringVarP(&m.opt.Config, "config", "c", "", "Config file")
	m.flags.StringVar(&m.opt.CredentialId, "credential-id", "", "Module Credential Id")
	m.flags.StringVar(&m.opt.CredentialSecret, "credential-secret", "", "Module Credential Secret")
	m.flags.BoolVar(&m.opt.TransportCredential.Insecure, "insecure", false, "Transport data in tls with insecure mode")
	m.flags.BoolVar(&m.opt.TransportCredential.PlainText, "plaintext", false, "Transport data without tls")
	m.flags.StringVar(&m.opt.TransportCredential.KeyFile, "key-file", "", "Transport credential key")
	m.flags.StringVar(&m.opt.TransportCredential.CertFile, "cert-file", "", "Transport credential cert")

	err := m.flags.Parse(m.args)
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
		opt.ServiceEndpoints = m.opt.ServiceEndpoints
		opt.TransportCredential.Insecure = m.opt.TransportCredential.Insecure
		opt.TransportCredential.PlainText = m.opt.TransportCredential.PlainText
		opt.TransportCredential.KeyFile = m.opt.TransportCredential.KeyFile
		opt.TransportCredential.CertFile = m.opt.TransportCredential.CertFile
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

func (m *Module) init_version() error {
	return m.Kernel().PutObject(fmt.Sprintf("/sys/firmware/modules/%s/version/current", m.Name()), strings.NewReader(m.GetVersion()))
}

func (m *Module) Name() string {
	m.name_once.Do(func() {
		mdl, err := m.Kernel().Show()
		// TODO(Peer): should not panic
		if err != nil {
			panic(err)
		}
		m.name = mdl.Name
	})

	return m.name
}

func (m *Module) WithNamespace(name string) string {
	return path.Join("modules", m.Name(), name)
}

func (m *Module) PutObject(name string, content io.Reader) error {
	return m.Kernel().PutObject(m.WithNamespace(name), content)
}

func (m *Module) PutObjects(objects map[string]io.Reader) error {
	with_namespace_objects := make(map[string]io.Reader)
	for name, content := range objects {
		with_namespace_objects[m.WithNamespace(name)] = content
	}

	return m.Kernel().PutObjects(with_namespace_objects)
}

func (m *Module) PutObjectStreaming(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) error {
	return m.Kernel().PutObjectStreaming(m.WithNamespace(name), content, opt)
}

func (m *Module) PutObjectStreamingWithCancel(name string, content io.ReadSeeker, opt *PutObjectStreamingOption) (context.CancelFunc, chan error, error) {
	return m.Kernel().PutObjectStreamingWithCancel(m.WithNamespace(name), content, opt)
}

func (m *Module) GetObject(name string) (*deviced_pb.Object, error) {
	return m.Kernel().GetObject(m.WithNamespace(name))
}

func (m *Module) GetObjectContent(name string) ([]byte, error) {
	return m.Kernel().GetObjectContent(m.WithNamespace(name))
}

func (m *Module) RemoveObject(name string) error {
	return m.Kernel().RemoveObject(m.WithNamespace(name))
}

func (m *Module) RemoveObjects(names []string) error {
	with_namespace_names := []string{}
	for _, name := range names {
		with_namespace_names = append(with_namespace_names, m.WithNamespace(name))
	}

	return m.Kernel().RemoveObjects(with_namespace_names)
}

func (m *Module) RenameObject(src, dst string) error {
	return m.Kernel().RenameObject(m.WithNamespace(src), m.WithNamespace(dst))
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

	err = m.init_version()
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
	for m.IsRunning() {
		err := m.Kernel().Heartbeat()
		if err != nil {
			m.logger.WithError(err).Warningf("failed to heartbeat")
		}
		time.Sleep(time.Duration(m.Kernel().Config().GetInt("heartbeat.interval")) * time.Second)
	}
}

func (m *Module) Serve() error {
	logger := m.Logger()
	cfg := m.Kernel().Config()

	hbs := cfg.GetString("heartbeat.strategy")
	if hbs == "" {
		hbs = "auto"
	}
	logger.WithFields(log.Fields{
		"heartbeat_strategy": hbs,
	}).Info("module serve")

	switch hbs {
	case "auto":
		go m.HeartbeatLoop()
	default:
	}

	return m.srv.Serve()
}

func (m *Module) Stop() {
	m.srv.Stop()

	if m.IsRunning() {
		close(m.closed)
	}
}

func (m *Module) IsRunning() bool {
	select {
	case _, alive := <-m.closed:
		return alive
	default:
		return true
	}
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

func NewDefaultModuleOption() objx.Map {
	return objx.New(map[string]interface{}{
		"version": "unknown",
		"args":    os.Args[1:],
	})
}

func NewModule(opts ...NewModuleOption) (*Module, error) {
	o := NewDefaultModuleOption()

	for _, opt := range opts {
		opt(o)
	}

	return &Module{
		Versioner: version_helper.NewVersioner(o.Get("version").String())(),
		name_once: new(sync.Once),
		tgt:       o.Get("target").Inter(),
		opt:       &ModuleOption{},
		args:      o.Get("args").StringSlice(),
		flags:     pflag.NewFlagSet("module", pflag.ExitOnError),
		closed:    make(chan struct{}),
	}, nil
}
