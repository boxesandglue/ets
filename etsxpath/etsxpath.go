package etsxpath

import (
	"os"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/speedata/goxpath"
)

const modulename = "xpath"

type jsXPath struct {
	rt *goja.Runtime
}

func (jm *jsXPath) newParser(call goja.FunctionCall) goja.Value {
	firstArg := call.Arguments[0]
	filename := firstArg.String()
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	p, err := goxpath.NewParser(r)
	if err != nil {
		panic(err)
	}
	ret := jm.rt.ToValue(p)
	return ret
}

func init() {
	require.RegisterNativeModule(modulename, Require)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	jm := jsXPath{
		rt: runtime,
	}
	o := module.Get("exports").(*goja.Object)
	o.Set("newParser", jm.newParser)
}
