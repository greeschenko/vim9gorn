package vim9gorn

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	PluginTypeOpt   = "opt"
	PluginTypeStart = "start"
)

type Plugin struct {
	Name   string
	Repo   string
	Type   string
	Source string
}

func NewPlugin(name, ptype string) *Plugin {
	return &Plugin{
		Name: name,
		Type: ptype,
	}
}

func (p *Plugin) SetRepo(repo string) *Plugin {
	p.Repo = repo
	return p
}

func (p *Plugin) GetPath() string {
	return filepath.Join(".vim", "pack", "plugins", p.Type, p.Name)
}

type PluginManifest struct {
	OutputDir string
	Plugins   []Plugin
}

func NewManifest() *PluginManifest {
	return &PluginManifest{
		OutputDir: ".vim",
		Plugins:   make([]Plugin, 0),
	}
}

func (m *PluginManifest) Add(p Plugin) *PluginManifest {
	m.Plugins = append(m.Plugins, p)
	return m
}

type Directory struct {
	Path string
}

func (d Directory) Generate() string {
	return d.Path
}

func (m *PluginManifest) Generate() string {
	var b strings.Builder
	b.WriteString("# === Plugin Structure ===\n")
	b.WriteString("# vim9gorn plugin scaffold\n")

	for _, p := range m.Plugins {
		b.WriteString(fmt.Sprintf("# Plugin: %s\n", p.Name))
		b.WriteString(fmt.Sprintf("# Path: %s\n", p.GetPath()))

		dirs := []string{
			filepath.Join(p.GetPath(), "plugin"),
			filepath.Join(p.GetPath(), "autoload"),
			filepath.Join(p.GetPath(), "ftplugin"),
			filepath.Join(p.GetPath(), "syntax"),
			filepath.Join(p.GetPath(), "doc"),
		}
		for _, dir := range dirs {
			b.WriteString(fmt.Sprintf("mkdir -p %s\n", dir))
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func (m *PluginManifest) Forge(path string, writer FileWriter) error {
	content := m.Generate()
	return writer.Write(path, content)
}

func (m *PluginManifest) CreateDirectories(pluginsDir string) error {
	for _, p := range m.Plugins {
		dirs := []string{
			filepath.Join(pluginsDir, p.GetPath(), "plugin"),
			filepath.Join(pluginsDir, p.GetPath(), "autoload"),
			filepath.Join(pluginsDir, p.GetPath(), "ftplugin"),
			filepath.Join(pluginsDir, p.GetPath(), "syntax"),
			filepath.Join(pluginsDir, p.GetPath(), "doc"),
		}
		for _, dir := range dirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
	}
	return nil
}
