/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/egon12/ghr/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rejectCmd represents the reject command
var rejectCmd = &cobra.Command{
	Use:   "reject messages",
	Short: "Request changes in the commit",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := app.Config{}
		_ = viper.Unmarshal(&cfg)

		a := app.InitApp(cfg)
		err := a.ReviewProcess.RequestChanges(args[0])
		if err != nil {
			os.Stderr.Write([]byte(err.Error()))
		}
	},
}

func init() {
	rootCmd.AddCommand(rejectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rejectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rejectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
