/*
Copyright © 2020 Laurent VAYLET <laurent.vaylet@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/xanzy/go-gitlab"
)

// rootRunnerJobCmd represents the rootRunnerJob command
var rootRunnerJobCmd = &cobra.Command{
	Use:   "rootRunnerJob",
	Short: "Manage Root Runner Jobs",
	Long:  `Simplify management of Root Runner Jobs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootRunnerJob called")

		// TODO Do something more relevant than listing users.
		// Use GitLab client to list users
		users, _, err := gitlabClient.Users.ListUsers(&gitlab.ListUsersOptions{})
		if err == nil {
			fmt.Println("Found", len(users), "users.")
		}
	},
}

func init() {
	rootCmd.AddCommand(rootRunnerJobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootRunnerJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootRunnerJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
