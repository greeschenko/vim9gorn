package vim9gorn

import (
	"fmt"
)

type Keymap struct {
	Mode   string
	LHS    string
	RHS    string
	Silent bool
}

func (k Keymap) Generate() string {
	cmd := k.Mode + "map"
	if k.Silent {
		cmd += " <silent>"
	}
	return fmt.Sprintf("%s %s %s", cmd, k.LHS, k.RHS)
}
