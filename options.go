package vim9gorn

import (
	"fmt"
	"strings"
)

type Options struct {
	data map[string]any
}

// NewOptions створює новий контейнер для options
func NewOptions() *Options {
	return &Options{
		data: make(map[string]any),
	}
}

// Set додає/оновлює option
func (o *Options) Set(name string, value any) *Options {
	o.data[name] = value
	return o
}

// Generate повертає блок налаштувань у форматі vim9script
func (o *Options) Generate() string {
	if len(o.data) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("# === Options ===\n")
	for k, v := range o.data {
		line, _ := formatOption(k, v)
		b.WriteString(line + "\n")
	}
	b.WriteByte('\n')
	return b.String()
}

func formatOption(name string, value any) (string, error) {
	switch v := value.(type) {
	case bool:
		if v {
			return fmt.Sprintf("set %s", name), nil
		}
		return fmt.Sprintf("set no%s", name), nil
	case int:
		return fmt.Sprintf("set %s=%d", name, v), nil
	case string:
		return fmt.Sprintf("set %s=%s", name, v), nil
	default:
		return "", fmt.Errorf("unsupported option type for %s", name)
	}
}
