//
// Copyright © 2021 Kris Nóva <kris@nivenly.com>
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
//
//   ███╗   ██╗ █████╗ ███╗   ███╗██╗
//   ████╗  ██║██╔══██╗████╗ ████║██║
//   ██╔██╗ ██║███████║██╔████╔██║██║
//   ██║╚██╗██║██╔══██║██║╚██╔╝██║██║
//   ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗
//   ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝
//

package naml

import (
	"fmt"
	"runtime/debug"
)

// Version is this specific version on naml
var Version string

func Banner() {
	if Version == "" {
		if build, ok := debug.ReadBuildInfo(); ok {
			Version = build.Main.Version[:5]
		}
	}
	fmt.Printf("\n+---------------------------------------------+\n")
	fmt.Printf("|    ███╗   ██╗ █████╗ ███╗   ███╗██╗         |\n")
	fmt.Printf("|    ████╗  ██║██╔══██╗████╗ ████║██║         |\n")
	if Version != "" {
		fmt.Printf("|    ██╔██╗ ██║███████║██╔████╔██║██║ v%s  |\n", Version)
	} else {
		fmt.Printf("|    ██╔██╗ ██║███████║██╔████╔██║██║      |\n")
	}
	fmt.Printf("|    ██║╚██╗██║██╔══██║██║╚██╔╝██║██║         |\n")
	fmt.Printf("|    ██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗    |\n")
	fmt.Printf("|    ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    |\n")
	fmt.Printf("|       Not Another Markup Language           |\n")
	fmt.Printf("|       Kris Nóva <kris@nivenly.com>          |\n")
	fmt.Printf("+---------------------------------------------+\n\n")

}
