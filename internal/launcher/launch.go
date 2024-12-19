package launcher

import (
	"errors"
	"os/exec"
)

type Executor interface {
	Command(name string, arg ...string) *exec.Cmd
}

type DefaultExecutor struct{}

func (e *DefaultExecutor) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

var executor Executor = &DefaultExecutor{}

func SetExecutor(e Executor) {
	executor = e
}

func LaunchApp(appName string) error {
	cmd := executor.Command(appName)
	err := cmd.Run()
	if err != nil {
		return errors.New("failed to start application: " + err.Error())
	}
	return nil
}

func LaunchShortcut(appName string, shortcuts map[string]string) error {
	path, exists := shortcuts[appName]
	if !exists {
		return errors.New("shortcut '" + appName + "' not found")
	}

	return LaunchApp(path)
}
