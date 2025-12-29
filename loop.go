package vim9gorn

import (
	"fmt"
	"strings"
)

// =====================
// ForLoop (key/value based)
// =====================
type ForLoop struct {
	Key      string // "_" to ignore
	Value    string // "_" to ignore
	Iterable string // range(...), list, dict, values(...), keys(...)
	Body     []CodeBlock
}

// NewForLoop creates a new ForLoop
func NewForLoop(key, value, iterable string) *ForLoop {
	if key == "_" && value == "_" {
		panic("ForLoop: at least one variable must be defined")
	}

	// Invalid: key without value in Vim9
	if key != "_" && value == "_" {
		panic("ForLoop: key without value is not supported in Vim9")
	}

	return &ForLoop{
		Key:      key,
		Value:    value,
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

	b.WriteString("for ")

	switch {
	case f.Key != "_" && f.Value != "_":
		b.WriteString(f.Key)
		b.WriteString(", ")
		b.WriteString(f.Value)
	case f.Key == "_" && f.Value != "_":
		b.WriteString(f.Value)
	default:
		panic("ForLoop: invalid key/value combination")
	}

	b.WriteString(" in ")
	b.WriteString(f.Iterable)
	b.WriteByte('\n')

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

func NewWhileLoop(cond string) *WhileLoop {
	return &WhileLoop{
		Condition: cond,
		Body:      make([]CodeBlock, 0),
	}
}

func (w *WhileLoop) Add(block CodeBlock) *WhileLoop {
	w.Body = append(w.Body, block)
	return w
}

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
