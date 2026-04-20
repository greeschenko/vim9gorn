package vim9gorn

import (
	"testing"
)

func TestReturn_SimpleValue(t *testing.T) {
	r := NewReturn("42")
	expect := "return 42"
	if got := r.Generate(); got != expect {
		t.Errorf("Return() = %q, want %q", got, expect)
	}
}

func TestReturn_StringValue(t *testing.T) {
	r := NewReturn(`"hello"`)
	expect := "return \"hello\""
	if got := r.Generate(); got != expect {
		t.Errorf("Return string = %q, want %q", got, expect)
	}
}

func TestReturn_Variable(t *testing.T) {
	r := NewReturn("myvar")
	expect := "return myvar"
	if got := r.Generate(); got != expect {
		t.Errorf("Return var = %q, want %q", got, expect)
	}
}

func TestReturn_Expression(t *testing.T) {
	r := NewReturn("a + b")
	expect := "return a + b"
	if got := r.Generate(); got != expect {
		t.Errorf("Return expr = %q, want %q", got, expect)
	}
}
