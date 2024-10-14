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

func newBackendColor(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	instance := &color.Color{}
	instanceValue := rt.ToValue(instance).(*goja.Object)
	instanceValue.SetPrototype(call.This.Prototype())

	if len(call.Arguments) > 0 {
		firstArg := call.Arguments[0]
		obj := firstArg.ToObject(rt)

		for _, key := range obj.Keys() {
			instanceValue.Set(key, obj.Get(key))
		}
	}
	return instanceValue
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("color", newBackendColor)
	o.Set("colorNone", runtime.ToValue(color.ColorNone).ToNumber())
	o.Set("colorRGB", runtime.ToValue(color.ColorRGB).ToNumber())
	o.Set("colorCMYK", runtime.ToValue(color.ColorRGB).ToNumber())
	o.Set("colorGray", runtime.ToValue(color.ColorGray).ToNumber())
	o.Set("colorSpotcolor", runtime.ToValue(color.ColorSpotcolor).ToNumber())
}
