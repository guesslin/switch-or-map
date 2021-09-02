package plugs

import (
	"plugin"

	"gxn.tw/performance/switch-or-map/internal/helpers"
	"gxn.tw/performance/switch-or-map/pkg/plugs"
)

func LoadGetter(path string) plugs.Getter {
	p := loadPlugin(path)
	f, err := p.Lookup(plugs.FactoryName)
	helpers.PanicOn(err)
	factory, ok := f.(*plugs.Factory)
	if !ok {
		panic("assert failed")
	}
	return (*factory)()
}

func loadPlugin(path string) *plugin.Plugin {
	p, err := plugin.Open(path)
	helpers.PanicOn(err)
	return p
}
