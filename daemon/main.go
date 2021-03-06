// Copyright 2016-2018 Authors of Cilium
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

import (
	"fmt"
	"os"
	"path"

	"github.com/cilium/cilium/cilium/cmd"
	"github.com/cilium/cilium/monitor"
)

func main() {
	base := path.Base(os.Args[0])

	switch base {
	case "cilium-agent":
		daemonMain()
	case "cilium":
		cmd.Execute()
	case "cilium-node-monitor":
		monitor.Execute()
	default:
		panic(fmt.Sprintf("Invalid executable name: %s. Only \"cilium-agent\", "+
			"\"cilium\" or \"cilium-node-monitor\" is supported.", base))
	}
}
