package core

import (
	"fmt"

	js "github.com/dop251/goja"
)

func log(call js.FunctionCall) js.Value {
	str := call.Argument(0)
	fmt.Print(str.String())
	return str
}

// RegisterConsole register a console.log to runtime
func RegisterConsole(c *Core) {
	r := c.GetRts()
	o := r.NewObject()
	o.Set("log", log)
	r.Set("console", o)
}
