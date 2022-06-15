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

func (t *Terminal) EraseTillScreenEnd() {
	t.PrintCSI("0J")
}

func (t *Terminal) EraseTillScreenBeginning() {
	t.PrintCSI("1J")
}

func (t *Terminal) EraseScreen() {
	t.PrintCSI("2J")
}

func (t *Terminal) EraseTillLineEnd() {
	t.PrintCSI("0K")
}

func (t *Terminal) EraseTillLineBeginning() {
	t.PrintCSI("1K")
}

func (t *Terminal) EraseLine() {
	t.PrintCSI("2K")
}
