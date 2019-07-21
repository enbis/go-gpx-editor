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
	"fmt"
	"log"

	"github.com/enbis/go-gpx-editor/editorgpx"
	"github.com/spf13/cobra"
)

// keepfromCmd represents the keepfrom command
var keepfromCmd = &cobra.Command{
	Use:   "keepfrom",
	Short: "Keeps unmodified the GPX trace from the start time setted via CLI, the rest of the lap is cutted",
	Long: `The command takes two parameters:
	
	-p for the path of the GPX file
	-t to set the start time in format 00h00m00s
	
	For example. 
	go run main.go keepfrom -p=/media/user/DATA/activity.gpx -t=0h6m0s`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("keepfrom called")
		//hh:mm:ss
		filepath, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal("error filepath empty")
		}
		time, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal(err)
		}

		editorgpx.KeepFrom(filepath, time)
	},
}

func init() {
	rootCmd.AddCommand(keepfromCmd)
	keepfromCmd.Flags().StringP("path", "p", "", "filepath")
	keepfromCmd.Flags().StringP("time", "t", "", "Keep gpx from this time")

}
