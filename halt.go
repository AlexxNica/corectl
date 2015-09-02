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

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	killCmd = &cobra.Command{
		Use:     "kill",
		Aliases: []string{"stop", "halt"},
		Short:   "halts a running VM",
		Run:     killCommand,
	}
)

func killCommand(cmd *cobra.Command, args []string) {
	SessionContext.canRun()
	if len(args) == 0 {
		log.Println("nothing to kill...")
		return
	}
	for _, v := range args {
		vm, err := findVM(v)
		if err == nil {
			if err := vm.runCommand([]string{"id", ";", "sync"}); err != nil {
				// ssh messed up for some reason
				if vm.isAlive() {
					if p, err := os.FindProcess(vm.Pid); err == nil {
						log.Println("hard kill...", err)
						if err := p.Kill(); err != nil {
							log.Fatalln(err)
						}
					}
				}
			} else {
				// will work. bellow returns an error, by design,
				// that we can safely ignore (because 'id' above worked)
				_ = vm.runCommand([]string{"sudo", "halt"})
			}
		}
		break
	}
}

func init() {
	RootCmd.AddCommand(killCmd)
}
