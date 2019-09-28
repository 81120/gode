package core

import (
	"io/ioutil"
	"path/filepath"

	js "github.com/dop251/goja"
)

func moduleTemplate(c string) string {
	return "(function(module, exports) {" + c + "\n})"
}

func createModule(c *Core) *js.Object {
	r := c.GetRts()
	m := r.NewObject()
	e := r.NewObject()
	m.Set("exports", e)

	return m
}

func loadModule(c *Core, p string) *js.Program {
	p = filepath.Clean(p)
	pkg := c.Pkg[p]
	if pkg == nil {
		code, _ := ioutil.ReadFile(p)
		text := moduleTemplate(string(code))
		pkg, _ = js.Compile(p, text, false)
		c.Pkg[p] = pkg
	}
	return pkg
}

// RegisterLoader register a simple commonjs style loader to runtime
func RegisterLoader(c *Core) {
	r := c.GetRts()

	r.Set("require", func(call js.FunctionCall) js.Value {
		p := call.Argument(0).String()
		prg := loadModule(c, p)

		f, _ := r.RunProgram(prg)
		g, _ := js.AssertFunction(f)

		m := createModule(c)
		jsExports := m.Get("exports")
		g(jsExports, m, jsExports)

		return m.Get("exports")
	})
}
