//go:build (!cgo || nocgo) && windows

package mpv

import (
	"fmt"

	"golang.org/x/sys/windows"
)

const (
	LibName = "libmpv.dll"
)

func loadLibrary() uintptr {
	handle, err := windows.LoadLibrary(LibName)
	if err != nil {
		return 0
	}

	return uintptr(handle)
}
