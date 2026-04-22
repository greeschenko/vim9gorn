package vim9gorn_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	vi "github.com/greeschenko/vim9gorn"
)

func TestGenerateUserVimrc(t *testing.T) {
	g := vi.New()

	// ============================================================
	// VARIABLES
	// ============================================================
	g.AddSection(
		vi.NewVariables().
			Var("plugin_dir", `expand('~/.vim/pack/plugins/opt')`),
	)

	// LSP servers - as raw vim code (complex dict)
	g.AddSection(vi.Raw{Code: `var lspServers = [
    {
        name: 'typescriptlang',
        filetype: ['javascript', 'typescript'],
        path: 'typescript-language-server',
        args: ['--stdio'],
    },
    {
        filetype: ['go', 'gomod'],
        path: '/home/olex/prodev/go/bin/gopls',
        args: [],
    }
]`})

	// -----------------------------
	// HELPER FUNCTIONS
	// -----------------------------

	checkGit := vi.NewFunction("CheckGit").
		SetScope(vi.Global).
		Add(vi.Raw{Code: `if !executable('git') | echoerr 'Git is required to install plugins' | endif`})

	ensurePluginDir := vi.NewFunction("EnsurePluginDir").
		SetScope(vi.Global).
		Add(vi.Raw{Code: `if !isdirectory(plugin_dir) | mkdir(plugin_dir, 'p') | endif`})

	pluginName := vi.NewFunction("PluginName").
		SetScope(vi.Global).
		Arg("repo", "string").
		Returns("string").
		Add(vi.Raw{Code: `return fnamemodify(repo, ':t')`})

	pluginPath := vi.NewFunction("PluginPath").
		SetScope(vi.Global).
		Arg("repo", "string").
		Returns("string").
		Add(vi.Raw{Code: `return plugin_dir .. '/' .. PluginName(repo)`})

	installPlugin := vi.NewFunction("InstallPlugin").
		SetScope(vi.Global).
		Arg("repo", "string").
		Add(vi.Raw{Code: `var path = PluginPath(repo)`}).
		Add(vi.NewIfElse("!isdirectory(path)").
			ThenAdd(vi.Raw{
				Code: `var url = 'https://github.com/' .. repo .. '.git' | echom 'Installing ' .. repo | system('git clone --depth 1 ' .. shellescape(url) .. ' ' .. shellescape(path))`,
			}))

	updatePlugin := vi.NewFunction("UpdatePlugin").
		SetScope(vi.Global).
		Arg("repo", "string").
		Add(vi.Raw{Code: `var path = PluginPath(repo)`}).
		Add(vi.NewIfElse("isdirectory(path)").
			ThenAdd(vi.Raw{
				Code: `echom 'Updating ' .. repo | system('git -C ' .. shellescape(path) .. ' pull --ff-only')`,
			}))

	loadPlugin := vi.NewFunction("LoadPlugin").
		SetScope(vi.Global).
		Arg("repo", "string").
		Add(vi.Raw{Code: `var name = PluginName(repo) | execute 'packadd ' .. name`})

	g.AddSection(checkGit)
	g.AddSection(ensurePluginDir)
	g.AddSection(pluginName)
	g.AddSection(pluginPath)
	g.AddSection(installPlugin)
	g.AddSection(updatePlugin)
	g.AddSection(loadPlugin)

	// -----------------------------
	// PLUGINS
	// -----------------------------

	// Load dev plugin from local source
	g.AddSection(
		vi.NewVariables().
			Legacy(vi.Global, "vimsidian_vault_path", `"~/prodev/MIND_VAULT"`),
	)

	// Plugin list - strings must be quoted in vim9script
	pluginList := vi.NewVariables().
		Var("plugin_list", `["greeschenko/cyberpunk99.vim", "yegappan/lsp", "liuchengxu/vim-which-key", "mg979/vim-visual-multi", "SirVer/ultisnips", "honza/vim-snippets", "vim-fuzzbox/fuzzbox.vim", "greeschenko/vim9-ollama"]`)

	// SetupPlugins function
	setupPlugins := vi.NewFunction("SetupPlugins").
		SetScope(vi.Global).
		Add(vi.Raw{Code: "CheckGit()"}).
		Add(vi.Raw{Code: "EnsurePluginDir()"}).
		Add(vi.Raw{Code: "for repo in plugin_list | InstallPlugin(repo) | LoadPlugin(repo) | endfor"})

	// UpdatePlugins function
	updatePlugins := vi.NewFunction("UpdatePlugins").
		SetScope(vi.Global).
		Add(vi.Raw{Code: "CheckGit()"}).
		Add(vi.Raw{Code: "for repo in plugin_list | UpdatePlugin(repo) | endfor"})

	g.AddSection(pluginList)
	g.AddSection(setupPlugins)
	g.AddSection(updatePlugins)

	// -----------------------------
	// USER COMMANDS
	// -----------------------------

	g.AddSection(&vi.Command{Name: "PluginsInstall", Cmd: "call.SetupPlugins()"})
	g.AddSection(&vi.Command{Name: "PluginsUpdate", Cmd: "call.UpdatePlugins()"})

	// -----------------------------
	// OPTIONS
	// -----------------------------

	g.AddSection(
		vi.NewOptions().
			Set("number", true).
			Set("wrap", false).
			Set("wildmenu", true).
			Set("autoread", true).
			Set("nobackup", true).
			Set("noswapfile", true).
			Set("termencoding", "utf-8").
			Set("encoding", "utf-8").
			Set("clipboard", "unnamed,unnamedplus").
			Set("incsearch", true).
			Set("hlsearch", true).
			Set("expandtab", true).
			Set("shiftwidth", 4).
			Set("softtabstop", 4).
			Set("tabstop", 4).
			Set("ttyfast", true).
			Set("laststatus", 2).
			Set("background", "dark").
			Set("termguicolors", true).
			Set("timeoutlen", 100).
			Set("smoothscroll", true).
			Set("smartindent", true).
			Set("autoindent", true).
			Set("nocursorcolumn", true).
			Set("nocursorline", true).
			Set("hidden", true).
			Set("maxmempattern", 5000),
	)

	// -----------------------------
	// COLORSCHEME
	// -----------------------------

	g.AddSection(vi.ColorScheme{
		Background:    "dark",
		TermGuiColors: true,
		SyntaxEnable:  true,
		Name:          "cyberpunk99",
	})

	// -----------------------------
	// HIGHLIGHTS
	// -----------------------------

	g.AddSection(vi.Highlight{LinkFrom: "Extra", LinkTo: "Comment"})
	g.AddSection(vi.Highlight{
		Group: "Git",
		Args:  "guibg=#F34F29 guifg=#FFFFFF ctermbg=202 ctermfg=231",
	})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagVirtualTextWarning", LinkFrom: "WarningMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagVirtualTextError", LinkFrom: "WarningMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagVirtualTextInfo", LinkFrom: "Comment"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagVirtualTextHint", LinkFrom: "MoreMsg"})

	g.AddSection(vi.Highlight{LinkTo: "LspDiagSignErrorText", LinkFrom: "WarningMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagSignWarningText", LinkFrom: "MoreMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagSignInfoText", LinkFrom: "Comment"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagSignHintText", LinkFrom: "MoreMsg"})

	g.AddSection(vi.Highlight{LinkTo: "LspDiagInlineError", LinkFrom: "WarningMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagInlineWarning", LinkFrom: "MoreMsg"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagInlineInfo", LinkFrom: "Comment"})
	g.AddSection(vi.Highlight{LinkTo: "LspDiagInlineHint", LinkFrom: "NONE"})

	// -----------------------------
	// KEYMAPS
	// -----------------------------

	g.AddSection(
		vi.NewVariables().
			Legacy(vi.Global, "mapleader", `"\<Space>"`).
			Legacy(vi.Global, "maplocalleader", "','"),
	)

	// which_key maps via Raw (external plugin)
	g.AddSection(vi.Raw{Code: `g:which_key_map = {}`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.f = { 'name': '+FZF' }`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.c = { 'name': '+COMMENT' }`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.g = { 'name': '+GIT' }`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.l = { 'name': '+LSP'}`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.l.e = { 'name': '+Diag'}`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.s = { 'name': '+STARGATE' }`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.n = { 'name': '+TREE' }`})
	g.AddSection(vi.Raw{Code: `g:which_key_map.v = { 'name': '+SELECT' }`})

	// WhichKey registration
	g.AddSection(vi.Raw{Code: `which_key#register('<Space>', "g:which_key_map", "n")`})
	g.AddSection(vi.Raw{Code: `which_key#register('<Space>', "g:which_key_map_visual", "v")`})

	// Normal mode keymaps
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>", RHS: `:<c-u>WhichKey '<Space>'<CR>`, Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<C-h>", RHS: "<C-w>h"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<C-l>", RHS: "<C-w>l"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<C-j>", RHS: "<C-w>j"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<C-k>", RHS: "<C-w>k"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>vl", RHS: "^v$h"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<esc>", RHS: ":nohlsearch<return>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>q", RHS: "ZZ"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<C-s>", RHS: ":w<cr>"})

	// Buffer navigation
	g.AddSection(vi.Keymap{Mode: "n", LHS: ">", RHS: ":bn<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<", RHS: ":bp<CR>"})

	// Visual mode keymaps
	g.AddSection(vi.Keymap{Mode: "v", LHS: "<leader>", RHS: `:<c-u>WhichKeyVisual '<Space>'<CR>`, Silent: true})
	g.AddSection(vi.Keymap{Mode: "v", LHS: "<", RHS: "<gv"})
	g.AddSection(vi.Keymap{Mode: "v", LHS: ">", RHS: ">gv"})

	// FZF mappings
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>/", RHS: ":FuzzyInBuffer<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>b", RHS: ":FuzzyBuffers<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>;", RHS: ":FuzzyRegisters<CR>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>.", RHS: ":FuzzyFiles<cr>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>fo", RHS: ":FuzzyMru<cr>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>fg", RHS: ":FuzzyGrep<cr>"})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>fm", RHS: ":FuzzyMarks<cr>"})

	// LSP mappings
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lD", RHS: ":LspGotoDeclaration<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>ld", RHS: ":LspGotoDefinition<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>li", RHS: ":LspGotoImpl<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lt", RHS: ":LspGotoTypeDef<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lr", RHS: ":LspShowReferences<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>la", RHS: ":LspCodeAction<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>ln", RHS: ":LspRename<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>K", RHS: ":LspHover<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>len", RHS: ":LspDiagNext<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lep", RHS: ":LspDiagPrev<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lel", RHS: ":LspDiagShow<CR>", Silent: true})
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>lf", RHS: ":LspFormat<CR>", Silent: true})

	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>e", RHS: ":Explore<CR>"})

	// -----------------------------
	// AUTOINSTALL
	// -----------------------------

	autoInstall := vi.NewFunction("Main").
		SetScope(vi.Global).
		Add(vi.Raw{Code: "SetupPlugins()"})

	g.AddSection(autoInstall)

	// -----------------------------
	// LSP CONFIG (external)
	// -----------------------------

	lspConfig := vi.NewFunction("LspConfig").
		SetScope(vi.Global).
		Add(vi.Raw{
			Code: `call LspOptionsSet({ aleSupport: v:false, autoComplete: v:true, autoHighlight: v:false, autoHighlightDiags: v:true, autoPopulateDiags: v:false, completionMatcher: 'case', completionMatcherValue: 1, diagSignErrorText: '●', diagSignHintText: '●', diagSignInfoText: '●', diagSignWarningText: '●', echoSignature: v:false, hideDisabledCodeActions: v:false, highlightDiagInline: v:true, hoverInPreview: v:false, completionInPreview: v:false, closePreviewOnComplete: v:true, ignoreMissingServer: v:false, keepFocusInDiags: v:true, keepFocusInReferences: v:true, completionTextEdit: v:true, diagVirtualTextAlign: 'after', diagVirtualTextWrap: 'default', noNewlineInCompletion: v:false, omniComplete: v:null, omniCompleteAllowBare: v:false, outlineOnRight: v:false, outlineWinSize: 20, popupBorder: v:true, popupBorderHighlight: 'Title', popupBorderHighlightPeek: 'Special', popupBorderSignatureHelp: v:false, popupHighlightSignatureHelp: 'Pmenu', popupHighlight: 'Normal', semanticHighlight: v:true, showDiagInBalloon: v:true, showDiagInPopup: v:true, showDiagOnStatusLine: v:false, showDiagWithSign: v:true, showDiagWithVirtualText: v:true, showInlayHints: v:false, showSignature: v:true, snippetSupport: v:false, ultisnipsSupport: v:false, useBufferCompletion: v:false, usePopupInCodeAction: v:false, useQuickfixForLocations: v:false, vsnipSupport: v:false, bufferCompletionTimeout: 100, customCompletionKinds: v:false, completionKinds: {}, filterCompletionDuplicates: v:false, condensedCompletionMenu: v:false })`,
		})

	lspServersAdd := vi.NewFunction("LspServersAdd").
		SetScope(vi.Global).
		Add(vi.Raw{Code: "call LspAddServer(lspServers)"})

	g.AddSection(lspConfig)
	g.AddSection(lspServersAdd)

	// Generate output
	output := g.Generate()

	// Write to test file
	writer := &vi.DefaultFileWriter{}
	if err := g.Forge("testdata/output.vimrc", writer); err != nil {
		t.Fatalf("Failed to write vimrc: %v", err)
	}

	// Verify output is not empty
	if output == "" {
		t.Error("Generated output is empty")
	}

	// Check for critical sections
	checks := []string{
		"vim9script",
		"def g:CheckGit",
		"def g:EnsurePluginDir",
		"def g:InstallPlugin",
		"def g:UpdatePlugin",
		"def g:LoadPlugin",
		"def g:SetupPlugins",
		"def g:UpdatePlugins",
		"command PluginsInstall",
		"command PluginsUpdate",
		"set number",
		"set tabstop=4",
		"set background=dark",
		"nmap",
		"vmap",
	}

	for _, check := range checks {
		if !strings.Contains(output, check) {
			t.Errorf("Missing section in generated vimrc: %s", check)
		}
	}

	t.Logf("Generated vimrc length: %d bytes", len(output))
}

func TestValidateGeneratedVimrcWithVim(t *testing.T) {
	// Skip if vim is not installed
	_, err := exec.LookPath("vim")
	if err != nil {
		t.Skip("vim not installed, skipping validation test")
	}

	// Read generated vimrc
	content, err := os.ReadFile("testdata/output.vimrc")
	if err != nil {
		t.Fatalf("Failed to read generated vimrc: %v", err)
	}

	// Split into lines and filter out plugin-dependent code
	// (which_key, packadd, etc. require plugins to be loaded)
	var filteredLines []string
	skipPatterns := []string{
		"which_key#",
		"packadd",
		"LspOptionsSet",
		"LspAddServer",
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		skip := false
		for _, pattern := range skipPatterns {
			if strings.Contains(line, pattern) {
				skip = true
				break
			}
		}
		if !skip {
			filteredLines = append(filteredLines, line)
		}
	}

	// Write filtered vimrc to temp file
	filteredContent := strings.Join(filteredLines, "\n")
	tmpFile, err := os.CreateTemp("", "vimrc-*.vim")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(filteredContent); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}
	tmpFile.Close()

	// Run vim in silent ex mode to validate
	cmd := exec.Command("vim", "-es", "-u", "NONE", "-c", "source "+tmpFile.Name(), "-c", "q")
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("Vim failed to parse generated vimrc: %v\nOutput: %s", err, output)
	}

	t.Log("Vim validation passed (filtered for plugin-dependent code)")
}
