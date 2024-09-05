package autostart

// #cgo LDFLAGS: -lole32 -luuid
/*
#define WIN32_LEAN_AND_MEAN
#include <stdint.h>
#include <stdlib.h>
#include <windows.h>

uint64_t CreateShortcut(char *shortcutA, char *path, char *args);
*/
import "C"

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

var startupDir string

func init() {
	startupDir = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming", "Microsoft", "Windows", "Start Menu", "Programs", "Startup")
}

func (a *App) path() string {
	return filepath.Join(startupDir, a.Name+".lnk")
}

func (a *App) IsEnabled() bool {
	_, err := os.Stat(a.path())
	return err == nil
}

func (a *App) Enable() error {
	path := a.Exec[0]
	args := strings.Join(a.Exec[1:], " ")

	if err := os.MkdirAll(startupDir, 0777); err != nil {
		return err
	}

	cShortcutPath := C.CString(a.path())
	defer C.free(unsafe.Pointer(cShortcutPath))
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cArgs := C.CString(args)
	defer C.free(unsafe.Pointer(cArgs))

	res := C.CreateShortcut(C.CString(a.path()), C.CString(path), C.CString(args))
	if res != 0 {
		return fmt.Errorf("autostart: cannot create shortcut '%s' error code: 0x%.8x", a.path(), res)
	}

	return nil
}

func (a *App) Disable() error {
	return os.Remove(a.path())
}
