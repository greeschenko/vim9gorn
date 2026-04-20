package vim9gorn

import (
	"fmt"
	"strings"
)

type TryCatch struct {
	Try       []CodeBlock
	CatchVar  string
	CatchPat  string
	CatchBody []CodeBlock
	Finally   []CodeBlock
}

func NewTryCatch() *TryCatch {
	return &TryCatch{
		Try:       make([]CodeBlock, 0),
		CatchBody: make([]CodeBlock, 0),
		Finally:   make([]CodeBlock, 0),
	}
}

func (t *TryCatch) SetCatch(varname, pattern string) *TryCatch {
	t.CatchVar = varname
	t.CatchPat = pattern
	return t
}

func (t *TryCatch) AddTry(block CodeBlock) *TryCatch {
	t.Try = append(t.Try, block)
	return t
}

func (t *TryCatch) AddCatch(block CodeBlock) *TryCatch {
	t.CatchBody = append(t.CatchBody, block)
	return t
}

func (t *TryCatch) AddFinally(block CodeBlock) *TryCatch {
	t.Finally = append(t.Finally, block)
	return t
}

func (t *TryCatch) Generate() string {
	var b strings.Builder

	b.WriteString("try\n")

	for _, block := range t.Try {
		lines := strings.Split(block.Generate(), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			b.WriteString("  " + line + "\n")
		}
	}

	if t.CatchVar != "" {
		b.WriteString(fmt.Sprintf("catch /%s/ as %s\n", t.CatchPat, t.CatchVar))
		for _, block := range t.CatchBody {
			lines := strings.Split(block.Generate(), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				b.WriteString("  " + line + "\n")
			}
		}
	}

	if len(t.Finally) > 0 {
		b.WriteString("finally\n")
		for _, block := range t.Finally {
			lines := strings.Split(block.Generate(), "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) == "" {
					continue
				}
				b.WriteString("  " + line + "\n")
			}
		}
	}

	b.WriteString("endtry\n")
	return b.String()
}
