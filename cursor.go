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

// MoveCursorHome moves the cursor to it's home(0,0) position.
func (t *Terminal) MoveCursorHome() {
	t.PrintCSI("H")
}

// MoveCursor moves the cursor to the given line and column.
func (t *Terminal) MoveCursor(line, column int) {
	t.PrintCSI("%d;%dH", line, column)
}

// MoveCursorUp moves the cursor up by the given amount of lines.
func (t *Terminal) MoveCursorUp(lines int) {
	t.PrintCSI("%dA", lines)
}

// MoveCursorDown moves the cursor down by the given amount of lines.
func (t *Terminal) MoveCursorDown(lines int) {
	t.PrintCSI("%dB", lines)
}

// MoveCursorRight moves the cursor right by the given amount of columns.
func (t *Terminal) MoveCursorRight(columns int) {
	t.PrintCSI("%dC", columns)
}

// MoveCursorLeft moves te cursor left by the given amount of columns.
func (t *Terminal) MoveCursorLeft(columns int) {
	t.PrintCSI("%dD", columns)
}

// MoveCursorToColumn moves the cursor to the given column.
func (t *Terminal) MoveCursorToColumn(column int) {
	t.PrintCSI("%dG", column)
}

// SaveCursorPosition saves the current position of the cursor.
func (t *Terminal) SaveCursorPosition() {
	t.PrintCSI("s")
}

// Restore cursor position restores the cursor to the most recently saved
// position.
func (t *Terminal) RestoreCursorPosition() {
	t.PrintCSI("u")
}

// HideCursor makes the cursor invisible on the screen.
func (t *Terminal) HideCursor() {
	t.PrintCSI("?25l")
}

// ShowCursor makes the cursor visible on the screen.
func (t *Terminal) ShowCursor() {
	t.PrintCSI("25h")
}
