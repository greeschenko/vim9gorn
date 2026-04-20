package vim9gorn

import (
	"testing"
)

func TestOptions_SetBool(t *testing.T) {
	o := NewOptions().
		Set("number", true)

	expect := "# === Options ===\nset number\n\n"
	if got := o.Generate(); got != expect {
		t.Errorf("Set(bool) = %q, want %q", got, expect)
	}
}

func TestOptions_SetBoolFalse(t *testing.T) {
	o := NewOptions().
		Set("relativenumber", false)

	expect := "# === Options ===\nset norelativenumber\n\n"
	if got := o.Generate(); got != expect {
		t.Errorf("Set(false) = %q, want %q", got, expect)
	}
}

func TestOptions_SetInt(t *testing.T) {
	o := NewOptions().
		Set("tabstop", 4)

	expect := "# === Options ===\nset tabstop=4\n\n"
	if got := o.Generate(); got != expect {
		t.Errorf("Set(int) = %q, want %q", got, expect)
	}
}

func TestOptions_SetString(t *testing.T) {
	o := NewOptions().
		Set("expandtab", "shiftwidth=4")

	expect := "# === Options ===\nset expandtab=shiftwidth=4\n\n"
	if got := o.Generate(); got != expect {
		t.Errorf("Set(string) = %q, want %q", got, expect)
	}
}

func TestOptions_Chained(t *testing.T) {
	o := NewOptions().
		Set("number", true).
		Set("tabstop", 4).
		Set("relativenumber", true)

	got := o.Generate()
	if got == "" {
		t.Error("Generate() = empty, want non-empty")
	}
}

func TestOptions_Empty(t *testing.T) {
	o := NewOptions()
	if got := o.Generate(); got != "" {
		t.Errorf("Empty Generate() = %q, want empty", got)
	}
}
