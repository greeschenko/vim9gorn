package vim9gorn

import (
	"testing"
)

func TestVariables_Var(t *testing.T) {
	v := NewVariables().
		Var("foo", `"bar"`)

	expect := "# === Variables ===\nvar foo = \"bar\"\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("Var() = %q, want %q", got, expect)
	}
}

func TestVariables_VarTyped(t *testing.T) {
	v := NewVariables().
		VarTyped("count", "number", "42")

	expect := "# === Variables ===\nvar count: number = 42\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("VarTyped() = %q, want %q", got, expect)
	}
}

func TestVariables_Const(t *testing.T) {
	v := NewVariables().
		Const("PI", "3.14")

	expect := "# === Variables ===\nconst PI = 3.14\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("Const() = %q, want %q", got, expect)
	}
}

func TestVariables_ConstTyped(t *testing.T) {
	v := NewVariables().
		ConstTyped("name", "string", `"test"`)

	expect := "# === Variables ===\nconst name: string = \"test\"\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("ConstTyped() = %q, want %q", got, expect)
	}
}

func TestVariables_Legacy(t *testing.T) {
	v := NewVariables().
		Legacy(Global, "mapleader", `"\\<Space>"`)

	expect := "# === Variables ===\ng:mapleader = \"\\\\<Space>\"\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("Legacy() = %q, want %q", got, expect)
	}
}

func TestVariables_LegacyBuffer(t *testing.T) {
	v := NewVariables().
		Legacy(Buffer, "myvar", "true")

	expect := "# === Variables ===\nb:myvar = true\n\n"
	if got := v.Generate(); got != expect {
		t.Errorf("Legacy(Buffer) = %q, want %q", got, expect)
	}
}

func TestVariables_Multiple(t *testing.T) {
	v := NewVariables().
		Var("foo", "1").
		Const("BAR", "2").
		Legacy(Global, "baz", "3")

	got := v.Generate()
	if got == "" {
		t.Error("Generate() = empty, want non-empty")
	}
}

func TestVariables_Empty(t *testing.T) {
	v := NewVariables()
	if got := v.Generate(); got != "" {
		t.Errorf("Empty Generate() = %q, want empty", got)
	}
}
