package tests

import (
	"os/exec"
	"testing"

	"github.com/souvik03-136/Executo/internal/launcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockExecutor struct with corrected Command method signature
type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) Command(name string, arg ...string) *exec.Cmd {
	args := m.Called(name, arg)
	return args.Get(0).(*exec.Cmd)
}

func TestLaunchApp_Success(t *testing.T) {
	// Create a mock executor
	mockExecutor := new(MockExecutor)

	// Define the mock behavior using `cmd` for Windows
	mockExecutor.On("Command", "myapp", mock.Anything).Return(exec.Command("cmd", "/C", "echo", "hello"))

	// Replace the real executor with the mock in the launcher
	launcher.SetExecutor(mockExecutor)

	// Call the function
	err := launcher.LaunchApp("myapp")

	// Assert that no error occurred
	assert.NoError(t, err)
	mockExecutor.AssertExpectations(t)

	// Reset executor to default
	launcher.SetExecutor(&launcher.DefaultExecutor{})
}

func TestLaunchApp_Failure(t *testing.T) {
	// Create a mock executor
	mockExecutor := new(MockExecutor)

	// Define mock behavior for failure
	mockExecutor.On("Command", "nonexistentapp", mock.Anything).Return(exec.Command("nonexistentapp"))
	launcher.SetExecutor(mockExecutor)

	// Call the function
	err := launcher.LaunchApp("nonexistentapp")

	// Assert that an error occurred and contains the correct message
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start application")

	// Reset executor to default
	launcher.SetExecutor(&launcher.DefaultExecutor{})
}
func TestLaunchShortcut_Success(t *testing.T) {
	// Define a shortcut map
	shortcuts := map[string]string{
		"myapp": "path/to/myapp",
	}

	// Create a mock executor
	mockExecutor := new(MockExecutor)

	// Define the mock behavior
	mockExecutor.On("Command", "path/to/myapp", mock.Anything).Return(exec.Command("cmd", "/C", "echo", "hello"))
	launcher.SetExecutor(mockExecutor)

	// Call the function
	err := launcher.LaunchShortcut("myapp", shortcuts)

	// Assert that no error occurred
	assert.NoError(t, err)
	mockExecutor.AssertExpectations(t)

	// Reset executor to default
	launcher.SetExecutor(&launcher.DefaultExecutor{})
}

func TestLaunchShortcut_Failure(t *testing.T) {
	// Define an empty shortcut map
	shortcuts := map[string]string{}

	// Call the function
	err := launcher.LaunchShortcut("nonexistent", shortcuts)

	// Assert that an error occurred and contains the correct message
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shortcut 'nonexistent' not found")
}
