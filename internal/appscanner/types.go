package appscanner

// Application represents the metadata of an installed application.
type Application struct {
	Name string // Name of the application
	Path string // Path to the application executable
}

// Shortcut represents metadata for a shortcut to an application.
type Shortcut struct {
	Alias string // Alias for quick access
	App   string // Name or path of the linked application
}
