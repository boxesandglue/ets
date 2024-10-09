package node

import (
	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/bag"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("sp", bag.SP)
		o.Set("mustSP", bag.MustSP)
		o.Set("factor", bag.Factor)
		o.Set("scaledPointFromFloat", bag.ScaledPointFromFloat)
		o.Set("max", bag.Max)
		o.Set("min", bag.Min)
		o.Set("maxSP", bag.MaxSP)
		o.Set("multiplyFloat", bag.MultiplyFloat)

	}(runtime, module)
}
