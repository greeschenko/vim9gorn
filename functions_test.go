package vim9gorn

import (
	"strings"
	"testing"
)

func TestFunction_Simple(t *testing.T) {
	f := NewFunction("MyFunc").
		Add(Raw{Code: "echo 'hello'"})

	expect := "def MyFunc()\n  echo 'hello'\nenddef\n\n"
	if got := f.Generate(); got != expect {
		t.Errorf("Simple function = %q, want %q", got, expect)
	}
}

func TestFunction_WithArg(t *testing.T) {
	f := NewFunction("Greet").
		Arg("name", "string").
		Add(Raw{Code: "echo 'Hello ' .. name"})

	if !strings.Contains(f.Generate(), "name: string") {
		t.Error("Function should contain typed arg")
	}
}

func TestFunction_WithReturnType(t *testing.T) {
	f := NewFunction("Add").
		Arg("a", "number").
		Arg("b", "number").
		Returns("number").
		Add(Raw{Code: "return a + b"})

	expect := "def Add(a: number, b: number): number\n  return a + b\nenddef\n\n"
	if got := f.Generate(); got != expect {
		t.Errorf("Function with return = %q, want %q", got, expect)
	}
}

func TestFunction_WithScope(t *testing.T) {
	f := NewFunction("Init").
		SetScope(Global).
		Add(Raw{Code: "echo 'init'"})

	if !strings.Contains(f.Generate(), "g:Init") {
		t.Error("Function should have g: prefix")
	}
}

func TestFunction_EmptyBody(t *testing.T) {
	f := NewFunction("Empty")
	expect := "def Empty()\nenddef\n\n"
	if got := f.Generate(); got != expect {
		t.Errorf("Empty body = %q, want %q", got, expect)
	}
}

func TestFunction_MultipleArgs(t *testing.T) {
	f := NewFunction("Calc").
		Arg("x", "number").
		Arg("y", "number").
		Arg("z", "number").
		Add(Raw{Code: "return x + y + z"})

	if got := f.Generate(); !strings.Contains(got, "x: number, y: number, z: number") {
		t.Errorf("Multiple args failed, got: %s", got)
	}
}
