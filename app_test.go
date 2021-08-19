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

package tgik

import (
	"os"
	"testing"

	"github.com/kris-nova/naml"

	"github.com/kris-nova/logger"
)

// TestMain will bootstrap and tear down our testing cluster.
func TestMain(m *testing.M) {
	err := naml.TestClusterStart()
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	q := m.Run()
	naml.TestClusterStop()
	os.Exit(q)
}

// TestApp is an example integration test that can be used to
// install and uninstall a sample application in Kubernetes.
func TestApp(t *testing.T) {
	client, err := naml.ClientFromPath(naml.TestClusterKubeConfigPath())
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}
	app := NewApp("default", "sample-app")
	err = app.Install(client)
	if err != nil {
		t.Errorf("unable to install sample-app: %v", err)
	}
	err = app.Uninstall(client)
	if err != nil {
		t.Errorf("unable to uninstall sample-app: %v", err)
	}
}
