package node

import (
	"github.com/boxesandglue/boxesandglue/backend/node"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:backend/node"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

func newBackendLinebreakSettings(call goja.ConstructorCall, rt *goja.Runtime) *goja.Object {
	instance := node.NewLinebreakSettings()
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

	func(runtime *goja.Runtime, module *goja.Object) {
		o := module.Get("exports").(*goja.Object)
		o.Set("linebreakSettings", newBackendLinebreakSettings)
		o.Set("linebreak", node.Linebreak)
		o.Set("appendLineEndAfter", node.AppendLineEndAfter)
		o.Set("boxit", node.Boxit)
		o.Set("copyList", node.CopyList)
		o.Set("dimensions", node.Dimensions)
		o.Set("hpack", node.Hpack)
		o.Set("hpackTo", node.HpackTo)
		o.Set("hpackToWithEnd", node.HpackToWithEnd)
		o.Set("insertAfter", node.InsertAfter)
		o.Set("insertBefore", node.InsertBefore)
		o.Set("newDisc", node.NewDisc)
		o.Set("newGlue", node.NewGlue)
		o.Set("newGlyph", node.NewGlyph)
		o.Set("newHList", node.NewHList)
		o.Set("newImage", node.NewImage)
		o.Set("newKern", node.NewKern)
		o.Set("newPenalty", node.NewPenalty)
		o.Set("newRule", node.NewRule)
		o.Set("newStartStop", node.NewStartStop)
		o.Set("newVList", node.NewVList)
		o.Set("tail", node.Tail)
		o.Set("vpack", node.Vpack)
		o.Set("debug", node.Debug)

	}(runtime, module)
}
