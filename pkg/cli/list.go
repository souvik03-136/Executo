package cli

import (
	"fmt"

	"github.com/souvik03-136/Executo/internal/appscanner"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available applications",
	Long:  "Scan the system and list all installed applications.",
	Run: func(cmd *cobra.Command, args []string) {
		apps, err := appscanner.Scan()
		if err != nil {
			fmt.Printf("Error scanning applications: %v\n", err)
			return
		}

		fmt.Println("Available Applications:")
		for _, app := range apps {
			fmt.Printf("- %s (%s)\n", app.Name, app.Path)
		}
	},
}
