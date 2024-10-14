package image

import (
	"github.com/boxesandglue/boxesandglue/backend/image"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/image"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

func newBackendImage(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	instance := &image.Image{}
	instanceValue := rt.ToValue(instance).(*goja.Object)
	instanceValue.SetPrototype(call.This.Prototype())
	return instanceValue
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("image", newBackendImage)
}
