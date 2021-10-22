package plugin

import (
	"fmt"
	"io"
	"strings"
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
		println(fmt.Sprintf("plugin {%s} already exists, skip", name))
		return
	}
	pluginFactory.Store(name, f)
	println(fmt.Sprintf("plugin {%s} registered", name))
}

func GetPlugin(name string) (Plugin, error) {
	name = strings.ToLower(name)
	if p, ok := plugins.Load(name); ok {
		return p.(Plugin), nil
	}

	f, ok := pluginFactory.Load(name)
	if !ok {
		println(fmt.Sprintf("plugin {%s} not found", name))
		return nil, nil
	}

	p, err := f.(Factory)()
	if err != nil {
		println(fmt.Sprintf("create plugin {%s} failed.", name))
		return nil, err
	}

	act, ok := plugins.LoadOrStore(name, p)
	if ok {
		err := p.Close()
		if err != nil {
			println(fmt.Sprintf("close plugin {%s} failed.", name))
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