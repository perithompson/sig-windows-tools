/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"burrito/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var mc utils.BurritoConfig
		if err := viper.Unmarshal(&mc); err != nil {
			fmt.Println(err)
		}
		d := make(map[string]string)
		for _, c := range mc.Components {
			if c.Variable_Src != "" {
				r, _ := http.NewRequest("GET", c.Source, nil)
				val := fmt.Sprintf("http://%s/%s/%s/%s", mc.HttpRoot, mc.FsDir, c.Name, path.Base(r.URL.Path))
				d[c.Variable_Src] = val
			}
			if c.Variable_Checksum != "" {
				d[c.Variable_Checksum] = c.Sha256
			}

		}
		json, err := json.MarshalIndent(d, "\r", "   ")
		if err != nil {
			fmt.Printf("Unable to Marshall to json: %e", err)
		}
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
