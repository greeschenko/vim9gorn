package vim9gorn

import (
	"fmt"
	"strings"
)

type ErrorType struct {
	Name   string
	Number string
}

func NewErrorType(name, number string) *ErrorType {
	return &ErrorType{
		Name:   name,
		Number: number,
	}
}

func (e *ErrorType) Generate() string {
	return fmt.Sprintf("error %s('custom error'): '%s'", e.Name, e.Number)
}

type Throw struct {
	Expr string
}

func NewThrow(expr string) *Throw {
	return &Throw{Expr: expr}
}

func (t *Throw) Generate() string {
	return "throw " + t.Expr
}

type Assert struct {
	Cond string
	Msg  string
}

func NewAssert(cond string) *Assert {
	return &Assert{
		Cond: cond,
	}
}

func (a *Assert) SetMsg(msg string) *Assert {
	a.Msg = msg
	return a
}

func (a *Assert) Generate() string {
	if a.Msg != "" {
		return fmt.Sprintf("assert %s, '%s'", a.Cond, a.Msg)
	}
	return "assert " + a.Cond
}

type AssertEqual struct {
	Left  string
	Right string
	Msg   string
}

func NewAssertEqual(left, right string) *AssertEqual {
	return &AssertEqual{
		Left:  left,
		Right: right,
	}
}

func (a *AssertEqual) SetMsg(msg string) *AssertEqual {
	a.Msg = msg
	return a
}

func (a *AssertEqual) Generate() string {
	if a.Msg != "" {
		return fmt.Sprintf("assert_equal(%s, %s, '%s')", a.Left, a.Right, a.Msg)
	}
	return fmt.Sprintf("assert_equal(%s, %s)", a.Left, a.Right)
}

type AssertNotequal struct {
	Left  string
	Right string
}

func NewAssertNotequal(left, right string) *AssertNotequal {
	return &AssertNotequal{
		Left:  left,
		Right: right,
	}
}

func (a *AssertNotequal) Generate() string {
	return fmt.Sprintf("assert_notequal(%s, %s)", a.Left, a.Right)
}

type AssertTrue struct {
	Cond string
	Msg  string
}

func NewAssertTrue(cond string) *AssertTrue {
	return &AssertTrue{Cond: cond}
}

func (a *AssertTrue) SetMsg(msg string) *AssertTrue {
	a.Msg = msg
	return a
}

func (a *AssertTrue) Generate() string {
	if a.Msg != "" {
		return fmt.Sprintf("assert_true(%s, '%s')", a.Cond, a.Msg)
	}
	return "assert_true(" + a.Cond + ")"
}

type AssertFalse struct {
	Cond string
	Msg  string
}

func NewAssertFalse(cond string) *AssertFalse {
	return &AssertFalse{Cond: cond}
}

func (a *AssertFalse) SetMsg(msg string) *AssertFalse {
	a.Msg = msg
	return a
}

func (a *AssertFalse) Generate() string {
	if a.Msg != "" {
		return fmt.Sprintf("assert_false(%s, '%s')", a.Cond, a.Msg)
	}
	return "assert_false(" + a.Cond + ")"
}

type AssertException struct {
	Expr  string
	Error string
	Msg   string
}

func NewAssertException(expr string) *AssertException {
	return &AssertException{
		Expr: expr,
	}
}

func (a *AssertException) SetError(err string) *AssertException {
	a.Error = err
	return a
}

func (a *AssertException) SetMsg(msg string) *AssertException {
	a.Msg = msg
	return a
}

func (a *AssertException) Generate() string {
	if a.Error != "" && a.Msg != "" {
		return fmt.Sprintf("assert_exception(%s, '%s', '%s')", a.Expr, a.Error, a.Msg)
	}
	if a.Error != "" {
		return fmt.Sprintf("assert_exception(%s, '%s')", a.Expr, a.Error)
	}
	return "assert_exception(" + a.Expr + ")"
}

type Comment struct {
	Text string
}

func NewComment(text string) *Comment {
	return &Comment{Text: text}
}

func (c *Comment) Generate() string {
	return "# " + c.Text
}

type MultiLineComment struct {
	Lines []string
}

func NewMultiLineComment(lines ...string) *MultiLineComment {
	return &MultiLineComment{Lines: lines}
}

func (m *MultiLineComment) AddLine(line string) *MultiLineComment {
	m.Lines = append(m.Lines, line)
	return m
}

func (m *MultiLineComment) Generate() string {
	var b strings.Builder
	for _, line := range m.Lines {
		b.WriteString("# " + line + "\n")
	}
	return b.String()
}
