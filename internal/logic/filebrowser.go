package logic

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gitlab.wige.one/wigeon/sage/internal/models"
)

type FileBrowser struct {
	currentDirectory string
	historyStack     *models.HistoryStack
}

func FileBrowserNew() *FileBrowser {
	fb := FileBrowser{}
	fb.currentDirectory = ""
	fb.historyStack = models.HistoryStackNew()

	return &fb
}

func (fb *FileBrowser) ChangeDirectory(path string) error {

	if path == fb.currentDirectory {
		return nil
	}

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

	fb.historyStack.Push(fb.currentDirectory)
	// FIX: strip trailing '/' or '/.' from path
	fb.currentDirectory = path

	return nil

}

func (fb *FileBrowser) CurrentDirectory() string {
	return fb.currentDirectory
}

func (fb *FileBrowser) NavigateUp() string {
	parentDirectory := filepath.Dir(fb.currentDirectory)

	if parentDirectory != fb.currentDirectory {
		fb.historyStack.Push(fb.currentDirectory)
	}

	fb.currentDirectory = parentDirectory
	return fb.currentDirectory
}

// TODO: implement a history stack to store previous locations
func (fb *FileBrowser) NavigateBack() string {
	fb.currentDirectory = fb.historyStack.Pop()
	return fb.currentDirectory
}
