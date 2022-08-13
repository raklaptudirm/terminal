// Copyright © 2022 Rak Laptudirm <raklaptudirm@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terminal

import (
	"fmt"
	"os"
)

// Constants representing escape headers.
const (
	Esc = "\x1b"
	Csi = Esc + "["
)

// New creates an instance of a Terminal from the given *os.File.
func New(w *os.File) *Terminal {
	term := Terminal{w}
	term.init()

	return &term
}

// Terminal represents a terminal supporting ansi escape sequences.
type Terminal struct {
	*os.File
}

// Print acts like the fmt.Fprint function on a Terminal io.Writer.
func (t *Terminal) Print(a ...any) (int, error) {
	return fmt.Fprint(t, a...)
}

// Println acts like the fmt.Fprintln function on a Terminal io.Writer.
func (t *Terminal) Println(a ...any) (int, error) {
	return fmt.Fprintln(t, a...)
}

// Printf acts like the fmt.Fprintf function on a Terminal io.Writer.
func (t *Terminal) Printf(format string, a ...any) (int, error) {
	return fmt.Fprintf(t, format, a...)
}

// PrintCSI prints the given escape sequence appended to a CSI(ESC [)
// escape header.
func (t *Terminal) PrintCSI(format string, a ...any) {
	format = Csi + format
	t.Printf(format, a...)
}
