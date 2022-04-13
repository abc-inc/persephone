// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strings"

	cmd "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
)

var version string

var DriverVersionCmd = &cobra.Command{
	Use:         ":driver-version",
	Aliases:     []string{"driver-version"},
	Short:       "Print version of the Neo4j Driver used.",
	Annotations: map[string]string{},
	Run:         func(cmd *cobra.Command, args []string) { driverVersionCmd() },
}

var VersionCmd = &cobra.Command{
	Use:         ":version",
	Aliases:     []string{"version"},
	Short:       "Print version information.",
	Annotations: cmd.Annotate(cmd.Offline),
	Run:         func(cmd *cobra.Command, args []string) { versionCmd() },
}

func driverVersionCmd() {
	ua := strings.Split(neo4j.UserAgent, "/")
	fmt.Println("Neo4j Driver", ua[len(ua)-1])
}

func versionCmd() {
	fmt.Println("persephone ", version)
}
