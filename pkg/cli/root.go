package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "executo",
	Short: "Executo - A smart application launcher",
	Long: `Executo is a tool that scans for installed applications and 
lets you launch them quickly and easily.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executo CLI - Use 'executo --help' to see available commands.")
	},
}
