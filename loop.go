package vim9gorn

import (
	"fmt"
	"strings"
)

// =====================
// ForLoop
// =====================
type ForLoop struct {
	VarName  string // loop variable
	Iterable string // e.g., "range(1,5)" або "[1,2,3]"
	Body     []CodeBlock
}

// NewForLoop creates a new ForLoop
func NewForLoop(varName, iterable string) *ForLoop {
	return &ForLoop{
		VarName:  varName,
		Iterable: iterable,
		Body:     make([]CodeBlock, 0),
	}
}

// Add adds a CodeBlock to the loop body
func (f *ForLoop) Add(block CodeBlock) *ForLoop {
	f.Body = append(f.Body, block)
	return f
}

// Generate generates Vim9 for loop code
func (f *ForLoop) Generate() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("for %s in %s\n", f.VarName, f.Iterable))
	for _, block := range f.Body {
		code := block.Generate()
		lines := strings.Split(code, "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString(functionIndent)
			b.WriteString(line)
			b.WriteByte('\n')
		}
	}
	b.WriteString("endfor\n")
	return b.String()
}

// =====================
// WhileLoop
// =====================
type WhileLoop struct {
	Condition string
	Body      []CodeBlock
}

// NewWhileLoop creates a new WhileLoop
func NewWhileLoop(cond string) *WhileLoop {
	return &WhileLoop{
		Condition: cond,
		Body:      make([]CodeBlock, 0),
	}
}

// Add adds a CodeBlock to the while loop body
func (w *WhileLoop) Add(block CodeBlock) *WhileLoop {
	w.Body = append(w.Body, block)
	return w
}

// Generate generates Vim9 while loop code
func (w *WhileLoop) Generate() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("while %s\n", w.Condition))
	for _, block := range w.Body {
		code := block.Generate()
		lines := strings.Split(code, "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString(functionIndent)
			b.WriteString(line)
			b.WriteByte('\n')
		}
	}
	b.WriteString("endwhile\n")
	return b.String()
}

// =====================
// Break & Continue
// =====================
type Break struct{}

func NewBreak() *Break            { return &Break{} }
func (b *Break) Generate() string { return "break" }

type Continue struct{}

func NewContinue() *Continue         { return &Continue{} }
func (c *Continue) Generate() string { return "continue" }
