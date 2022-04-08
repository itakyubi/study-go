package context

import (
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"study-go/web/log"
	"study-go/web/utils"
	"sync"
	"syscall"
)

type Context interface {
	ConfFile() string
	Log() *log.Logger
	Wait()
	WaitChan() <-chan os.Signal
	LoadCustomConfig(cfg interface{}, files ...string) error
}

type ctx struct {
	sync.Map // global cache
	log      *log.Logger
}

func NewContext(confFile string) Context {
	if confFile == "" {
		confFile = os.Getenv(ConfFile)
	}

	c := &ctx{}
	c.Store(ConfFile, confFile)
	c.log = log.L()

	sc := &SystemConfig{}
	err := c.LoadCustomConfig(sc)
	if err != nil {
		c.log.Error("failed to load system config, to use default config", log.Error(err))
		utils.UnmarshalYAML(nil, sc)
	}

	_log, err := log.Init(sc.Logger)
	if err != nil {
		c.log.Error("failed to init logger", log.Error(err))
	}
	c.log = _log
	c.log.Debug("context is created", log.Any("file", confFile))
	return c
}

func (c *ctx) ConfFile() string {
	v, ok := c.Load(ConfFile)
	if !ok {
		return ""
	}
	return v.(string)
}

func (c *ctx) Log() *log.Logger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
}

func (c *ctx) WaitChan() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	return sig
}

func (c *ctx) LoadCustomConfig(cfg interface{}, files ...string) error {
	f := c.ConfFile()
	if len(files) > 0 && len(files[0]) > 0 {
		f = files[0]
	}
	if utils.FileExists(f) {
		return errors.WithStack(utils.LoadYAML(f, cfg))
	}
	return errors.WithStack(utils.UnmarshalYAML(nil, cfg))
}
