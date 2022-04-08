package plugin

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strings"
	"study-go/web/log"
	"sync"
)

type Plugin interface {
	io.Closer
}

type Factory func() (Plugin, error)

var pluginFactory sync.Map
var plugins sync.Map

func RegisterFactory(name string, f Factory) {
	if _, ok := pluginFactory.Load(name); ok {
		log.L().Info("plugin already exists, skip", log.Any("plugin", name))
		return
	}
	pluginFactory.Store(name, f)
	log.L().Info("plugin is registered", log.Any("plugin", name))
}

func GetPlugin(name string) (Plugin, error) {
	name = strings.ToLower(name)
	if p, ok := plugins.Load(name); ok {
		return p.(Plugin), nil
	}

	f, ok := pluginFactory.Load(name)
	if !ok {
		return nil, errors.New(fmt.Sprintf("plugin {%s} not found", name))
	}

	p, err := f.(Factory)()
	if err != nil {
		log.L().Error("failed to create plugin", log.Error(err))
		return nil, err
	}

	act, ok := plugins.LoadOrStore(name, p)
	if ok {
		err := p.Close()
		if err != nil {
			log.L().Warn("failed to close plugin", log.Error(err))
		}
		return act.(Plugin), nil
	}
	return p, nil
}

func ClosePlugins() {
	plugins.Range(func(key, value interface{}) bool {
		value.(Plugin).Close()
		return true
	})
}
