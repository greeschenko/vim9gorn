package vim9gorn

import (
	"testing"
)

func TestLambda_Simple(t *testing.T) {
	l := NewLambda("echo 'hello'")
	got := l.Generate()
	expect := "() => echo 'hello'"
	if got != expect {
		t.Errorf("Lambda = %q, want %q", got, expect)
	}
}

func TestLambda_WithArgs(t *testing.T) {
	l := NewLambda("x + 1").
		Arg("x")

	got := l.Generate()
	expect := "(x) => x + 1"
	if got != expect {
		t.Errorf("Lambda = %q, want %q", got, expect)
	}
}

func TestLambda_MultipleArgs(t *testing.T) {
	l := NewLambda("x + y").
		Arg("x").
		Arg("y")

	got := l.Generate()
	expect := "(x, y) => x + y"
	if got != expect {
		t.Errorf("Lambda = %q, want %q", got, expect)
	}
}

func TestLambda_WithReturn(t *testing.T) {
	l := NewLambda("x + 1").
		Arg("x").
		SetReturn("number")

	got := l.Generate()
	expect := "(x) => x + 1: number"
	if got != expect {
		t.Errorf("Lambda = %q, want %q", got, expect)
	}
}

func TestLambda_WithTypedArgs(t *testing.T) {
	l := NewLambda("x + y").
		Arg("x: number").
		Arg("y: number").
		SetReturn("number")

	got := l.Generate()
	expect := "(x: number, y: number) => x + y: number"
	if got != expect {
		t.Errorf("Lambda = %q, want %q", got, expect)
	}
}

func TestClosure(t *testing.T) {
	c := NewClosure("a + b").
		Arg("a").
		AddCapture("b")

	got := c.Generate()
	expect := "(a) => a + b"
	if got != expect {
		t.Errorf("Closure = %q, want %q", got, expect)
	}
}

func TestLambdaCall(t *testing.T) {
	l := NewLambda("item * 2")
	lc := NewLambdaCall(l)
	got := lc.Call("items")
	expect := "() => item * 2(items)"
	if got != expect {
		t.Errorf("LambdaCall = %q, want %q", got, expect)
	}
}

func TestFilter(t *testing.T) {
	pred := NewLambda("x > 0").Arg("x")
	got := Filter("[1, -1, 2, -2]", pred)
	expect := "filter([1, -1, 2, -2], (x) => x > 0)"
	if got != expect {
		t.Errorf("Filter = %q, want %q", got, expect)
	}
}

func TestMap(t *testing.T) {
	fn := NewLambda("x * 2").Arg("x")
	got := Map("[1, 2, 3]", fn)
	expect := "map([1, 2, 3], (x) => x * 2)"
	if got != expect {
		t.Errorf("Map = %q, want %q", got, expect)
	}
}

func TestForEach(t *testing.T) {
	fn := NewLambda("echo item").Arg("item")
	fl := ForEach("mylist", fn)
	got := fl.Generate()
	if got == "" {
		t.Error("ForEach generated empty output")
	}
}
