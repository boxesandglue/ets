package document

import (
	"os"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/boxesandglue/boxesandglue/backend/document"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/document"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

type jsBackend struct {
	runtime *goja.Runtime
}

func (jb *jsBackend) jsBackendNewDocument(call goja.FunctionCall) goja.Value {
	firstArg := call.Arguments[0]
	w, err := os.Create(firstArg.String())

	if err != nil {
		panic(jb.runtime.ToValue(err))
	}

	return jb.runtime.ToValue(document.NewDocument(w))
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	jb := &jsBackend{
		runtime: runtime,
	}

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("newDocument", jb.jsBackendNewDocument)
		o.Set("mustSP", bag.MustSP)

	}(runtime, module)
}
