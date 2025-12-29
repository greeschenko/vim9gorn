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
	v.Variables.Legacy(vim9gorn.Global, "mapleader", `"<Space>"`)
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
		SetScope(vim9gorn.Global)

	// ---------------------
	// if / elseif / else
	// ---------------------
	timeCheck := vim9gorn.NewIfElse(`str2nr(strftime("%H")) < 12`).
		ThenAdd(vim9gorn.Raw{Code: `echo "Good morning from vim9gorn 👋"`}).
		ElseIfAdd(`str2nr(strftime("%H")) < 18`,
			vim9gorn.Raw{Code: `echo "Good afternoon from vim9gorn 👋"`},
		).
		ElseAdd(vim9gorn.Raw{Code: `echo "Good evening from vim9gorn 👋"`})

	greet.Add(timeCheck)

	// ---------------------
	// for loop with continue & break
	// ---------------------
	loop := vim9gorn.NewForLoop("i", "[1, 2, 3, 4, 5]").
		Add(
			vim9gorn.NewIfElse("i == 2").
				ThenAdd(vim9gorn.Raw{Code: "continue"}),
		).
		Add(
			vim9gorn.NewIfElse("i == 4").
				ThenAdd(vim9gorn.Raw{Code: "break"}),
		).
		Add(
			vim9gorn.Raw{Code: `echo "Loop value: " .. i`},
		)

	greet.Add(loop)

	// ---------------------
	// register function
	// ---------------------
	v.AddSection(greet)

	// =====================
	// Forge
	// =====================
	writer := &vim9gorn.DefaultFileWriter{}
	if err := v.Forge(".vimrc", writer); err != nil {
		panic(err)
	}
}
