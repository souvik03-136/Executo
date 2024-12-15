package appscanner

import (
	"errors"
	"os"
	"path/filepath"
)

// Default function to get app directories.
var GetAppDirectories = func() []string {
	if os.Getenv("OS") == "Windows_NT" {
		return []string{os.Getenv("PROGRAMFILES"), os.Getenv("PROGRAMFILES(X86)"), os.Getenv("LOCALAPPDATA")}
	}
	return []string{"/usr/bin", "/usr/local/bin", "/Applications"}
}

// Scan scans the system for available applications and returns a list.
func Scan() ([]Application, error) {
	var applications []Application

	directories := GetAppDirectories()

	for _, dir := range directories {
		apps, err := scanDirectory(dir)
		if err != nil {
			return nil, err
		}
		applications = append(applications, apps...)
	}

	if len(applications) == 0 {
		return nil, errors.New("no applications found")
	}

	return applications, nil
}

// scanDirectory and isExecutable functions remain the same.
func scanDirectory(directory string) ([]Application, error) {
	var apps []Application

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isExecutable(info) {
			apps = append(apps, Application{Name: info.Name(), Path: path})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return apps, nil
}

func isExecutable(info os.FileInfo) bool {
	mode := info.Mode()
	return mode&0111 != 0 // Check if any execute bits are set
}
