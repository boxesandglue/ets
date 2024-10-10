package frontend

import (
	"github.com/boxesandglue/boxesandglue/frontend"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

const modulename = "bag:frontend"

func init() {
	require.RegisterNativeModule(modulename, Require)
}

type jsFrontend struct {
	runtime *goja.Runtime
}

func (jf *jsFrontend) frontendNew(call goja.FunctionCall) goja.Value {
	firstArg := call.Arguments[0]
	filename := firstArg.String()
	fd, err := frontend.New(filename)
	if err != nil {
		panic(err)
	}
	jf.runtime.Set("__filename", filename)
	jf.runtime.Set("__pw", fd.Doc.PDFWriter)
	return jf.runtime.ToValue(fd)
}

// Require is called on load.
func Require(runtime *goja.Runtime, module *goja.Object) {
	jf := &jsFrontend{
		runtime: runtime,
	}

	o := module.Get("exports").(*goja.Object)
	o.Set("new", jf.frontendNew)
	o.Set("getLanguage", frontend.GetLanguage)
	o.Set("fontSource", func(call goja.FunctionCall) goja.Value {
		firstArg := call.Arguments[0]
		return runtime.ToValue(&frontend.FontSource{Location: firstArg.String()})
	})
	o.Set("newText", frontend.NewText)
	o.Set("leading", frontend.Leading)
	o.Set("fontsize", frontend.FontSize)
	o.Set("family", frontend.Family)
	o.Set("fontStyleNormal", frontend.FontStyleNormal)
	o.Set("fontStyleItalic", frontend.FontStyleItalic)
	o.Set("fontStyleOblique", frontend.FontStyleOblique)
	o.Set("fontWeight100", frontend.FontWeight100)
	o.Set("fontWeight200", frontend.FontWeight200)
	o.Set("fontWeight300", frontend.FontWeight300)
	o.Set("fontWeight400", frontend.FontWeight400)
	o.Set("fontWeight500", frontend.FontWeight500)
	o.Set("fontWeight600", frontend.FontWeight600)
	o.Set("fontWeight700", frontend.FontWeight700)
	o.Set("fontWeight800", frontend.FontWeight800)
	o.Set("fontWeight900", frontend.FontWeight900)
	o.Set("textDecorationUnderline", frontend.TextDecorationUnderline)
	o.Set("settingTextDecorationLine", runtime.ToValue(frontend.SettingTextDecorationLine).ToNumber())
	o.Set("settingBox", runtime.ToValue(frontend.SettingBox).ToNumber())
	o.Set("settingBackgroundColor", runtime.ToValue(frontend.SettingBackgroundColor).ToNumber())
	o.Set("settingBorderBottomWidth", runtime.ToValue(frontend.SettingBorderBottomWidth).ToNumber())
	o.Set("settingBorderLeftWidth", runtime.ToValue(frontend.SettingBorderLeftWidth).ToNumber())
	o.Set("settingBorderRightWidth", runtime.ToValue(frontend.SettingBorderRightWidth).ToNumber())
	o.Set("settingBorderTopWidth", runtime.ToValue(frontend.SettingBorderTopWidth).ToNumber())
	o.Set("settingBorderBottomColor", runtime.ToValue(frontend.SettingBorderBottomColor).ToNumber())
	o.Set("settingBorderLeftColor", runtime.ToValue(frontend.SettingBorderLeftColor).ToNumber())
	o.Set("settingBorderRightColor", runtime.ToValue(frontend.SettingBorderRightColor).ToNumber())
	o.Set("settingBorderTopColor", runtime.ToValue(frontend.SettingBorderTopColor).ToNumber())
	o.Set("settingBorderBottomStyle", runtime.ToValue(frontend.SettingBorderBottomStyle).ToNumber())
	o.Set("settingBorderLeftStyle", runtime.ToValue(frontend.SettingBorderLeftStyle).ToNumber())
	o.Set("settingBorderRightStyle", runtime.ToValue(frontend.SettingBorderRightStyle).ToNumber())
	o.Set("settingBorderTopStyle", runtime.ToValue(frontend.SettingBorderTopStyle).ToNumber())
	o.Set("settingBorderTopLeftRadius", runtime.ToValue(frontend.SettingBorderTopLeftRadius).ToNumber())
	o.Set("settingBorderTopRightRadius", runtime.ToValue(frontend.SettingBorderTopRightRadius).ToNumber())
	o.Set("settingBorderBottomLeftRadius", runtime.ToValue(frontend.SettingBorderBottomLeftRadius).ToNumber())
	o.Set("settingBorderBottomRightRadius", runtime.ToValue(frontend.SettingBorderBottomRightRadius).ToNumber())
	o.Set("settingColor", runtime.ToValue(frontend.SettingColor).ToNumber())
	o.Set("settingDebug", runtime.ToValue(frontend.SettingDebug).ToNumber())
	o.Set("settingFontExpansion", runtime.ToValue(frontend.SettingFontExpansion).ToNumber())
	o.Set("settingFontFamily", runtime.ToValue(frontend.SettingFontFamily).ToNumber())
	o.Set("settingFontWeight", runtime.ToValue(frontend.SettingFontWeight).ToNumber())
	o.Set("settingHAlign", runtime.ToValue(frontend.SettingHAlign).ToNumber())
	o.Set("settingHangingPunctuation", runtime.ToValue(frontend.SettingHangingPunctuation).ToNumber())
	o.Set("settingHeight", runtime.ToValue(frontend.SettingHeight).ToNumber())
	o.Set("settingHyperlink", runtime.ToValue(frontend.SettingHyperlink).ToNumber())
	o.Set("settingIndentLeft", runtime.ToValue(frontend.SettingIndentLeft).ToNumber())
	o.Set("settingIndentLeftRows", runtime.ToValue(frontend.SettingIndentLeftRows).ToNumber())
	o.Set("settingLeading", runtime.ToValue(frontend.SettingLeading).ToNumber())
	o.Set("settingMarginBottom", runtime.ToValue(frontend.SettingMarginBottom).ToNumber())
	o.Set("settingMarginLeft", runtime.ToValue(frontend.SettingMarginLeft).ToNumber())
	o.Set("settingMarginRight", runtime.ToValue(frontend.SettingMarginRight).ToNumber())
	o.Set("settingMarginTop", runtime.ToValue(frontend.SettingMarginTop).ToNumber())
	o.Set("settingOpenTypeFeature", runtime.ToValue(frontend.SettingOpenTypeFeature).ToNumber())
	o.Set("settingPaddingBottom", runtime.ToValue(frontend.SettingPaddingBottom).ToNumber())
	o.Set("settingPaddingLeft", runtime.ToValue(frontend.SettingPaddingLeft).ToNumber())
	o.Set("settingPaddingRight", runtime.ToValue(frontend.SettingPaddingRight).ToNumber())
	o.Set("settingPaddingTop", runtime.ToValue(frontend.SettingPaddingTop).ToNumber())
	o.Set("settingPrepend", runtime.ToValue(frontend.SettingPrepend).ToNumber())
	o.Set("settingPreserveWhitespace", runtime.ToValue(frontend.SettingPreserveWhitespace).ToNumber())
	o.Set("settingSize", runtime.ToValue(frontend.SettingSize).ToNumber())
	o.Set("settingStyle", runtime.ToValue(frontend.SettingStyle).ToNumber())
	o.Set("settingTabSizeSpaces", runtime.ToValue(frontend.SettingTabSizeSpaces).ToNumber())
	o.Set("settingTabSize", runtime.ToValue(frontend.SettingTabSize).ToNumber())
	o.Set("settingTextDecorationLine", runtime.ToValue(frontend.SettingTextDecorationLine).ToNumber())
	o.Set("settingWidth", runtime.ToValue(frontend.SettingWidth).ToNumber())
	o.Set("settingVAlign", runtime.ToValue(frontend.SettingVAlign).ToNumber())
	o.Set("settingYOffset", runtime.ToValue(frontend.SettingYOffset).ToNumber())
}
