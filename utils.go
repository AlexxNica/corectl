// Copyright 2015 - António Meireles  <antonio.meireles@reformi.st>
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

package main

import "github.com/spf13/cobra"

var (
	utilsCmd = &cobra.Command{
		Use:   "utils",
		Short: "some developer focused utilies",
		Run:   utilsCommand,
	}
	manCmd = &cobra.Command{
		Use:   "mkMan",
		Short: "generates man pages",
		Run:   manCommand,
	}
	mkdownCmd = &cobra.Command{
		Use:   "mkMkdown",
		Short: "generates Markdown documentation",
		Run:   mkdownCommand,
	}
)

func utilsCommand(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func manCommand(cmd *cobra.Command, args []string) {
	RootCmd.GenManTree("coreos", SessionContext.pwd+"/documentation/man/")
}

func mkdownCommand(cmd *cobra.Command, args []string) {
	cobra.GenMarkdownTree(RootCmd,
		SessionContext.pwd+"/documentation/markdown/")
}

func init() {
	if SessionContext.debug {
		utilsCmd.AddCommand(manCmd)
		utilsCmd.AddCommand(mkdownCmd)
		RootCmd.AddCommand(utilsCmd)
	}
}
