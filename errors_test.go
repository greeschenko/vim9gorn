package vim9gorn

import (
	"testing"
)

func TestErrorType(t *testing.T) {
	e := NewErrorType("MyError", "E1001")
	got := e.Generate()
	expect := "error MyError('custom error'): 'E1001'"
	if got != expect {
		t.Errorf("ErrorType = %q, want %q", got, expect)
	}
}

func TestThrow(t *testing.T) {
	tw := NewThrow(`"something went wrong"`)
	got := tw.Generate()
	expect := `throw "something went wrong"`
	if got != expect {
		t.Errorf("Throw = %q, want %q", got, expect)
	}
}

func TestAssert(t *testing.T) {
	a := NewAssert("x > 0")
	got := a.Generate()
	expect := "assert x > 0"
	if got != expect {
		t.Errorf("Assert = %q, want %q", got, expect)
	}
}

func TestAssertWithMsg(t *testing.T) {
	a := NewAssert("x > 0").SetMsg("x must be positive")
	got := a.Generate()
	expect := "assert x > 0, 'x must be positive'"
	if got != expect {
		t.Errorf("Assert = %q, want %q", got, expect)
	}
}

func TestAssertEqual(t *testing.T) {
	a := NewAssertEqual("actual", "expected")
	got := a.Generate()
	expect := "assert_equal(actual, expected)"
	if got != expect {
		t.Errorf("AssertEqual = %q, want %q", got, expect)
	}
}

func TestAssertEqualWithMsg(t *testing.T) {
	a := NewAssertEqual("a", "b").SetMsg("values must match")
	got := a.Generate()
	expect := "assert_equal(a, b, 'values must match')"
	if got != expect {
		t.Errorf("AssertEqual = %q, want %q", got, expect)
	}
}

func TestAssertTrue(t *testing.T) {
	a := NewAssertTrue("condition")
	got := a.Generate()
	expect := "assert_true(condition)"
	if got != expect {
		t.Errorf("AssertTrue = %q, want %q", got, expect)
	}
}

func TestAssertFalse(t *testing.T) {
	a := NewAssertFalse("condition")
	got := a.Generate()
	expect := "assert_false(condition)"
	if got != expect {
		t.Errorf("AssertFalse = %q, want %q", got, expect)
	}
}

func TestAssertException(t *testing.T) {
	a := NewAssertException("may_fail()")
	got := a.Generate()
	expect := "assert_exception(may_fail())"
	if got != expect {
		t.Errorf("AssertException = %q, want %q", got, expect)
	}
}

func TestAssertExceptionWithError(t *testing.T) {
	a := NewAssertException("may_fail()").SetError("E1234")
	got := a.Generate()
	expect := "assert_exception(may_fail(), 'E1234')"
	if got != expect {
		t.Errorf("AssertException = %q, want %q", got, expect)
	}
}

func TestComment(t *testing.T) {
	c := NewComment("This is a comment")
	got := c.Generate()
	expect := "# This is a comment"
	if got != expect {
		t.Errorf("Comment = %q, want %q", got, expect)
	}
}

func TestMultiLineComment(t *testing.T) {
	c := NewMultiLineComment("Line 1", "Line 2", "Line 3")
	got := c.Generate()
	expect := "# Line 1\n# Line 2\n# Line 3\n"
	if got != expect {
		t.Errorf("MultiLineComment = %q, want %q", got, expect)
	}
}
