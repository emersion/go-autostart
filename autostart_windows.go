package autostart

import (
	"os"
	"path/filepath"
	"strings"
)

var startupDir string

func init() {
	startupDir = filepath.Join(os.Getenv("USERPROFILE"), "Start Menu", "Programs", "Startup")
}

func (a *App) path() string {
	return filepath.Join(startupDir, a.Name + ".bat")
}

func (a *App) IsEnabled() bool {
	_, err := os.Stat(a.path())
	return !os.IsNotExist(err)
}

func (a *App) Enable() error {
	s := "start " + strings.Join(a.Exec, " ") + "\r\n"

	f, err := os.Create(a.path())
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(s))
	return err
}

func (a *App) Disable() error {
	return os.Remove(a.path())
}
