package vim9gorn

import (
	"strings"
)

type ColorScheme struct {
	Background    string
	TermGuiColors bool
	SyntaxEnable  bool
	Name          string
}

func (c ColorScheme) Generate() string {
	var b strings.Builder
	b.WriteString("# === Colorscheme ===\n")
	if c.Background != "" {
		b.WriteString("set background=" + c.Background + "\n")
	}
	if c.TermGuiColors {
		b.WriteString("if has(\"termguicolors\")\n  set termguicolors\nendif\n")
	}
	if c.SyntaxEnable {
		b.WriteString("syntax enable\n")
	}
	if c.Name != "" {
		b.WriteString("colorscheme " + c.Name + "\n")
	}
	return b.String()
}
