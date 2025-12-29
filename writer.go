package vim9gorn

import (
	"os"
)

// FileWriter — інтерфейс для запису в файл (DIP)
type FileWriter interface {
	Write(path string, content string) error
}

// DefaultFileWriter — стандартна реалізація через os.WriteFile
type DefaultFileWriter struct{}

func (w *DefaultFileWriter) Write(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
