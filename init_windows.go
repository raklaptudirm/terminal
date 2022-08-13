//go:build windows
// +build windows

package terminal

import (
	"golang.org/x/sys/windows"
)

func (t *Terminal) init() {
	handle := windows.Handle(t.Fd())
	var originalMode uint32

	windows.GetConsoleMode(handle, &originalMode)
	// enable virtual terminal processing to properly parse ansi escapes
	windows.SetConsoleMode(handle, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
