package font

import (
	"github.com/boxesandglue/boxesandglue/backend/font"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/font"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)

		o.Set("newFont", font.NewFont)
	}(runtime, module)
}
