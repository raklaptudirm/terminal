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

package ansi

func (t *Terminal) MoveCursorHome() {
	t.PrintCSI("H")
}

func (t *Terminal) MoveCursor(line, column int) {
	t.PrintCSI("%d;%dH", line, column)
}

func (t *Terminal) MoveCursorUp(lines int) {
	t.PrintCSI("%dA", lines)
}

func (t *Terminal) MoveCursorDown(lines int) {
	t.PrintCSI("%dB", lines)
}

func (t *Terminal) MoveCursorRight(columns int) {
	t.PrintCSI("%dC", columns)
}

func (t *Terminal) MoveCursorLeft(columns int) {
	t.PrintCSI("%dD", columns)
}

func (t *Terminal) MoveCursorToColumn(column int) {
	t.PrintCSI("%dG", column)
}

func (t *Terminal) SaveCursorPosition() {
	t.PrintCSI("s")
}

func (t *Terminal) RestoreCursorPosition() {
	t.PrintCSI("u")
}

func (t *Terminal) HideCursor() {
	t.PrintCSI("?25l")
}

func (t *Terminal) ShowCursor() {
	t.PrintCSI("25h")
}
