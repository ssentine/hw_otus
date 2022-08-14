package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrFailToGetStat       = errors.New("failed to get file info")
	ErrIsNotDirectory      = errors.New("it is not a directory")
	ErrFailToReadDirectory = errors.New("failed to read directory")
	ErrFailToOpenFile      = errors.New("failed to open file")
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		return nil, ErrFailToGetStat
	}

	if !dirInfo.IsDir() {
		return nil, ErrIsNotDirectory
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, ErrFailToReadDirectory
	}

	envMap := Environment{}
	for _, file := range files {
		fileName := file.Name()
		if file.IsDir() || !file.Mode().IsRegular() || strings.Contains(fileName, "=") {
			continue
		}

		fd, err := os.Open(path.Join(dir, fileName))
		if err != nil {
			return nil, ErrFailToOpenFile
		}
		defer fd.Close()

		scanner := bufio.NewScanner(fd)
		if scanner.Scan() {
			value := scanner.Text()
			value = strings.ReplaceAll(value, "\000", "\n")
			value = strings.TrimRight(value, " \t\n")
			envMap[fileName] = EnvValue{Value: value, NeedRemove: false}
		} else {
			envMap[fileName] = EnvValue{Value: "", NeedRemove: true}
		}
	}
	return envMap, nil
}
