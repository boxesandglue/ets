package lang

import (
	"github.com/boxesandglue/boxesandglue/backend/lang"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/lang"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("loadPatternFile", lang.LoadPatternFile)
		o.Set("lang", func() *lang.Lang { return &lang.Lang{} })
	}(runtime, module)
}
