package vim9gorn

/*
Raw represents a raw Vim9 code block.

It is an escape hatch that allows embedding
any valid Vim9Script directly into the output.

Indentation is handled by the parent (e.g. Function).
*/
type Raw struct {
	Code string
}

// Generate returns the raw Vim9 code as-is.
func (r Raw) Generate() string {
	return r.Code
}
