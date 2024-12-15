package cli

// Execute initializes the CLI and runs the root command
func Execute() error {
	return RootCmd.Execute()
}
