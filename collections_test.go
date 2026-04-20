package vim9gorn

import (
	"testing"
)

func TestList(t *testing.T) {
	l := NewList("1", "2", "3")
	expect := "[1, 2, 3]"
	if got := l.Generate(); got != expect {
		t.Errorf("List = %q, want %q", got, expect)
	}
}

func TestList_Single(t *testing.T) {
	l := NewList("42")
	expect := "[42]"
	if got := l.Generate(); got != expect {
		t.Errorf("Single item = %q, want %q", got, expect)
	}
}

func TestList_Add(t *testing.T) {
	l := NewList("1").Add("2").Add("3")
	expect := "[1, 2, 3]"
	if got := l.Generate(); got != expect {
		t.Errorf("List.Add = %q, want %q", got, expect)
	}
}

func TestList_Strings(t *testing.T) {
	l := NewList(`"a"`, `"b"`, `"c"`)
	expect := "[\"a\", \"b\", \"c\"]"
	if got := l.Generate(); got != expect {
		t.Errorf("String list = %q, want %q", got, expect)
	}
}

func TestDict(t *testing.T) {
	d := NewDict().
		Set("a", "1").
		Set("b", "2")

	expect := "{\"a\": 1, \"b\": 2}"
	if got := d.Generate(); got != expect {
		t.Errorf("Dict = %q, want %q", got, expect)
	}
}

func TestDict_Empty(t *testing.T) {
	d := NewDict()
	expect := "{}"
	if got := d.Generate(); got != expect {
		t.Errorf("Empty dict = %q, want %q", got, expect)
	}
}

func TestDict_Chained(t *testing.T) {
	d := NewDict().
		Set("name", `"vim"`).
		Set("version", "9").
		Set("license", `"MIT"`)

	got := d.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}

func TestValues(t *testing.T) {
	d := NewDict().Set("a", "1")
	v := Values(d)
	expect := "values({\"a\": 1})"
	if got := v; got != expect {
		t.Errorf("Values() = %q, want %q", got, expect)
	}
}

func TestKeys(t *testing.T) {
	d := NewDict().Set("a", "1")
	k := Keys(d)
	expect := "keys({\"a\": 1})"
	if got := k; got != expect {
		t.Errorf("Keys() = %q, want %q", got, expect)
	}
}

func TestItems(t *testing.T) {
	d := NewDict().Set("a", "1")
	i := Items(d)
	expect := "items({\"a\": 1})"
	if got := i; got != expect {
		t.Errorf("Items() = %q, want %q", got, expect)
	}
}
