package vim9gorn

import (
	"testing"
)

func TestPlugin_New(t *testing.T) {
	p := NewPlugin("mylib", PluginTypeOpt)

	if p.Name != "mylib" {
		t.Errorf("Name = %q, want %q", p.Name, "mylib")
	}
	if p.Type != PluginTypeOpt {
		t.Errorf("Type = %q, want %q", p.Type, PluginTypeOpt)
	}
}

func TestPlugin_GetPath(t *testing.T) {
	p := NewPlugin("vim-commentary", PluginTypeStart)
	got := p.GetPath()
	expect := ".vim/pack/plugins/start/vim-commentary"
	if got != expect {
		t.Errorf("GetPath() = %q, want %q", got, expect)
	}
}

func TestPlugin_SetRepo(t *testing.T) {
	p := NewPlugin("fugitive", PluginTypeOpt)
	p.SetRepo("tpope/vim-fugitive")

	if p.Repo != "tpope/vim-fugitive" {
		t.Errorf("Repo = %q, want %q", p.Repo, "tpope/vim-fugitive")
	}
}

func TestPlugin_GetPathTypeOpt(t *testing.T) {
	p := NewPlugin("nerdtree", PluginTypeOpt)
	got := p.GetPath()
	expect := ".vim/pack/plugins/opt/nerdtree"
	if got != expect {
		t.Errorf("GetPath() = %q, want %q", got, expect)
	}
}

func TestManifest_New(t *testing.T) {
	m := NewManifest()
	if m.OutputDir != ".vim" {
		t.Errorf("OutputDir = %q", m.OutputDir)
	}
}

func TestManifest_Add(t *testing.T) {
	m := NewManifest().
		Add(Plugin{Name: "A"}).
		Add(Plugin{Name: "B"})

	if len(m.Plugins) != 2 {
		t.Errorf("len(Plugins) = %d, want 2", len(m.Plugins))
	}
}

func TestManifest_Generate(t *testing.T) {
	m := NewManifest().
		Add(Plugin{Name: "testplugin", Type: PluginTypeOpt})

	got := m.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}
