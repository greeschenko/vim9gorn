package vim9gorn

import (
	"testing"
)

func TestIfElse_Simple(t *testing.T) {
	i := NewIfElse("true").
		ThenAdd(Raw{Code: "echo 'yes'"})

	expect := "if true\n  echo 'yes'\nendif\n"
	if got := i.Generate(); got != expect {
		t.Errorf("Simple if = %q, want %q", got, expect)
	}
}

func TestIfElse_IfElse(t *testing.T) {
	i := NewIfElse("count > 0").
		ThenAdd(Raw{Code: "echo 'positive'"}).
		ElseAdd(Raw{Code: "echo 'zero or negative'"})

	got := i.Generate()
	if got == "" {
		t.Error("Generate() = empty, want non-empty")
	}
}

func TestIfElse_ElseIf(t *testing.T) {
	i := NewIfElse("hour < 12").
		ThenAdd(Raw{Code: "echo 'morning'"}).
		ElseIfAdd("hour < 18", Raw{Code: "echo 'afternoon'"}).
		ElseIfAdd("hour < 22", Raw{Code: "echo 'evening'"}).
		ElseAdd(Raw{Code: "echo 'night'"})

	got := i.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}

func TestIfElse_Nested(t *testing.T) {
	i := NewIfElse("a == 1").
		ThenAdd(
			NewIfElse("b == 2").
				ThenAdd(Raw{Code: "echo 'a=1, b=2'"}),
		)

	got := i.Generate()
	if got == "" {
		t.Error("Nested if failed")
	}
}

func TestIfElse_EmptyThen(t *testing.T) {
	i := NewIfElse("cond")
	got := i.Generate()
	if got == "" {
		t.Error("Empty then branch failed")
	}
}
