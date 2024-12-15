package launcher

import (
	"fmt"
	"os/exec"
)

// LaunchApp launches an application by its path or name.
func LaunchApp(appPath string) error {
	cmd := exec.Command(appPath)
	cmd.Stderr = nil
	cmd.Stdout = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start application: %w", err)
	}

	return nil
}

// LaunchShortcut launches an application using a shortcut.
func LaunchShortcut(alias string, shortcuts map[string]string) error {
	appPath, exists := shortcuts[alias]
	if !exists {
		return fmt.Errorf("shortcut '%s' not found", alias)
	}

	return LaunchApp(appPath)
}
