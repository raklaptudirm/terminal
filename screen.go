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

// SaveScreen saves the state of the current screen of the terminal.
func (t *Terminal) SaveScreen() {
	t.PrintCSI("?47h")
}

// RestoreScreen restores the saved screen back state into the terminal.
func (t *Terminal) RestoreScreen() {
	t.PrintCSI("?47l")
}

// EnableAlternateBuffer enables the alternate screen buffer.
func (t *Terminal) EnableAlternateBuffer() {
	t.PrintCSI("?1049h")
}

// DisableAlternateBuffer disables the alternate screen buffer.
func (t *Terminal) DisableAlternateBuffer() {
	t.PrintCSI("?1049l")
}
