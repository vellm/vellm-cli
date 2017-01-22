// Copyright © 2017 Henrik Fricke <hello@vellm.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/vellm/vellm-cli/cmd"

const (
	// VERSION will be set in deploy process
	VERSION = "latest"
)

func main() {
	cmd.Execute(VERSION)
}
