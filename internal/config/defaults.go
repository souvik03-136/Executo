package config

// GetDefaultConfig returns the default configuration settings.
func GetDefaultConfig() Config {
	return Config{
		ScanDirectories: []string{
			"/usr/bin",
			"/usr/local/bin",
			"/Applications",
		},
		DefaultApp: "",
		Shortcuts: map[string]string{
			"browser": "/usr/bin/firefox",
			"editor":  "/usr/bin/vim",
		},
	}
}
