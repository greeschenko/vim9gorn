package vim9gorn

import (
	"fmt"
)

type Highlight struct {
	LinkFrom string
	LinkTo   string
	Group    string
	Args     string
}

func (h Highlight) Generate() string {
	if h.LinkFrom != "" {
		return fmt.Sprintf("highlight link %s %s", h.LinkFrom, h.LinkTo)
	}
	return fmt.Sprintf("highlight %s %s", h.Group, h.Args)
}
