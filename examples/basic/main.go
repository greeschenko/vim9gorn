package main

import (
	"github.com/greeschenko/vim9gorn"
)

func main() {
	v := vim9gorn.New()

	// =====================
	// Options
	// =====================
	v.Options.Set("wildmenu", true)
	v.Options.Set("number", true)
	v.Options.Set("relativenumber", true)
	v.Options.Set("tabstop", 4)

	// =====================
	// Variables (legacy scopes)
	// =====================
	v.Variables.Legacy(vim9gorn.Global, "mapleader", "\"\\<Space>\"")
	v.Variables.Legacy(vim9gorn.Global, "maplocalleader", "','")

	// =====================
	// Colorscheme
	// =====================
	v.AddSection(vim9gorn.ColorScheme{
		Background:    "dark",
		TermGuiColors: true,
		SyntaxEnable:  true,
		Name:          "retrobox",
	})

	// =====================
	// Highlights
	// =====================
	v.AddSection(vim9gorn.Highlight{LinkFrom: "Extra", LinkTo: "Comment"})
	v.AddSection(vim9gorn.Highlight{
		Group: "Git",
		Args:  "guibg=#F34F29 guifg=#FFFFFF ctermbg=202 ctermfg=231",
	})

	// =====================
	// Keymaps
	// =====================
	v.AddSection(vim9gorn.Keymap{
		Mode: "n", LHS: "<leader>vl", RHS: "^v$h",
	})
	v.AddSection(vim9gorn.Keymap{
		Mode: "n", LHS: "<esc>", RHS: ":nohlsearch<return>", Silent: true,
	})
	v.AddSection(vim9gorn.Keymap{Mode: "v", LHS: "<", RHS: "<gv"})
	v.AddSection(vim9gorn.Keymap{Mode: "v", LHS: ">", RHS: ">gv"})

	// =====================
	// Function (Vim9)
	// =====================
	greet := vim9gorn.NewFunction("Greet").
		SetScope(vim9gorn.Global).
		Add(vim9gorn.Raw{
			Code: `echo "Hello from vim9gorn 👋"`,
		})

	v.AddSection(greet)

	// =====================
	// Forge
	// =====================
	writer := &vim9gorn.DefaultFileWriter{}
	if err := v.Forge(".vimrc", writer); err != nil {
		panic(err)
	}
}
