package tests

import (
	"os/exec"
	"testing"

	"github.com/souvik03-136/Executo/internal/launcher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking exec.Command for testing LaunchApp and LaunchShortcut
type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) Command(name string, arg ...string) *exec.Cmd {
	args := m.Called(name, arg)
	return args.Get(0).(*exec.Cmd)
}

// Test the LaunchApp function
func TestLaunchApp_Success(t *testing.T) {
	// Create a mock executor
	mockExecutor := new(MockExecutor)

	// Define the mock behavior
	mockExecutor.On("Command", "myapp").Return(exec.Command("echo", "hello"))

	// Call the function
	err := launcher.LaunchApp("myapp")

	// Assert that no error occurred
	assert.NoError(t, err)
	mockExecutor.AssertExpectations(t)
}

func TestLaunchApp_Failure(t *testing.T) {
	// Create a mock executor
	mockExecutor := new(MockExecutor)

	// Define the mock behavior to simulate a failure
	mockExecutor.On("Command", "nonexistentapp").Return(nil)

	// Call the function
	err := launcher.LaunchApp("nonexistentapp")

	// Assert that an error occurred
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start application")
}

// Test the LaunchShortcut function
func TestLaunchShortcut_Success(t *testing.T) {
	shortcuts := map[string]string{
		"myapp": "path/to/myapp",
	}

	err := launcher.LaunchShortcut("myapp", shortcuts)

	// Assert that no error occurred
	assert.NoError(t, err)
}

func TestLaunchShortcut_Failure(t *testing.T) {
	shortcuts := map[string]string{
		"myapp": "path/to/myapp",
	}

	err := launcher.LaunchShortcut("nonexistent", shortcuts)

	// Assert that an error occurred
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shortcut 'nonexistent' not found")
}

// Test AddShortcut function
func TestAddShortcut_Success(t *testing.T) {
	shortcuts := map[string]string{}
	err := launcher.AddShortcut("myapp", "path/to/myapp", shortcuts)

	// Assert that no error occurred
	assert.NoError(t, err)
	assert.Equal(t, "path/to/myapp", shortcuts["myapp"])
}

func TestAddShortcut_Failure(t *testing.T) {
	shortcuts := map[string]string{
		"myapp": "path/to/myapp",
	}

	err := launcher.AddShortcut("myapp", "new/path/to/myapp", shortcuts)

	// Assert that an error occurred because the shortcut already exists
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shortcut already exists")
}

// Test RemoveShortcut function
func TestRemoveShortcut_Success(t *testing.T) {
	shortcuts := map[string]string{
		"myapp": "path/to/myapp",
	}

	err := launcher.RemoveShortcut("myapp", shortcuts)

	// Assert that no error occurred and the shortcut was removed
	assert.NoError(t, err)
	_, exists := shortcuts["myapp"]
	assert.False(t, exists)
}

func TestRemoveShortcut_Failure(t *testing.T) {
	shortcuts := map[string]string{}

	err := launcher.RemoveShortcut("myapp", shortcuts)

	// Assert that an error occurred because the shortcut does not exist
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shortcut does not exist")
}

// Test ListShortcuts function
func TestListShortcuts(t *testing.T) {
	shortcuts := map[string]string{
		"myapp":    "path/to/myapp",
		"otherapp": "path/to/otherapp",
	}

	result := launcher.ListShortcuts(shortcuts)

	// Assert that the correct shortcuts are returned
	assert.Equal(t, shortcuts, result)
}
