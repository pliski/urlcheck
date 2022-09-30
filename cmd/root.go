/*
Copyright © 2022 pliski <pliski@pli.ski>

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
	"errors"
	"fmt"
	"net/url"
	"os"

	action "urlcheck/action"
	cliState "urlcheck/cliState"
	model "urlcheck/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	Verbose     = false
	authorFlag  bool
	timeout     uint // seconds
	statusFlag  bool
	urlFilename string
	uriList     model.UriList
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{

	Use:   "urlcheck [uri ...]",
	Short: "checks url/s",
	Long: `Tries to connect to servers from the URLs in input and print the responses. 
	Accepts:
	- a single url, 
	- a filename containing one url per line `,
	// - by default search in ~/.config/urlcheck/sites.txt`,

	Args: func(cmd *cobra.Command, args []string) error {
		for i := 0; i < len(args); i++ {
			_, err := url.ParseRequestURI(args[i])
			if err != nil {
				errMsg := fmt.Sprintln("Invalid URL:", args[i])
				return errors.New(errMsg)
			}
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Parsing input ... ")
		if urlFilename != "" {
			err := uriList.UriListFromFile(urlFilename)
			if err != nil {
				return err
			}
		}

		if len(args) > 0 {
			uriList.UriListFromArgs(args)
		}

		if authorFlag {
			fmt.Fprintln(os.Stdout, "Author: pliski <pliski@pli.ski>")
			return nil
		}

		if len(uriList.Entry) == 0 {
			return errors.New("nothing to do")
		}

		if statusFlag {
			if action.IsStatusOK(uriList.Entry[0], timeout) {
				os.Exit(0)
			}
			os.Exit(-1)
		}

		model := cliState.NewModel(uriList.Entry)
		err := tea.NewProgram(model).Start()
		return err

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func init() {
	uriList = model.NewUriList()
	rootCmd.PersistentFlags().UintVarP(&timeout, "timeout", "t", 10, "http call timeout (in seconds)")
	rootCmd.PersistentFlags().BoolVarP(&statusFlag, "status", "s", false, "minimal check for http status code (intended for inline use)")
	rootCmd.PersistentFlags().BoolVarP(&authorFlag, "author", "a", false, "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&urlFilename, "filename", "f", "", "name of the file containing the list of URLs to check.")
}
