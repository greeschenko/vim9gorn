package vim9gorn

import (
	"fmt"
	"strings"
)

// VarKind визначає синтаксичний тип змінної у Vim9script
type VarKind int

const (
	VarLegacy   VarKind = iota // g:, b:, w:, t:
	VarScript                  // var
	ConstScript                // const
)

// Scope використовується тільки для legacy-змінних
type Scope string

const (
	None    Scope = ""
	Global  Scope = "g"
	Buffer  Scope = "b"
	Window  Scope = "w"
	Tabpage Scope = "t"
	Script  Scope = "s"
)

// Variable описує одну змінну Vim9script
type Variable struct {
	Kind  VarKind
	Scope Scope // тільки для VarLegacy
	Name  string
	Type  string // optional: number, string, bool, etc.
	Value string
}

// Variables — контейнер змінних
type Variables struct {
	data []Variable
}

// NewVariables створює новий контейнер змінних
func NewVariables() *Variables {
	return &Variables{
		data: make([]Variable, 0),
	}
}

//
// ===== Public API (зручний і безпечний) =====
//

// Var додає script-local змінну (vim9: var)
func (v *Variables) Var(name, value string) {
	v.data = append(v.data, Variable{
		Kind:  VarScript,
		Name:  name,
		Value: value,
	})
}

// VarTyped додає script-local змінну з типом
func (v *Variables) VarTyped(name, typ, value string) {
	v.data = append(v.data, Variable{
		Kind:  VarScript,
		Name:  name,
		Type:  typ,
		Value: value,
	})
}

// Const додає константу (vim9: const)
func (v *Variables) Const(name, value string) {
	v.data = append(v.data, Variable{
		Kind:  ConstScript,
		Name:  name,
		Value: value,
	})
}

// ConstTyped додає константу з типом
func (v *Variables) ConstTyped(name, typ, value string) {
	v.data = append(v.data, Variable{
		Kind:  ConstScript,
		Name:  name,
		Type:  typ,
		Value: value,
	})
}

// Legacy додає legacy-змінну з префіксом (g:, b:, w:, t:)
func (v *Variables) Legacy(scope Scope, name, value string) {
	v.data = append(v.data, Variable{
		Kind:  VarLegacy,
		Scope: scope,
		Name:  name,
		Value: value,
	})
}

//
// ===== Code generation =====
//

// Generate повертає блок змінних у форматі Vim9script
func (v *Variables) Generate() string {
	if len(v.data) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("# === Variables ===\n")

	for _, vr := range v.data {
		switch vr.Kind {

		case VarScript:
			if vr.Type != "" {
				b.WriteString(fmt.Sprintf("var %s: %s = %s\n", vr.Name, vr.Type, vr.Value))
			} else {
				b.WriteString(fmt.Sprintf("var %s = %s\n", vr.Name, vr.Value))
			}

		case ConstScript:
			if vr.Type != "" {
				b.WriteString(fmt.Sprintf("const %s: %s = %s\n", vr.Name, vr.Type, vr.Value))
			} else {
				b.WriteString(fmt.Sprintf("const %s = %s\n", vr.Name, vr.Value))
			}

		case VarLegacy:
			b.WriteString(fmt.Sprintf("%s:%s = %s\n", vr.Scope, vr.Name, vr.Value))
		}
	}

	b.WriteByte('\n')
	return b.String()
}
