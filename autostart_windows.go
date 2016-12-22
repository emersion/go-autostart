package autostart

import (
	"os"
	"path/filepath"
)

var startupDir string

func init() {
	startupDir = filepath.Join(os.Getenv("USERPROFILE"), "Start Menu", "Programs", "Startup")
}

func (a *App) path() string {
	return filepath.Join(startupDir, a.Name+".exe")
}

func (a *App) IsEnabled() bool {
	_, err := os.Stat(a.path())
	return !os.IsNotExist(err)
}

func (a *App) Enable() error {
	return os.Link(a.Exec[0], filepath.Join(startupDir, a.Name+".exe"))
}

func (a *App) Disable() error {
	return os.Remove(filepath.Join(startupDir, a.Name+".exe"))
}
