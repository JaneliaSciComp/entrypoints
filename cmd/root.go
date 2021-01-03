package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ScriptDir represents the configured script directory containing app scripts
var ScriptDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "entrypoints [APP] [ARG...]",
	Short: "Exposes multiple entrypoint apps for an OCI container",
	Long: `This command can be included as the entrypoint of an OCI (e.g. Docker) container in order to allow the user ` +
		`to call multiple entrypoint scripts inside of the container. When called without an app, it lists all of the ` +
		`available apps.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
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
	rootCmd.PersistentFlags().StringVarP(&ScriptDir, "scriptDir", "s", "", "Directory containing entry scripts")
	rootCmd.MarkPersistentFlagRequired("scriptDir")
}
