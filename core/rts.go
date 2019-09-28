package core

import (
	js "github.com/dop251/goja"
)

// Core is the basic struct of gode
type Core struct {
	Rts *js.Runtime
	Pkg map[string]js.Value
}

// New create a *Core
func New() *Core {
	vm := js.New()
	pkg := make(map[string]js.Value)

	return &Core{
		Rts: vm,
		Pkg: pkg,
	}
}

// GetRts get the object of javascript runtime
func (c *Core) GetRts() *js.Runtime {
	return c.Rts
}

// RegisterBuildInModule register some build in modules to the runtime
func (c *Core) RegisterBuildInModule() {
	RegisterConsole(c)
	RegisterLoader(c)
}
