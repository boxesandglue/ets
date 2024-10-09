package core

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	pdf "github.com/boxesandglue/baseline-pdf"
	"github.com/boxesandglue/ets/console"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

// Configuration keeps the configuration for the javascript environment.
type Configuration struct {
	Starttime        time.Time
	Loglevel         string
	Verbose          bool
	Version          string
	protocolfilename string
}

var (
	cfg      Configuration
	pw       *pdf.PDF
	filename string
)

func runJavascript(vm *goja.Runtime, scriptname string) error {
	var err error
	if _, err = os.Stat(scriptname); err != nil {
		return err
	}
	data, err := os.ReadFile(scriptname)
	if err != nil {
		return err
	}
	_, err = vm.RunScript(scriptname, string(data))
	if err != nil {
		return err
	}
	return nil
}

// Runs the Javascript file if it exists. If the exename is "foo", it looks for
// a JS file called "foo.js"
func runDefaultJavascript(vm *goja.Runtime, exename string) error {
	var extension = filepath.Ext(exename)
	var name = exename[0:len(exename)-len(extension)] + ".js"
	var err error
	if _, err = os.Stat(name); err != nil {
		return nil
	}
	data, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	_, err = vm.RunScript(name, string(data))
	if err != nil {
		return err
	}
	return nil
}

func dothings(exename string, args []string) error {
	runtime := goja.New()
	var err error
	runtime.SetFieldNameMapper(goja.UncapFieldNameMapper())
	new(require.Registry).Enable(runtime)
	console.Enable(runtime)
	runtime.Set("toRunes", func(s string) []rune {
		return []rune(s)
	})

	if err = runDefaultJavascript(runtime, exename); err != nil {
		return err
	}
	if err = runJavascript(runtime, args[0]); err != nil {
		return err
	}

	if pdfwriter := runtime.Get("__pw"); pdfwriter != nil {
		pw = pdfwriter.Export().(*pdf.PDF)
	}
	if fn := runtime.Get("__filename"); fn != nil {
		filename = fn.ToString().String()
	}
	return nil
}

// SetCfg sets the configuration struct
func SetCfg(config Configuration) {
	cfg = config
}

func pluralize(what string, count int) string {
	switch what {
	default:
		if count == 1 {
			return "1 " + what
		}
		return fmt.Sprintf("%d %ss", count, what)
	}
}

// RunETS starts the execution of the javascript
func RunETS(exename string, args []string) error {
	defer teardownLog()
	err := dothings(exename, args)
	if err != nil {
		if exception, ok := err.(*goja.Exception); ok {
			stk := exception.Stack()
			slog.Error(exception.Error(), "position", stk[len(stk)-1].Position())
		}
		slog.Error(err.Error())
	}

	fmt.Printf("Finished with %s and %s\n", pluralize("error", errCount), pluralize("warning", warnCount))
	fmt.Printf("Transcript written to %s\n", cfg.protocolfilename)
	if pw != nil {
		fmt.Printf("Output written on %s (%s, %d bytes)\n", filename, pluralize("page", pw.NoPages), pw.Size())
	}
	slog.Info("Duration", "dur", time.Now().Sub(cfg.Starttime))
	return nil
}
