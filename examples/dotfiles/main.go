package main

import (
	"fmt"
	"os"

	vi "github.com/greeschenko/vim9gorn"
)

func main() {
	g := vi.New()

	addHeader(g)
	addVariables(g)
	addHelperFunctions(g)
	addPlugins(g)
	addUserCommands(g)
	addOptions(g)
	addColorscheme(g)
	addHighlights(g)
	addKeymaps(g)
	addAutoCommands(g)
	addLspConfig(g)
	addAutoLoad(g)

	output := g.Generate()
	if err := os.WriteFile("generated.vimrc", []byte(output), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Generated: generated.vimrc")
	fmt.Printf("Size: %d bytes\n", len(output))
}

func addHeader(g *vi.Gorn) {
}

func addVariables(g *vi.Gorn) {
	g.AddSection(
		vi.NewVariables().
			Var("plugin_dir", `expand('~/.vim/pack/plugins/opt')`),
	)

	g.AddSection(
		vi.NewVariables().
			Var("leader", `"<Space>"`).
			Var("localleader", `"','"`),
	)
}

func addHelperFunctions(g *vi.Gorn) {
	checkGit := vi.NewFunction("CheckGit").
		SetScope(vi.Global).
		Add(vi.Raw{Code: `if !executable('git') | echoerr 'Git is required' | endif`})

	ensureDir := vi.NewFunction("EnsureDir").
		SetScope(vi.Global).
		Arg("dir", "string").
		Add(vi.Raw{Code: `if !isdirectory(dir) | call mkdir(dir, 'p') | endif`})

	installPlugin := vi.NewFunction("InstallPlugin").
		SetScope(vi.Global).
		Arg("repo", "string").
		Add(vi.Raw{Code: `var name = fnamemodify(repo, ':t')`}).
		Add(vi.Raw{Code: `var path = plugin_dir .. '/' .. name`}).
		Add(vi.NewIfElse("!isdirectory(path)").
			ThenAdd(vi.Raw{Code: `execute '!git clone --depth 1 https://github.com/' .. repo .. '.git ' .. path`}))

	g.AddSection(checkGit)
	g.AddSection(ensureDir)
	g.AddSection(installPlugin)
}

func addPlugins(g *vi.Gorn) {
	plugins := vi.NewVariables().
		Var("g:plugins", `["tpope/vim-sensible", "junegunn/fzf.vim", "neoclide/coc.nvim"]`)

	setupPlugins := vi.NewFunction("SetupPlugins").
		SetScope(vi.Global).
		Add(vi.Raw{Code: "call CheckGit()"}).
		Add(vi.Raw{Code: "call EnsureDir(plugin_dir)"}).
		Add(vi.Raw{Code: "for repo in g:plugins | call InstallPlugin(repo) | endfor"})

	g.AddSection(plugins)
	g.AddSection(setupPlugins)

	g.AddSection(vi.Raw{Code: "call SetupPlugins()"})
}

func addUserCommands(g *vi.Gorn) {
	g.AddSection(&vi.Command{Name: "PluginsInstall", Cmd: "call.SetupPlugins()"})
}

func addOptions(g *vi.Gorn) {
	g.AddSection(
		vi.NewOptions().
			Set("number", true).
			Set("relativenumber", true).
			Set("tabstop", 4).
			Set("shiftwidth", 4).
			Set("expandtab", true).
			Set("smartindent", true).
			Set("autoindent", true).
			Set("wrap", false).
			Set("hidden", true).
			Set("wildmenu", true).
			Set("incsearch", true).
			Set("hlsearch", true).
			Set("clipboard", "unnamedplus").
			Set("encoding", "utf-8").
			Set("termencoding", "utf-8").
			Set("background", "dark").
			Set("laststatus", 2).
			Set("showmode", false).
			Set("showcmd", true).
			Set("cursorline", true).
			Set("signcolumn", "yes").
			Set("timeoutlen", 300).
			Set("updatetime", 300),
	)
}

func addColorscheme(g *vi.Gorn) {
	g.AddSection(vi.ColorScheme{
		Background:    "dark",
		TermGuiColors: true,
		SyntaxEnable:  true,
		Name:          "gruvbox",
	})
}

func addHighlights(g *vi.Gorn) {
	g.AddSection(vi.Highlight{LinkFrom: "Extra", LinkTo: "Comment"})
	g.AddSection(vi.Highlight{LinkFrom: "Todo", LinkTo: "WarningMsg"})
}

func addKeymaps(g *vi.Gorn) {
	g.AddSection(
		vi.NewVariables().
			Legacy(vi.Global, "mapleader", `"<Space>"`).
			Legacy(vi.Global, "maplocalleader", "','"),
	)

	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>w", RHS: ":w<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>q", RHS: ":q<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>h", RHS: "<C-w>h"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>j", RHS: "<C-w>j"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>k", RHS: "<C-w>k"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>l", RHS: "<C-w>l"})

	g.AddSection(vi.Keymap{Mode: "v", LHS: "<", RHS: "<gv"})
	g.AddSection(vi.Keymap{Mode: "v", LHS: ">", RHS: ">gv"})

	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>p", RHS: `"_dP`})
}

func addAutoCommands(g *vi.Gorn) {
	ag := vi.NewAutocmdGroup("MySettings").
		Add(*vi.NewAutocmd("FileType", "go").SetCmd("setlocal shiftwidth=4 noexpandtab")).
		Add(*vi.NewAutocmd("FileType", "python").SetCmd("setlocal shiftwidth=4 expandtab"))

	g.AddSection(ag)
}

func addLspConfig(g *vi.Gorn) {
	g.AddSection(vi.Raw{Code: `let g:lsp_semantic_enabled = v:true`})
	g.AddSection(vi.Raw{Code: `let g:lsp_diagnostics_enabled = v:true`})
}

func addAutoLoad(g *vi.Gorn) {
	af := vi.NewAutoloadFile("utils").
		AddFunc(*vi.NewAutoloadFunc("utils", "Hello").
			Add(vi.Raw{Code: "echo 'Hello from autoload!'"})).
		AddFunc(*vi.NewAutoloadFunc("utils", "Echo").
			Arg("msg", "string").
			Add(vi.Raw{Code: "echo msg"}))

	g.AddSection(af)
}
