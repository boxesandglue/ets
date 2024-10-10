package color

import (
	"github.com/boxesandglue/boxesandglue/backend/color"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/color"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("color", &color.Color{})
	o.Set("colorRGB", runtime.ToValue(color.ColorRGB).ToNumber())
	o.Set("colorCMYK", runtime.ToValue(color.ColorRGB).ToNumber())
	o.Set("colorGray", runtime.ToValue(color.ColorGray).ToNumber())
	o.Set("colorSpotcolor", runtime.ToValue(color.ColorSpotcolor).ToNumber())
}
