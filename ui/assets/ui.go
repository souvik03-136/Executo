package assets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// UIConfig holds configurations for the UI
type UIConfig struct {
	Theme  string   // Current theme ("light", "dark", or user-defined)
	Layout string   // Layout style (e.g., "grid", "list")
	Apps   []string // List of apps to display in the UI
}

// DefaultConfig provides default UI configurations
func DefaultConfig() UIConfig {
	return UIConfig{
		Theme:  "light",
		Layout: "grid",
		Apps:   []string{"App1", "App2", "App3", "App4"},
	}
}

// InitializeUI starts the user interface with the provided configuration
func InitializeUI(config UIConfig) {
	application := app.New()
	window := application.NewWindow("Executo Launcher")

	// Apply theme
	if config.Theme == "dark" {
		application.Settings().SetTheme(theme.DarkTheme())
	} else {
		application.Settings().SetTheme(theme.LightTheme())
	}

	// Build UI components
	var content fyne.CanvasObject
	if config.Layout == "grid" {
		content = buildGridLayout(config)
	} else {
		content = buildListLayout(config)
	}

	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}

// buildGridLayout creates a grid layout for the apps
func buildGridLayout(config UIConfig) fyne.CanvasObject {
	buttons := make([]fyne.CanvasObject, len(config.Apps))
	for i, appName := range config.Apps {
		buttons[i] = widget.NewButton(appName, func(appName string) func() {
			return func() {
				fmt.Printf("Launching %s\n", appName)
			}
		}(appName))
	}
	return container.NewGridWithColumns(3, buttons...)
}

// buildListLayout creates a list layout for the apps
func buildListLayout(config UIConfig) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return len(config.Apps)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("App Name")
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(config.Apps[i])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		fmt.Printf("Launching %s\n", config.Apps[id])
	}

	return container.NewVBox(list)
}

// Example usage
func Example() {
	config := DefaultConfig()
	config.Theme = "dark"  // Set theme to dark
	config.Layout = "grid" // Choose grid layout
	InitializeUI(config)
}
