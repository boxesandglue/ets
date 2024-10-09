package console

import (
	"log/slog"

	"github.com/boxesandglue/boxesandglue/backend/bag"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/util"
)

const ModuleName = "node:console"

type Console struct {
	runtime *goja.Runtime
	util    *goja.Object
}

func (c *Console) makeLogFunc(loglevel slog.Level) func(call goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		if format, ok := goja.AssertFunction(c.util.Get("format")); ok {
			ret, err := format(c.util, call.Arguments...)
			if err != nil {
				panic(err)
			}
			retStr := ret.String()
			switch loglevel {
			case slog.LevelDebug:
				bag.Logger.Debug(retStr)
			case slog.LevelInfo:
				bag.Logger.Info(retStr)
			case slog.LevelWarn:
				bag.Logger.Warn(retStr)
			case slog.LevelError:
				bag.Logger.Error(retStr)
			}
		} else {
			panic(c.runtime.NewTypeError("util.format is not a function"))
		}

		return nil
	}
}

func (c *Console) log(p func(string)) func(goja.FunctionCall) goja.Value {
	return func(call goja.FunctionCall) goja.Value {
		if format, ok := goja.AssertFunction(c.util.Get("format")); ok {
			ret, err := format(c.util, call.Arguments...)
			if err != nil {
				panic(err)
			}

			p(ret.String())
		} else {
			panic(c.runtime.NewTypeError("util.format is not a function"))
		}

		return nil
	}
}

func Require(runtime *goja.Runtime, module *goja.Object) {

	c := &Console{
		runtime: runtime,
	}

	c.util = require.Require(runtime, util.ModuleName).(*goja.Object)

	o := module.Get("exports").(*goja.Object)
	o.Set("debug", c.makeLogFunc(slog.LevelDebug))
	o.Set("info", c.makeLogFunc(slog.LevelInfo))
	o.Set("log", c.makeLogFunc(slog.LevelInfo))
	o.Set("warn", c.makeLogFunc(slog.LevelWarn))
}

func Enable(runtime *goja.Runtime) {
	runtime.Set("console", require.Require(runtime, ModuleName))
}

func init() {
	require.RegisterNativeModule(ModuleName, Require)
}
