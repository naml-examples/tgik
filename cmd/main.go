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

package main

import (
	"fmt"
	"os"


	dba "github.com/naml-examples/tgik/dashboardauth"
	"github.com/kris-nova/naml"
	tgik "github.com/naml-examples/tgik"
)

func main() {
	// Load the application into the NAML registery
	// Note: naml.Register() can be used multiple times.
	naml.Register(tgik.NewApp("tgik-kube-dashboard", "Kubernetes dashboard for TGIK"))
	naml.Register(dba.NewApp("tgik-kube-dashboard-roles-thingies", "These are extra role thingies required for dashboard"))

	// Run the generic naml command line program with
	// the application loaded.
	err := naml.RunCommandLine()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
