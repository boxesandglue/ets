package core

import (
	"context"
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/boxesandglue/boxesandglue/backend/bag"
)

const (
	// LevelMessage is used for messages from Message
	LevelMessage = slog.Level(3)
)

var (
	lvl                = new(slog.LevelVar)
	logEncoder         *xml.Encoder
	msgAttr            = xml.Attr{Name: xml.Name{Local: "msg"}}
	lvlAttr            = xml.Attr{Name: xml.Name{Local: "level"}}
	logElement         = xml.StartElement{Name: xml.Name{Local: "entry"}}
	protocolWriter     io.Writer
	statusWriter       io.Writer
	errCount           = 0
	warnCount          = 0
	loglevel           slog.LevelVar
	repl               = strings.NewReplacer(" ", "-") // for XML attribute names
	logStartElement    xml.StartElement
	statusStartElement xml.StartElement
)

type logHandler struct {
}

func (lh *logHandler) Enabled(_ context.Context, level slog.Level) bool {
	if level == slog.LevelError {
		errCount++
	} else if level == slog.LevelWarn {
		warnCount++
	}
	return level >= loglevel.Level()
}

func (lh *logHandler) Handle(_ context.Context, r slog.Record) error {
	lvlAttr.Value = getLoglevelString(r.Level)
	msgAttr.Value = r.Message
	values := []string{}
	le := logElement.Copy()
	le.Attr = append(le.Attr, lvlAttr, msgAttr)
	r.Attrs(
		func(a slog.Attr) bool {
			var val string
			switch t := a.Value.Any().(type) {
			case slog.LogValuer:
				val = t.LogValue().String()
				values = append(values, fmt.Sprintf("%s=%s", a.Key, val))
			default:
				t = a.Value
				val = a.Value.String()
				values = append(values, fmt.Sprintf("%s=%s", a.Key, a.Value))
			}
			le.Attr = append(le.Attr, xml.Attr{Name: xml.Name{Local: repl.Replace(a.Key)}, Value: val})
			return true
		})
	logEncoder.EncodeToken(xml.CharData([]byte("  ")))
	logEncoder.EncodeToken(le)
	logEncoder.EncodeToken(le.End())
	logEncoder.EncodeToken(xml.CharData([]byte("\n")))

	if cfg.Verbose {
		lparen := ""
		rparen := ""
		if len(values) > 0 {
			lparen = "("
			rparen = ")"
		}
		lvlString := "·  "
		switch r.Level {
		case slog.LevelWarn:
			lvlString = "W: "
		case slog.LevelError:
			lvlString = "E: "
		}
		msg := r.Message
		if r.Level == LevelMessage {
			msg = "Message: " + msg
		}
		fmt.Println(lvlString, msg, lparen+strings.Join(values, ",")+rparen)
	}

	// status file
	eltName := xml.StartElement{}
	switch r.Level {
	case LevelMessage:
		eltName.Name = xml.Name{Local: "Message"}
	case slog.LevelWarn:
		eltName.Name = xml.Name{Local: "Warning"}
	case slog.LevelError:
		eltName.Name = xml.Name{Local: "Error"}
	}

	return nil
}

func (lh *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return lh
}

func (lh *logHandler) WithGroup(name string) slog.Handler {
	return lh
}

func getLoglevelString(lvl slog.Level) string {
	switch lvl {
	case -8:
		return "trace"
	case slog.LevelDebug:
		return "debug"
	case slog.LevelInfo:
		return "info"
	case LevelMessage:
		return "message"
	case slog.LevelWarn:
		return "warn"
	case slog.LevelError:
		return "error"
	}
	return ""
}

// SetupLog creates a protocol file. filename is the name of the XML file.
func SetupLog(filename string) error {
	cfg.protocolfilename = filename
	var err error
	switch cfg.Loglevel {
	case "trace":
		loglevel.Set(slog.Level(-8))
	case "debug":
		loglevel.Set(slog.LevelDebug)
	case "info":
		loglevel.Set(slog.LevelInfo)
	case "message":
		loglevel.Set(LevelMessage)
	case "warn":
		loglevel.Set(slog.LevelWarn)
	case "error":
		loglevel.Set(slog.LevelError)
	}
	protocolWriter, err = os.Create(filename)
	if err != nil {
		return err
	}
	sl := slog.New(&logHandler{})
	logEncoder = xml.NewEncoder(protocolWriter)
	logStartElement = xml.StartElement{
		Name: xml.Name{Local: "log"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "loglevel"}, Value: getLoglevelString(loglevel.Level())},
			{Name: xml.Name{Local: "time"}, Value: time.Now().Format(time.Stamp)},
			{Name: xml.Name{Local: "version"}, Value: cfg.Version},
		},
	}
	if err = logEncoder.EncodeToken(logStartElement); err != nil {
		return err
	}
	logEncoder.EncodeToken(xml.CharData([]byte("\n")))
	slog.SetDefault(sl)
	bag.SetLogger(sl)

	return nil
}

func teardownLog() error {
	summaryElement := xml.StartElement{
		Name: xml.Name{Local: "summary"},
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "errors"}, Value: fmt.Sprint(errCount)},
			{Name: xml.Name{Local: "warnings"}, Value: fmt.Sprint(warnCount)},
		},
	}
	var err error
	if err = logEncoder.EncodeToken(xml.CharData([]byte("  "))); err != nil {
		return err
	}
	if err = logEncoder.EncodeToken(summaryElement); err != nil {
		return err
	}
	if err = logEncoder.EncodeToken(summaryElement.End()); err != nil {
		return err
	}
	if err = logEncoder.EncodeToken(xml.CharData([]byte("\n"))); err != nil {
		return err
	}
	if err = logEncoder.EncodeToken(logStartElement.End()); err != nil {
		return err
	}
	if err := logEncoder.Flush(); err != nil {
		return err
	}

	return nil
}

// see the section on performance considerations in the slog package for a
// rationale why I chose this way to do the calculation.
type md5calc string

func (e md5calc) LogValue() slog.Value {
	data, err := os.ReadFile(string(e))
	if err != nil {
		return slog.StringValue(err.Error())
	}
	return slog.AnyValue(fmt.Sprintf("%x", md5.Sum(data)))
}
