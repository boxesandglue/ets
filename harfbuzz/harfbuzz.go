package harfbuzz

import (
	"github.com/boxesandglue/textlayout/harfbuzz"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:harfbuzz"

func init() {
	require.RegisterNativeModule(modulename, Require)

}

type jsHarfbuzz struct {
	runtime *goja.Runtime
}

func (jh *jsHarfbuzz) jsHarfbuzzFeatures(call goja.FunctionCall) goja.Value {
	f := make([]harfbuzz.Feature, 0, len(call.Arguments))
	for _, arg := range call.Arguments {
		feat, err := harfbuzz.ParseFeature(arg.String())
		if err != nil {
			panic(err)
		}
		f = append(f, feat)
	}
	return jh.runtime.ToValue(f)
}
func (jh *jsHarfbuzz) jsHarfbuzzNewBuffer(call goja.FunctionCall) goja.Value {
	hb := harfbuzz.NewBuffer()
	return jh.runtime.ToValue(hb)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	jp := &jsHarfbuzz{
		runtime: runtime,
	}

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("newBuffer", jp.jsHarfbuzzNewBuffer)
		o.Set("features", jp.jsHarfbuzzFeatures)
	}(runtime, module)
}
