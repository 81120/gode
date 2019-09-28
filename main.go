package main

import (
	"fmt"

	"github.com/81120/gode/core"
)

func main() {
	gode := core.New()
	gode.RegisterBuildInModule()

	r := gode.GetRts()
	v, err := r.RunString(`
		var t = require('./test.js');
		t.test();
	`)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(v)
	}
}
