package autostart

// #cgo LDFLAGS: -lole32 -luuid
/*
#define WIN32_LEAN_AND_MEAN
#include <windows.h>

int CreateShortcut(char *shortcutA, char *path, char *args);
*/
import "C"

import (
	"errors"
	"os"
	"path/filepath"
)

var startupDir string

func init() {
	startupDir = filepath.Join(os.Getenv("USERPROFILE"), "Start Menu", "Programs", "Startup")
}

func (a *App) path() string {
	return filepath.Join(startupDir, a.Name+".lnk")
}

func (a *App) IsEnabled() bool {
	_, err := os.Stat(a.path())
	return !os.IsNotExist(err)
}

func (a *App) Enable() error {
	path := a.Exec[0]
	args := quote(a.Exec[1:])

	res := C.CreateShortcut(C.CString(a.path()), C.CString(path), C.CString(args))
	if res == 0 {
		return errors.New("autostart: cannot create shortcut")
	}
	return nil
}

func (a *App) Disable() error {
	return os.Remove(a.path())
}
