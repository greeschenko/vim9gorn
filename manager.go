package vim9gorn

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ExternalPlugin struct {
	Repo string
	Dir  string
}

func NewExternalPlugin(repo string) *ExternalPlugin {
	parts := strings.Split(repo, "/")
	name := parts[len(parts)-1]
	return &ExternalPlugin{
		Repo: repo,
		Dir:  filepath.Join(".vim", "pack", "plugins", "opt", name),
	}
}

type PluginManager struct {
	Plugins []ExternalPlugin
	VimDir  string
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		Plugins: make([]ExternalPlugin, 0),
		VimDir:  ".vim",
	}
}

func (m *PluginManager) Add(repo string) *PluginManager {
	m.Plugins = append(m.Plugins, *NewExternalPlugin(repo))
	return m
}

func (m *PluginManager) FetchAll() error {
	for _, p := range m.Plugins {
		if err := m.fetchPlugin(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *PluginManager) fetchPlugin(p ExternalPlugin) error {
	targetDir := filepath.Join(m.VimDir, p.Dir)

	if _, err := os.Stat(targetDir); err == nil {
		fmt.Printf("Plugin %s already exists, skipping\n", p.Repo)
		return nil
	}

	parentDir := filepath.Dir(targetDir)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return err
	}

	cloneURL := fmt.Sprintf("https://github.com/%s.git", p.Repo)
	cmd := exec.Command("git", "clone", "--depth", "1", cloneURL, targetDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Cloning %s...\n", p.Repo)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to clone %s: %w", p.Repo, err)
	}

	return nil
}

func (m *PluginManager) UpdateAll() error {
	for _, p := range m.Plugins {
		if err := m.updatePlugin(p); err != nil {
			return err
		}
	}
	return nil
}

func (m *PluginManager) updatePlugin(p ExternalPlugin) error {
	targetDir := filepath.Join(m.VimDir, p.Dir)

	if _, err := os.Stat(targetDir); err != nil {
		return fmt.Errorf("plugin %s not found, run FetchAll first", p.Repo)
	}

	cmd := exec.Command("git", "-C", targetDir, "pull")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Updating %s...\n", p.Repo)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to update %s: %w", p.Repo, err)
	}

	return nil
}

func (m *PluginManager) Generate() string {
	var b strings.Builder
	b.WriteString("# External Plugins\n")
	b.WriteString("# Run: go run . to fetch plugins\n\n")

	for _, p := range m.Plugins {
		b.WriteString(fmt.Sprintf("# %s -> %s\n", p.Repo, p.Dir))
	}

	return b.String()
}
