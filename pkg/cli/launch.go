package cli

import (
	"fmt"

	"github.com/souvik03-136/Executo/internal/launcher"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(launchCmd)
}

var launchCmd = &cobra.Command{
	Use:   "launch [app_name]",
	Short: "Launch an application by name",
	Long:  "Provide the name of an installed application to launch it.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		fmt.Printf("Launching application: %s\n", appName)
		err := launcher.LaunchApp(appName)
		if err != nil {
			fmt.Printf("Error launching application: %v\n", err)
			return
		}
		fmt.Printf("Application '%s' launched successfully.\n", appName)
	},
}
