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

func newBackendFont(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	instance := &font.Font{}
	instanceValue := rt.ToValue(instance).(*goja.Object)
	instanceValue.SetPrototype(call.This.Prototype())
	return instanceValue
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("font", newBackendFont)
	o.Set("newFont", font.NewFont)
}
