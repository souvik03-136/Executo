package appscanner

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
	dirs := GetAppDirectories()
	for _, dir := range dirs {
		fmt.Printf("Scanning directory: %s\n", dir)
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory: %s\n", err)
			return nil, fmt.Errorf("failed to read directory: %w", err)
		}
		for _, file := range files {
			fmt.Printf("File found: %s, IsDir: %t, Mode: %s\n", file.Name(), file.IsDir(), file.Mode())
			if strings.HasSuffix(file.Name(), ".exe") {
				fmt.Printf("Application detected: %s\n", file.Name())
				applications = append(applications, Application{
					Name: file.Name(),
					Path: filepath.Join(dir, file.Name()),
				})
			}
		}
	}
	if len(applications) == 0 {
		fmt.Println("No applications found")
		return nil, fmt.Errorf("no applications found")
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
