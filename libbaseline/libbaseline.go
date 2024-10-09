package libbaseline

import (
	"os"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:baseline"

func init() {
	require.RegisterNativeModule(modulename, Require)

}

type jsPDF struct {
	runtime *goja.Runtime
}

type jsPDFString struct {
	runtime *goja.Runtime
}

func (jp *jsPDF) jsPDFString(call goja.FunctionCall) goja.Value {
	firstARg := call.Arguments[0]
	str := pdf.String(firstARg.String())
	return jp.runtime.ToValue(str)
}

func (jp *jsPDF) jsPDFNew(call goja.FunctionCall) goja.Value {
	firstArg := call.Arguments[0]
	w, err := os.Create(firstArg.String())

	if err != nil {
		panic(jp.runtime.ToValue(err))
	}

	pw := pdf.NewPDFWriter(w)
	return jp.runtime.ToValue(pw)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	jp := &jsPDF{
		runtime: runtime,
	}

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("new", jp.jsPDFNew)
		o.Set("string", jp.jsPDFString)
		o.Set("serialize", func(call goja.FunctionCall) goja.Value {
			firstArg := call.Arguments[0]
			return runtime.ToValue(pdf.Serialize(firstArg.Export()))
		})
		o.Set("nameDest", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(&pdf.NameDest{})
		})
		o.Set("outline", func(call goja.FunctionCall) goja.Value {
			return runtime.ToValue(&pdf.Outline{})
		})
		o.Set("annotation", func() *pdf.Annotation { return &pdf.Annotation{} })

	}(runtime, module)
}
