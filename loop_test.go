package vim9gorn

import (
	"testing"
)

func TestForLoop_ValueOnly(t *testing.T) {
	f := NewForLoop("_", "i", "[1, 2, 3]").
		Add(Raw{Code: "echo i"})

	expect := "for i in [1, 2, 3]\n  echo i\nendfor\n"
	if got := f.Generate(); got != expect {
		t.Errorf("for i in list = %q, want %q", got, expect)
	}
}

func TestForLoop_KeyValue(t *testing.T) {
	f := NewForLoop("k", "v", "items(dict)").
		Add(Raw{Code: "echo k .. ': ' .. v"})

	expect := "for [k, v] in items(dict)\n  echo k .. ': ' .. v\nendfor\n"
	if got := f.Generate(); got != expect {
		t.Errorf("for [k, v] = %q, want %q", got, expect)
	}
}

func TestForLoop_ChainedAdd(t *testing.T) {
	f := NewForLoop("_", "item", "mylist").
		Add(Raw{Code: "echo item"}).
		Add(Raw{Code: "sleep 1m"})

	got := f.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}

func TestForLoop_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for invalid key/value")
		}
	}()

	NewForLoop("_", "_", "[1,2,3]")
}

func TestForLoop_PanicKeyWithoutValue(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for key without value")
		}
	}()

	NewForLoop("k", "_", "[1,2,3]")
}

func TestWhileLoop(t *testing.T) {
	w := NewWhileLoop("i < 10").
		Add(Raw{Code: "let i += 1"})

	expect := "while i < 10\n  let i += 1\nendwhile\n"
	if got := w.Generate(); got != expect {
		t.Errorf("while = %q, want %q", got, expect)
	}
}

func TestWhileLoop_Nested(t *testing.T) {
	w := NewWhileLoop("i < 10").
		Add(NewForLoop("_", "j", "[1,2,3]").Add(Raw{Code: "echo j"}))

	got := w.Generate()
	if got == "" {
		t.Error("Nested loop failed")
	}
}

func TestBreak(t *testing.T) {
	b := NewBreak()
	expect := "break"
	if got := b.Generate(); got != expect {
		t.Errorf("Break = %q, want %q", got, expect)
	}
}

func TestContinue(t *testing.T) {
	c := NewContinue()
	expect := "continue"
	if got := c.Generate(); got != expect {
		t.Errorf("Continue = %q, want %q", got, expect)
	}
}

func TestRange(t *testing.T) {
	r := NewRange(1, 10)
	expect := "range(1, 10)"
	if got := r.Generate(); got != expect {
		t.Errorf("Range = %q, want %q", got, expect)
	}
}

func TestRangeWithStep(t *testing.T) {
	r := NewRangeWithStep(0, 100, 5)
	expect := "range(0, 100, 5)"
	if got := r.Generate(); got != expect {
		t.Errorf("RangeWithStep = %q, want %q", got, expect)
	}
}
