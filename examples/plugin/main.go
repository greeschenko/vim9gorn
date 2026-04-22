package main

import (
	"fmt"
	"os"
	"path/filepath"

	vi "github.com/greeschenko/vim9gorn"
)

const pluginName = "myvimlib"
const pluginDesc = "My Vim Library"
const pluginURL = "https://github.com/user/myvimlib"

func main() {
	baseDir := filepath.Join(".", pluginName)

	createDirectories(baseDir)
	genPlugin(baseDir)
	genAutoload(baseDir)
	genDoc(baseDir)
	genReadme(baseDir)

	fmt.Println("Plugin generated!")
}

func genPlugin(baseDir string) {
	g := vi.New()

	g.AddSection(vi.Raw{Code: "vim9script"})

	g.AddSection(vi.NewVariables().
		Var("loaded_"+pluginName, "1").
		Var(pluginName+"_version", `"1.0.0"`))

	g.AddSection(vi.NewVariables().
		Legacy(vi.Global, "mapleader", `"<Space>"`))

	setupFn := vi.NewFunction("Setup").
		SetScope(vi.Global).
		Add(vi.Raw{Code: `echom "` + pluginName + ` loaded"`})

	toggleFn := vi.NewFunction("Toggle").
		SetScope(vi.Global).
		Add(vi.Raw{Code: `if exists('g:loaded_myvimlib') | unlet g:loaded_myvimlib | else | Setup() | endif`})

	g.AddSection(setupFn)
	g.AddSection(toggleFn)
	g.AddSection(vi.Keymap{Mode: "n", LHS: "<leader>p", RHS: ":call Toggle()<CR>"})
	g.AddSection(&vi.Command{Name: pluginName + "Toggle", Cmd: "call Toggle()"})

	output := g.Generate()
	pluginFile := filepath.Join(baseDir, "plugin", pluginName+".vim")
	os.WriteFile(pluginFile, []byte(output), 0644)
	fmt.Printf("Generated: %s\n", pluginFile)
}

func genAutoload(baseDir string) {
	g := vi.New()

	af := vi.NewAutoloadFile(pluginName).
		AddFunc(*vi.NewAutoloadFunc(pluginName, "Hello").
			Add(vi.Raw{Code: `echo "` + pluginDesc})).
		AddFunc(*vi.NewAutoloadFunc(pluginName, "Echo").
			Arg("msg", "string").
			Add(vi.Raw{Code: "echo msg"}))

	g.AddSection(af)

	output := g.Generate()
	autoloadFile := filepath.Join(baseDir, "autoload", pluginName, pluginName+".vim")
	os.WriteFile(autoloadFile, []byte(output), 0644)
	fmt.Printf("Generated: %s\n", autoloadFile)
}

func genDoc(baseDir string) {
	content := "* " + pluginName + ".txt *\n\n" +
		pluginName + "(1)    " + pluginName + "\n\n" +
		"DESCRIPTION\n" +
		pluginDesc + "\n"

	docFile := filepath.Join(baseDir, "doc", pluginName+".txt")
	os.WriteFile(docFile, []byte(content), 0644)
	fmt.Printf("Generated: %s\n", docFile)
}

func genReadme(baseDir string) {
	content := "# " + pluginName + "\n\n" +
		pluginDesc + "\n\n" +
		"## Installation\n\n" +
		"```bash\n" +
		"git clone " + pluginURL + ".git ~/.vim/pack/plugins/opt/" + pluginName + "\n" +
		"```\n"

	readmeFile := filepath.Join(baseDir, "README.md")
	os.WriteFile(readmeFile, []byte(content), 0644)
	fmt.Printf("Generated: %s\n", readmeFile)
}

func createDirectories(baseDir string) {
	dirs := []string{
		filepath.Join(baseDir, "plugin"),
		filepath.Join(baseDir, "autoload", pluginName),
		filepath.Join(baseDir, "ftplugin"),
		filepath.Join(baseDir, "syntax"),
		filepath.Join(baseDir, "doc"),
	}

	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
		fmt.Printf("Created: %s/\n", dir)
	}
}
