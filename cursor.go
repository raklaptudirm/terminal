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
	t.PrintEsc("[H")
}

func (t *Terminal) MoveCursor(line, column int) {
	t.PrintEsc("[%d;%dH", line, column)
}

func (t *Terminal) MoveCursorUp(lines int) {
	t.PrintEsc("[%dA", lines)
}

func (t *Terminal) MoveCursorDown(lines int) {
	t.PrintEsc("[%dB", lines)
}

func (t *Terminal) MoveCursorRight(columns int) {
	t.PrintEsc("[%dC", columns)
}

func (t *Terminal) MoveCursorLeft(columns int) {
	t.PrintEsc("[%dD", columns)
}

func (t *Terminal) MoveCursorToColumn(column int) {
	t.PrintEsc("[%dG", column)
}

func (t *Terminal) SaveCursorPosition() {
	t.PrintEsc(" 7")
}

func (t *Terminal) RestoreCursorPosition() {
	t.PrintEsc(" 8")
}

func (t *Terminal) HideCursor() {
	t.PrintEsc("[?25l")
}

func (t *Terminal) ShowCursor() {
	t.PrintEsc("[25h")
}
