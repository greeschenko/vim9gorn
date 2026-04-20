package vim9gorn

import (
	"testing"
)

func TestAutocmd_Simple(t *testing.T) {
	a := NewAutocmd("FileType", "go").
		SetCmd("setlocal shiftwidth=4")

	expect := "autocmd FileType go setlocal shiftwidth=4"
	if got := a.Generate(); got != expect {
		t.Errorf("Autocmd = %q, want %q", got, expect)
	}
}

func TestAutocmd_Nested(t *testing.T) {
	a := NewAutocmd("BufRead", "*.vim").
		SetCmd("source ~/.vim/plugin/init.vim").
		SetNested()

	expect := "autocmd BufRead *.vim ++nested source ~/.vim/plugin/init.vim"
	if got := a.Generate(); got != expect {
		t.Errorf("Nested = %q, want %q", got, expect)
	}
}

func TestAutocmdGroup(t *testing.T) {
	g := NewAutocmdGroup("MyGroup").
		Add(*NewAutocmd("FileType", "go").SetCmd("echo 'go'")).
		Add(*NewAutocmd("FileType", "rs").SetCmd("echo 'rust'"))

	got := g.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}
