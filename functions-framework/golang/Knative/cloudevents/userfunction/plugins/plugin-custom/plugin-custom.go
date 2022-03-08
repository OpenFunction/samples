package plugin_custom

import (
	ofctx "github.com/OpenFunction/functions-framework-go/context"
	"github.com/OpenFunction/functions-framework-go/plugin"
	"github.com/fatih/structs"
)

const (
	Name    = "plugin-custom"
	Version = "v1"
)

type PluginCustom struct {
	PluginName    string
	PluginVersion string
	stateC        int64
	StateC        int64
}

var _ plugin.Plugin = &PluginCustom{}

func New() *PluginCustom {
	return &PluginCustom{}
}

func (p *PluginCustom) Name() string {
	return Name
}

func (p *PluginCustom) Version() string {
	return Version
}

func (p *PluginCustom) Init() plugin.Plugin {
	return New()
}

func (p *PluginCustom) ExecPreHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	p.stateC = 3
	p.StateC = 3
	return nil
}

func (p *PluginCustom) ExecPostHook(ctx ofctx.RuntimeContext, plugins map[string]plugin.Plugin) error {
	return nil
}

func (p *PluginCustom) Get(fieldName string) (interface{}, bool) {
	plgMap := structs.Map(p)
	value, ok := plgMap[fieldName]
	return value, ok
}
