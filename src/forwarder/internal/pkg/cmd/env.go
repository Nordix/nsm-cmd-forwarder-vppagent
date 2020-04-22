// Copyright (c) 2020 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spiffe/go-spiffe/workload"

	"github.com/networkservicemesh/sdk/pkg/tools/flags"
)

func init() {
	cmd := envCmd
	rootCmd.AddCommand(cmd)
	Flags(cmd.Flags())
	cobra.OnInitialize(flags.FromEnv(flags.EnvPrefix, flags.EnvReplacer, cmd.Flags()))
}

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Dumps env for current forwarder settings suitable for evaling",
	Long:  `Dumps env for current forwarder settings suitable for evaling`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Value.Type() != "bool" {
				fmt.Printf("%s=%q\n", flags.KeyToEnvVariable(f.Name), f.Value.String())
			}
		})
		fmt.Printf("%s=%q\n", workload.SocketEnv, os.Getenv(workload.SocketEnv))
	},
}