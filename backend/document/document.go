package document

import (
	"os"

	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/document"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

func newBackendDocument(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	firstArg := call.Arguments[0]
	w, err := os.Create(firstArg.String())
	if err != nil {
		panic(rt.ToValue(err))
	}
	instance := document.NewDocument(w)
	instanceValue := rt.ToValue(instance).(*goja.Object)
	instanceValue.SetPrototype(call.This.Prototype())
	return instanceValue
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	o := module.Get("exports").(*goja.Object)
	o.Set("document", newBackendDocument)
}
