package assets

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// UIConfig holds configurations for the UI
type UIConfig struct {
	Theme  string   // Theme ("light", "dark", or custom)
	Layout string   // Layout style ("grid", "list")
	Apps   []string // List of apps
}

// DefaultConfig returns default configurations
func DefaultConfig() UIConfig {
	return UIConfig{
		Theme:  "light",
		Layout: "grid",
		Apps:   []string{"App1", "App2", "App3", "App4"},
	}
}

// InitializeUI starts the GUI with the specified configuration
func InitializeUI(config UIConfig) {
	application := app.New()
	window := application.NewWindow("Executo Launcher")

	// Apply theme
	if config.Theme == "dark" {
		application.Settings().SetTheme(theme.DarkTheme())
	} else {
		application.Settings().SetTheme(theme.LightTheme())
	}

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

// buildGridLayout creates a grid layout
func buildGridLayout(config UIConfig) fyne.CanvasObject {
	buttons := make([]fyne.CanvasObject, len(config.Apps))
	for i, appName := range config.Apps {
		buttons[i] = widget.NewButtonWithIcon(appName, theme.DocumentIcon(), func(appName string) func() {
			return func() {
				fmt.Printf("Launching %s\n", appName)
			}
		}(appName))
	}
	return container.NewGridWithColumns(3, buttons...)
}

func buildListLayout(config UIConfig) fyne.CanvasObject {
	search := widget.NewEntry()
	search.SetPlaceHolder("Search apps...")

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

	search.OnChanged = func(input string) {
		filtered := filterApps(config.Apps, input)
		list.Length = func() int { return len(filtered) }
		list.UpdateItem = func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(filtered[id])
		}
		list.Refresh()
	}

	return container.NewVBox(search, list)
}

func filterApps(apps []string, input string) []string {
	var result []string
	for _, app := range apps {
		if strings.Contains(strings.ToLower(app), strings.ToLower(input)) {
			result = append(result, app)
		}
	}
	return result
}
