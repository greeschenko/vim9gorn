package vim9gorn

import (
	"strings"
	"testing"
)

func TestAutoloadFunc_Simple(t *testing.T) {
	f := NewAutoloadFunc("mylib", "Hello").
		Add(Raw{Code: "echo 'hello'"})

	got := f.Generate()
	if !strings.Contains(got, "mylib#Hello") {
		t.Errorf("Function name should contain mylib#Hello, got: %s", got)
	}
}

func TestAutoloadFunc_WithArgs(t *testing.T) {
	f := NewAutoloadFunc("utils", "Add").
		Arg("a", "number").
		Arg("b", "number").
		SetReturn("number").
		Add(NewReturn("a + b"))

	got := f.Generate()
	if !strings.Contains(got, "a: number, b: number") {
		t.Errorf("Should have typed args, got: %s", got)
	}
	if !strings.Contains(got, ": number") {
		t.Errorf("Should have return type, got: %s", got)
	}
}

func TestAutoloadFile(t *testing.T) {
	af := NewAutoloadFile("myplugin").
		AddFunc(*NewAutoloadFunc("myplugin", "Init").
			Add(Raw{Code: "echo 'init'"}),
		)

	got := af.Generate()
	if !strings.Contains(got, "vim9script") {
		t.Error("Should have vim9script header")
	}
	if !strings.Contains(got, "# Autoload: myplugin") {
		t.Error("Should have autoload comment")
	}
}
