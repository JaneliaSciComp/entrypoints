package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var scriptDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "entrypoints [APP] [ARG...]",
	Short: "Exposes multiple entrypoint apps for an OCI container",
	Long: `This command can be included as the entrypoint of an OCI (e.g. Docker) container in order to allow the user ` +
		`to call multiple entrypoint scripts inside of the container. When called without an app, it lists all of the ` +
		`available apps.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			app := args[0]
			appPath := scriptDir + "/" + app + ".sh"
			err := runCommand(appPath, args[1:]...)
			if err != nil {
				os.Exit(1)
			}
			os.Exit(0)
		}

		fmt.Println("Available apps:")

		err := filepath.Walk(scriptDir, func(path string, info os.FileInfo, err error) error {
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&scriptDir, "scriptDir", "s", "", "Directory containing entry scripts")
	rootCmd.MarkPersistentFlagRequired("scriptDir")
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
