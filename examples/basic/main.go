package main

import (
	vi "github.com/greeschenko/vim9gorn"
)

func main() {
	v := vi.New()

	// =====================
	// Options
	// =====================
	v.AddSection(
		vi.NewOptions().
			Set("wildmenu", true).
			Set("number", true).
			Set("relativenumber", true).
			Set("tabstop", 4),
	)

	// =====================
	// Variables (legacy scopes)
	// =====================
	v.AddSection(
		vi.NewVariables().
			Legacy(vi.Global, "mapleader", `"<Space>"`).
			Legacy(vi.Global, "maplocalleader", "','"),
	)

	// =====================
	// Colorscheme
	// =====================
	v.AddSection(vi.ColorScheme{
		Background:    "dark",
		TermGuiColors: true,
		SyntaxEnable:  true,
		Name:          "retrobox",
	})

	// =====================
	// Highlights
	// =====================
	v.AddSection(vi.Highlight{LinkFrom: "Extra", LinkTo: "Comment"})
	v.AddSection(vi.Highlight{
		Group: "Git",
		Args:  "guibg=#F34F29 guifg=#FFFFFF ctermbg=202 ctermfg=231",
	})

	// =====================
	// Keymaps
	// =====================
	v.AddSection(vi.Keymap{
		Mode: "n", LHS: "<leader>vl", RHS: "^v$h",
	})
	v.AddSection(vi.Keymap{
		Mode: "n", LHS: "<esc>", RHS: ":nohlsearch<return>", Silent: true,
	})
	v.AddSection(vi.Keymap{Mode: "v", LHS: "<", RHS: "<gv"})
	v.AddSection(vi.Keymap{Mode: "v", LHS: ">", RHS: ">gv"})

	// =====================
	// Function (Vim9)
	// =====================
	greet := vi.NewFunction("Greet").
		SetScope(vi.Global)

	// ---------------------
	// if / elseif / else
	// ---------------------
	timeCheck := vi.NewIfElse(`str2nr(strftime("%H")) < 12`).
		ThenAdd(vi.Raw{Code: `echo "Good morning from vim9gorn 👋"`}).
		ElseIfAdd(`str2nr(strftime("%H")) < 18`,
			vi.Raw{Code: `echo "Good afternoon from vim9gorn 👋"`},
		).
		ElseAdd(vi.Raw{Code: `echo "Good evening from vim9gorn 👋"`})

	greet.Add(timeCheck)

	// ---------------------
	// for loop with continue & break using List
	// ---------------------
	numbers := vi.NewList("1", "2", "3", "4", "5")

	loop := vi.NewForLoop("_", "i", numbers.Generate()).
		Add(
			vi.NewIfElse("i == 2").
				ThenAdd(vi.Raw{Code: "continue"}),
		).
		Add(
			vi.NewIfElse("i == 4").
				ThenAdd(vi.Raw{Code: "break"}),
		).
		Add(
			vi.Raw{Code: `echo "Loop value: " .. i`},
		)

	greet.Add(loop)

	// ---------------------
	// for loop over Dict keys/values
	// ---------------------
	myDict := vi.NewDict().
		Set("a", `"Apple"`).
		Set("b", `"Banana"`).
		Set("c", `"Cherry"`)

	dictLoop := vi.NewForLoop("k", "v", vi.Items(myDict)).
		Add(vi.Raw{Code: `echo k .. ": " .. v`})

	greet.Add(dictLoop)

	// ---------------------
	// register function
	// ---------------------
	v.AddSection(greet)

	// =====================
	// Forge
	// =====================
	writer := &vi.DefaultFileWriter{}
	if err := v.Forge(".vimrc", writer); err != nil {
		panic(err)
	}
}
