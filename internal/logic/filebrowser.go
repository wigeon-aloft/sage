package logic

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"gitlab.wige.one/wigeon/sage/internal/models"
)

var (
	NoMappingError = errors.New("No application mapping for given filetype")
)

type FileBrowser struct {
	currentDirectory string
	historyStack     *models.HistoryStack
	settings         *Settings
}

func FileBrowserNew() *FileBrowser {
	fb := FileBrowser{}
	fb.currentDirectory = ""
	fb.historyStack = models.HistoryStackNew()

	return &fb
}

func FileBrowserNewWithSettings(settings *Settings) *FileBrowser {
	fb := FileBrowserNew()
	fb.settings = settings

	return fb
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

func (fb *FileBrowser) NavigateBack() string {
	fb.currentDirectory = fb.historyStack.Pop()
	return fb.currentDirectory
}

func (fb *FileBrowser) GetCurrentDirContents() ([]os.FileInfo, error) {
	dirEntries, err := os.ReadDir(fb.currentDirectory)
	if err != nil {
		return nil, err
	}

	dirContents := make([]os.FileInfo, len(dirEntries))
	for idx, dirEntry := range dirEntries {
		info, err := dirEntry.Info()

		if err != nil {
			return nil, err
		}

		dirContents[idx] = info
	}

	return dirContents, nil
}

func SaveApplicationMapping(executablePath string, extension string) {
	// TODO: Save application-filetype mapping to users home directory
}

func (fb *FileBrowser) OpenFileExternallyWithMapping(filePath string) error {

	application := fb.settings.LookupApplication(path.Ext(filePath))

	if application == "" {
		return NoMappingError
	}

	err := fb.OpenFileExternally(application, filePath, false)
	if err != nil {
		return err
	}

	return nil
}

func (fb *FileBrowser) OpenFileExternally(executablePath string, filePath string, save bool) error {

	executablePathSplit := strings.Split(executablePath, " ")
	executable := executablePathSplit[0]
	args := append(executablePathSplit[1:], filePath)

	command := exec.Command(executable, args...)
	err := command.Start()

	if err != nil {
		return err
	}

	if save {
		fb.settings.AddApplicationFiletypeMapping(
			path.Ext(filePath),
			executablePath,
		)
	}

	return nil
}
