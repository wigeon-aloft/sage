package logic

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FileBrowser struct {
	currentDirectory string
}

func (fb *FileBrowser) ChangeDirectory(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	if fileInfo.Mode().IsRegular() {
		return errors.New(fmt.Sprint(path, " is a file, not a directory"))
	}

	fb.currentDirectory = path

	return nil

}

func (fb *FileBrowser) CurrentDirectory() string {
	return fb.currentDirectory
}

func (fb *FileBrowser) NavigateUp() string {
	fb.currentDirectory = filepath.Dir(fb.currentDirectory)
	return fb.currentDirectory
}

// TODO: implement a history stack to store previous locations
func (fb *FileBrowser) NavigateBack() string {
	return ""
}
