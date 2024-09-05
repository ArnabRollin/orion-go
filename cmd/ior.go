/*
Copyright © 2024 Arnab Phukan <iamarnab.phukan@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/Solarcode-org/orion-go/lib"
	"github.com/spf13/cobra"
	"log"
)

// iorCmd represents the ior command
var iorCmd = &cobra.Command{
	Use:   "ior",
	Short: "Run an individual Orion file",
	Long: `Run an individual Orion file.

Example:
	orion ior main.or`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("You must specify a Orion file.")
		}

		ast, err := lib.Parse(args[0])
		if err != nil {
			log.Fatalf("Failed to parse file: %s", args[0])
		}

		err = lib.Run(ast)
		if err != nil {
			log.Fatalf("Failed to run file: %s", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(iorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
