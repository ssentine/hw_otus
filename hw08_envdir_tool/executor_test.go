package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("ls -l -a", func(t *testing.T) {
		path, err := os.MkdirTemp(".", "temp_dir")
		require.NoError(t, err)
		defer os.RemoveAll(path)

		file, err := os.CreateTemp(path, "test_env")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		_, err = file.Write([]byte("HELLO WORLD!"))
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)

		env, err := ReadDir(path)
		require.NoError(t, err)

		args := []string{"ls", "-l", "-a"}

		code := RunCmd(args, env)

		testEnv, ok := os.LookupEnv(filepath.Base(file.Name()))
		require.True(t, ok)
		require.Equal(t, "HELLO WORLD!", testEnv)
		require.Equal(t, 0, code)
	})
	t.Run("command and args exist", func(t *testing.T) {
		path, err := os.MkdirTemp(".", "temp_dir")
		require.NoError(t, err)
		defer os.RemoveAll(path)

		file, err := os.CreateTemp(path, "text1.txt")
		require.NoError(t, err)
		defer os.Remove(file.Name())

		_, err = file.Write([]byte("some data\n"))
		require.NoError(t, err)
		file.Close()

		cmd := []string{"cat", file.Name()}
		env := Environment{}
		code := RunCmd(cmd, env)
		require.Equal(t, 0, code)
	})
	t.Run("command and args don`t exist", func(t *testing.T) {
		cmd := []string{"wrongCommand", "wrongArg"}
		env := Environment{}
		code := RunCmd(cmd, env)
		require.Equal(t, -1, code)
	})
}
