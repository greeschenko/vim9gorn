package vim9gorn

import (
	"strings"
	"testing"
)

func TestClass_Simple(t *testing.T) {
	c := NewClass("MyClass").
		AddField("name", "string")

	got := c.Generate()
	if !strings.Contains(got, "class MyClass") {
		t.Errorf("Should contain class declaration, got: %s", got)
	}
	if !strings.Contains(got, "this.name: string") {
		t.Errorf("Should contain field, got: %s", got)
	}
}

func TestClass_WithDefault(t *testing.T) {
	c := NewClass("Counter").
		AddFieldWithDefault("value", "number", "0")

	got := c.Generate()
	if !strings.Contains(got, "this.value: number = 0") {
		t.Errorf("Should contain field with default, got: %s", got)
	}
}

func TestClass_WithSuper(t *testing.T) {
	c := NewClass("Dog").
		SetSuper("Animal").
		AddField("name", "string")

	got := c.Generate()
	if !strings.Contains(got, "extends Animal") {
		t.Errorf("Should contain extends, got: %s", got)
	}
}

func TestClass_WithMethod(t *testing.T) {
	c := NewClass("Person").
		AddField("name", "string").
		AddMethod(
			NewFunction("GetName").
				Returns("string").
				Add(NewReturn("this.name")),
		)

	got := c.Generate()
	if !strings.Contains(got, "def GetName") {
		t.Errorf("Should contain method, got: %s", got)
	}
	if !strings.Contains(got, "this.name") {
		t.Errorf("Should access field, got: %s", got)
	}
}

func TestClass_MultipleFields(t *testing.T) {
	c := NewClass("Point").
		AddField("x", "number").
		AddField("y", "number").
		AddFieldWithDefault("z", "number", "0")

	got := c.Generate()
	if !strings.Contains(got, "this.x: number") {
		t.Errorf("Missing x field, got: %s", got)
	}
	if !strings.Contains(got, "this.y: number") {
		t.Errorf("Missing y field, got: %s", got)
	}
}

func TestClass_EndClass(t *testing.T) {
	c := NewClass("Test")
	got := c.Generate()
	if !strings.Contains(got, "endclass") {
		t.Errorf("Should contain endclass, got: %s", got)
	}
}

func TestClassInstance(t *testing.T) {
	ci := NewClassInstance("Person", "Alice", "30")
	got := ci.Generate()
	expect := "Person->new(Alice, 30)"
	if got != expect {
		t.Errorf("ClassInstance = %q, want %q", got, expect)
	}
}

func TestClassInstance_NoArgs(t *testing.T) {
	ci := NewClassInstance("Person")
	got := ci.Generate()
	expect := "Person->new()"
	if got != expect {
		t.Errorf("ClassInstance = %q, want %q", got, expect)
	}
}
