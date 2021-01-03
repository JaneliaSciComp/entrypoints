package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "List available apps",
	Long:  `Lists all apps which can be invoked`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available apps:")

		err := filepath.Walk(ScriptDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			filename := info.Name()
			ext := filepath.Ext(filename)
			if ext == ".sh" {
				appName := strings.TrimSuffix(filename, ext)
				fmt.Println(appName)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)
}
