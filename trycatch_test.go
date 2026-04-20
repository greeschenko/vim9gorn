package vim9gorn

import (
	"testing"
)

func TestTryCatch_TryOnly(t *testing.T) {
	tc := NewTryCatch().
		AddTry(Raw{Code: "echo 'trying'"})

	if got := tc.Generate(); got == "" {
		t.Error("Generate() = empty")
	}
}

func TestTryCatch_TryCatch(t *testing.T) {
	tc := NewTryCatch().
		AddTry(Raw{Code: "some_call()"}).
		SetCatch("e", "Vim:.*").
		AddCatch(Raw{Code: "echo 'caught: ' .. e"})

	got := tc.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}

func TestTryCatch_TryCatchFinally(t *testing.T) {
	tc := NewTryCatch().
		AddTry(Raw{Code: "let x = 1"}).
		SetCatch("err", ".*").
		AddCatch(Raw{Code: "echo err"}).
		AddFinally(Raw{Code: "unlet x"})

	got := tc.Generate()
	if got == "" {
		t.Error("Generate() = empty")
	}
}

func TestTryCatch_Nested(t *testing.T) {
	tc := NewTryCatch().
		AddTry(
			NewIfElse("cond").
				ThenAdd(Raw{Code: "throw 'error'"}),
		)

	got := tc.Generate()
	if got == "" {
		t.Error("Nested failed")
	}
}
