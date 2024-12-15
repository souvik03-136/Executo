package launcher

import (
	"errors"
)

// AddShortcut adds a new shortcut to the shortcuts map.
func AddShortcut(alias, appPath string, shortcuts map[string]string) error {
	if _, exists := shortcuts[alias]; exists {
		return errors.New("shortcut already exists")
	}
	shortcuts[alias] = appPath
	return nil
}

// RemoveShortcut removes a shortcut from the shortcuts map.
func RemoveShortcut(alias string, shortcuts map[string]string) error {
	if _, exists := shortcuts[alias]; !exists {
		return errors.New("shortcut does not exist")
	}
	delete(shortcuts, alias)
	return nil
}

// ListShortcuts returns all shortcuts in the shortcuts map.
func ListShortcuts(shortcuts map[string]string) map[string]string {
	return shortcuts
}
