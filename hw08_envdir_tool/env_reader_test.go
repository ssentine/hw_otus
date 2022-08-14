package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("wrong directory", func(t *testing.T) {
		env, err := ReadDir("wrong/path")
		require.Empty(t, env)
		require.Equal(t, err, ErrFailToGetStat)
	})

	t.Run("empty directory path", func(t *testing.T) {
		env, err := ReadDir("")
		require.Empty(t, env)
		require.Equal(t, err, ErrFailToGetStat)
	})

	t.Run("path is not a directory", func(t *testing.T) {
		env, err := ReadDir("testdata/echo.sh")
		require.Nil(t, env)
		require.ErrorIs(t, err, ErrIsNotDirectory)
	})

	t.Run("directory exists and has keys", func(t *testing.T) {
		expected := Environment{
			"BAR":   EnvValue{Value: "bar", NeedRemove: false},
			"HELLO": EnvValue{Value: `"hello"`, NeedRemove: false},
			"FOO":   EnvValue{Value: "   foo\nwith new line", NeedRemove: false},
			"EMPTY": EnvValue{Value: "", NeedRemove: false},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
		}
		env, err := ReadDir("testdata/env")
		require.Nil(t, err)
		for k, v := range expected {
			require.Equal(t, v, env[k], "not valid env value key = %v, %v != %v", k, v, env[k])
		}
	})

	t.Run("env name with =", func(t *testing.T) {
		dir, err := os.MkdirTemp("/tmp", "env_test")
		require.Nil(t, err)

		badFile, err := os.CreateTemp(dir, "WORLD=")
		require.Nil(t, err)

		envs, err := ReadDir(dir)
		require.Nil(t, err)

		os.Remove(badFile.Name())
		err = os.Remove(dir)
		require.Nil(t, err)

		require.Empty(t, envs)
	})
}
