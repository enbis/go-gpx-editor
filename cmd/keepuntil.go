// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"log"

	"github.com/enbis/go-gpx-editor/editorgpx"
	"github.com/spf13/cobra"
)

// keepuntilCmd represents the keepuntil command
var keepuntilCmd = &cobra.Command{
	Use:   "keepuntil",
	Short: "Keeps unmodified the GPX trace until the end time setted via CLI, the rest of the lap is cutted",
	Long: `The command takes two parameters:
	
	-p for the path of the GPX file
	-t to set the end time in format 00h00m00s
	
	For example. 
	go run main.go keepuntil -p=/media/user/DATA/activity.gpx -t=1h6m0s`,

	Run: func(cmd *cobra.Command, args []string) {

		//hh:mm:ss
		filepath, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}
		time, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal(err)
		}

		editorgpx.KeepUntil(filepath, time)
	},
}

func init() {
	rootCmd.AddCommand(keepuntilCmd)
	keepuntilCmd.Flags().StringP("path", "p", "", "filepath")
	keepuntilCmd.Flags().StringP("time", "t", "", "Keep gpx until this time")
}
