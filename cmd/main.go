package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/boxesandglue/ets/core"
	"github.com/pelletier/go-toml/v2"
	"github.com/speedata/optionparser"

	_ "github.com/boxesandglue/ets/backend/bag"
	_ "github.com/boxesandglue/ets/backend/document"
	_ "github.com/boxesandglue/ets/backend/font"
	_ "github.com/boxesandglue/ets/backend/image"
	_ "github.com/boxesandglue/ets/backend/lang"
	_ "github.com/boxesandglue/ets/backend/node"
	_ "github.com/boxesandglue/ets/frontend"
	_ "github.com/boxesandglue/ets/harfbuzz"
	_ "github.com/boxesandglue/ets/libbaseline"
)

var (
	version string
)

const (
	cmdRun     = "run"
	cmdHelp    = "help"
	cmdVersion = "version"
)

func dothings() error {
	pathToExefile, err := os.Executable()
	if err != nil {
		return err
	}
	exenameWithExt := filepath.Base(pathToExefile)
	var extension = filepath.Ext(exenameWithExt)
	var exename = exenameWithExt[0 : len(exenameWithExt)-len(extension)]
	var configFilesRead = []string{}

	cfgfilename := exename + ".cfg"
	cfg := core.Configuration{
		Version:   version,
		Starttime: time.Now(),
		Loglevel:  "message",
		Verbose:   false,
	}

	if data, err := os.ReadFile(cfgfilename); err == nil {
		if err = toml.Unmarshal(data, &cfg); err != nil {
			switch t := err.(type) {
			case *toml.DecodeError:
				fmt.Println(t.String())
			default:
				return err
			}
			return err
		}
		configFilesRead = append(configFilesRead, cfgfilename)
	}

	op := optionparser.NewOptionParser()
	op.Banner = "experimental typesetting system\nrun: ets somefile.js"
	op.Command(cmdVersion, "Show version information")
	op.Command(cmdHelp, "Show usage help")
	op.On("--loglevel LVL", "Set the log level for the console to one of debug, info, warn, error", &cfg.Loglevel)
	op.On("--verbose", "Print a bit of debugging output", &cfg.Verbose)

	err = op.Parse()
	if err != nil {
		return err
	}

	if len(op.Extra) == 0 {
		return fmt.Errorf("Please specify a command or file to run. See %s --help", exename)
	}
	switch op.Extra[0] {
	case cmdVersion:
		fmt.Printf("%s version %s\n", exename, version)
		os.Exit(0)
	case cmdHelp:
		op.Help()
		os.Exit(0)
	}

	core.SetCfg(cfg)
	core.SetupLog(exename + "-protocol.xml")
	for _, cfgFile := range configFilesRead {
		slog.Info("Use configuration file", "filename", cfgFile)
	}

	return core.RunETS(exename, op.Extra)
}

func main() {
	err := dothings()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
