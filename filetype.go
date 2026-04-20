package vim9gorn

import (
	"strings"
)

type FiletypeDetect struct {
	Pattern  string
	Filetype string
	Command  string
}

func NewFiletypeDetect(pattern string) *FiletypeDetect {
	return &FiletypeDetect{
		Pattern: pattern,
	}
}

func (f *FiletypeDetect) SetFiletype(ft string) *FiletypeDetect {
	f.Filetype = ft
	return f
}

func (f *FiletypeDetect) SetCommand(cmd string) *FiletypeDetect {
	f.Command = cmd
	return f
}

func (f *FiletypeDetect) Generate() string {
	var b strings.Builder
	b.WriteString("autocmd")
	b.WriteString(" filetypedetect")

	if f.Pattern != "" {
		b.WriteString(" ")
		b.WriteString(f.Pattern)
	}

	if f.Command != "" {
		b.WriteString(" ")
		b.WriteString(f.Command)
	} else if f.Filetype != "" {
		b.WriteString(" setf ")
		b.WriteString(f.Filetype)
	}

	return b.String()
}

type FiletypePlugin struct {
	Filetype string
	Settings []string
}

func NewFiletypePlugin(ft string) *FiletypePlugin {
	return &FiletypePlugin{
		Filetype: ft,
		Settings: make([]string, 0),
	}
}

func (f *FiletypePlugin) AddSetting(setting string) *FiletypePlugin {
	f.Settings = append(f.Settings, setting)
	return f
}

func (f *FiletypePlugin) Generate() string {
	var b strings.Builder
	b.WriteString("# ftplugin/")
	b.WriteString(f.Filetype)
	b.WriteString(".vim settings\n")

	for _, s := range f.Settings {
		b.WriteString(s)
		b.WriteByte('\n')
	}

	return b.String()
}
