package vim9gorn

import (
	"testing"
)

func TestExternalPlugin_New(t *testing.T) {
	p := NewExternalPlugin("tpope/vim-fugitive")
	if p.Repo != "tpope/vim-fugitive" {
		t.Errorf("Repo = %q", p.Repo)
	}
	if p.Dir != ".vim/pack/plugins/opt/vim-fugitive" {
		t.Errorf("Dir = %q, want .vim/pack/plugins/opt/vim-fugitive", p.Dir)
	}
}

func TestPluginManager_New(t *testing.T) {
	m := NewPluginManager()
	if m.VimDir != ".vim" {
		t.Errorf("VimDir = %q", m.VimDir)
	}
}

func TestPluginManager_Add(t *testing.T) {
	m := NewPluginManager().
		Add("foo/bar").
		Add("baz/qux")

	if len(m.Plugins) != 2 {
		t.Errorf("len(Plugins) = %d, want 2", len(m.Plugins))
	}
}

func TestPluginManager_Generate(t *testing.T) {
	m := NewPluginManager().
		Add("tpope/vim-fugitive")

	got := m.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}
