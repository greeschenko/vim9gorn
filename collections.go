package vim9gorn

import (
	"fmt"
	"strings"
)

// =====================
// List — represents a Vim9 list
// Example: [1, 2, 3] or ["a","b","c"]
// =====================
type List struct {
	Items []string
}

// NewList creates a new List
func NewList(items ...string) *List {
	return &List{
		Items: items,
	}
}

// Add appends an item to the list
func (l *List) Add(item string) *List {
	l.Items = append(l.Items, item)
	return l
}

// Generate returns Vim9 list syntax
func (l *List) Generate() string {
	return "[" + strings.Join(l.Items, ", ") + "]"
}

// =====================
// Dict — represents a Vim9 dictionary
// Example: {"a":1, "b":2}
// =====================
type Dict struct {
	Items map[string]string
}

// NewDict creates a new Dict
func NewDict() *Dict {
	return &Dict{
		Items: make(map[string]string),
	}
}

// Set adds or updates a key/value pair
func (d *Dict) Set(key, value string) *Dict {
	d.Items[key] = value
	return d
}

// Generate returns Vim9 dictionary syntax
func (d *Dict) Generate() string {
	if len(d.Items) == 0 {
		return "{}"
	}

	var parts []string
	for k, v := range d.Items {
		parts = append(parts, fmt.Sprintf("\"%s\": %s", k, v))
	}

	return "{" + strings.Join(parts, ", ") + "}"
}

// =====================
// Values — helper for iterating over dict values
// Example: for v in values(dict)
// =====================
func Values(dict *Dict) string {
	return "values(" + dict.Generate() + ")"
}

// =====================
// Keys — helper for iterating over dict keys
// Example: for k in keys(dict)
// =====================
func Keys(dict *Dict) string {
	return "keys(" + dict.Generate() + ")"
}

func Items(dict *Dict) string {
	return "items(" + dict.Generate() + ")"
}

