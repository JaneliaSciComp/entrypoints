package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// invokeCmd represents the invoke command
var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Invoke the given app",
	Long:  `Invoke the given containerized app with the provided args`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			invokeApp(args[0], args[1:])
		}
	},
}

func init() {
	rootCmd.AddCommand(invokeCmd)
	// Disable parsing because we want to pass through flags to the containerized application
	invokeCmd.DisableFlagParsing = true
}

func invokeApp(app string, args []string) {

	appPath := ScriptDir + "/" + app + ".sh"

	info, err := os.Stat(appPath)
	if os.IsNotExist(err) {
		fmt.Println("No such app: " + app)
		os.Exit(1)
	}

	if info.IsDir() {
		fmt.Println("App is directory: " + app)
		os.Exit(1)
	}

	if info.Mode()&0111 == 0 {
		fmt.Println("App script is not executable: " + appPath)
		os.Exit(1)
	}

	err = runCommand(appPath, args...)
	if err != nil {
		fmt.Println("App exited with error: " + err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
