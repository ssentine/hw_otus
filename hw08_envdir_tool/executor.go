package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	command.Env = createEnv(env)
	if result := command.Run(); result != nil {
		return -1
	}

	return returnCode
}

func createEnv(env Environment) []string {
	for k, v := range env {
		switch v.NeedRemove {
		case true:
			os.Unsetenv(k)
		case false:
			os.Setenv(k, v.Value)
		}
	}
	return os.Environ()
}
