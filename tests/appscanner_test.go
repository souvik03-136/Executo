package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/souvik03-136/Executo/internal/appscanner"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScan(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "testapp.exe")
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	require.NoError(t, err, "Failed to create test application file")

	originalGetAppDirectories := appscanner.GetAppDirectories
	appscanner.GetAppDirectories = func() []string {
		return []string{tempDir}
	}
	defer func() { appscanner.GetAppDirectories = originalGetAppDirectories }()

	apps, err := appscanner.Scan()
	if err != nil {
		t.Logf("Error during Scan: %s", err)
	}
	require.NoError(t, err, "Scan failed")
	assert.Equal(t, 1, len(apps), "Expected one application")
	assert.Equal(t, "testapp.exe", filepath.Base(apps[0].Path), "Application path mismatch")
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

func TestExample(t *testing.T) {
	t.Log("Test ran successfully!")
}
