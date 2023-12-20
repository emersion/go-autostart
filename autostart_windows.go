package autostart

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var startupKeyPath = `Software\Microsoft\Windows\CurrentVersion\Run`

func (a *App) path() string {
	return filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup", a.Name+".lnk")
}

func (a *App) IsEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, startupKeyPath, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()

	_, _, err = k.GetStringValue(a.Name)
	return err == nil
}

func (a *App) Enable() error {
	exePath := a.Exec[0]

	args := strings.Join(a.Exec[1:], " ")

	// Now fullPath will have the environment variables expanded
	fullPath := fmt.Sprintf("\"%s\" %s", exePath, args)

	fullPath = os.ExpandEnv(fullPath)

	k, err := registry.OpenKey(registry.CURRENT_USER, startupKeyPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.SetStringValue(a.Name, fullPath)
	if err != nil {
		return errors.New(fmt.Sprintf("autostart: cannot create registry value '%s': %v", a.Name, err))
	}
	return nil
}

func (a *App) Disable() error {
	k, err := registry.OpenKey(registry.CURRENT_USER, startupKeyPath, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	err = k.DeleteValue(a.Name)
	if err != nil {
		return errors.New(fmt.Sprintf("autostart: cannot delete registry value '%s': %v", a.Name, err))
	}
	return nil
}
