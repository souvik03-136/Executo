package tests

import (
	"os"
	"testing"

	"github.com/souvik03-136/Executo/internal/appscanner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	// Create a temporary directory and files to simulate an application directory
	tempDir := t.TempDir()
	tempFile := tempDir + "/testapp"
	err := os.WriteFile(tempFile, []byte{}, 0755) // Make it executable
	require.NoError(t, err, "Failed to create temporary file")

	// Override getAppDirectories for the test
	originalGetAppDirectories := appscanner.GetAppDirectories
	appscanner.GetAppDirectories = func() []string {
		return []string{tempDir}
	}
	defer func() { appscanner.GetAppDirectories = originalGetAppDirectories }()

	// Call Scan and check results
	apps, err := appscanner.Scan()
	require.NoError(t, err, "Scan failed")

	// Use assert to check the number of applications found
	assert.Equal(t, 1, len(apps), "Expected 1 application, but found a different number")

	// Use assert to check the application data
	assert.Equal(t, "testapp", apps[0].Name, "Application name does not match")
	assert.Equal(t, tempFile, apps[0].Path, "Application path does not match")
}

func TestScan_NoApps(t *testing.T) {
	// Create an empty temporary directory
	tempDir := t.TempDir()

	// Override getAppDirectories for the test
	originalGetAppDirectories := appscanner.GetAppDirectories
	appscanner.GetAppDirectories = func() []string {
		return []string{tempDir}
	}
	defer func() { appscanner.GetAppDirectories = originalGetAppDirectories }()

	// Call Scan and check results
	_, err := appscanner.Scan()

	// Assert that the error is expected
	assert.Error(t, err, "Expected an error when no applications are found")
	assert.EqualError(t, err, "no applications found", "Unexpected error message")
}
