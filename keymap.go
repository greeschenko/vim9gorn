package vim9gorn

import (
	"fmt"
)

type Keymap struct {
	Mode   string
	LHS    string
	RHS    string
	Silent bool
	Nor    bool
}

func (k Keymap) Generate() string {
	cmd := k.Mode + "map"
	if k.Nor {
		cmd = k.Mode + "noremap"
	}
	if k.Silent {
		cmd += " <silent>"
	}
	return fmt.Sprintf("%s %s %s", cmd, k.LHS, k.RHS)
}
