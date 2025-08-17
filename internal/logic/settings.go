package logic

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Settings struct {
	UserSettingsPath string

	applicationFiletypeMapping ApplicationFiletypeMapping
}

type ApplicationFiletypeMapping map[string]string

const (
	SETTINGS_DIRECTORY_NAME           = ".wagtail"
	APPLICATION_FILETYPE_MAPPING_FILE = "mappings"
)

func createSettingsDirectory() error {
	userSettingsPath := filepath.Join(os.Getenv("HOME"), SETTINGS_DIRECTORY_NAME)

	// Exit if the directory already exists
	if _, err := os.Stat(userSettingsPath); err == nil {
		return nil
	}

	err := os.Mkdir(userSettingsPath, 0755)

	if err != nil {
		return err
	}

	return nil
}

func SettingsNew() (*Settings, error) {

	var userSettingsPath = filepath.Join(os.Getenv("HOME"), SETTINGS_DIRECTORY_NAME)
	var applicationFiletypeMapping map[string]string

	err := createSettingsDirectory()
	if err != nil {
		return nil, err
	}

	applicationFiletypeMapping = make(map[string]string)

	return &Settings{
		UserSettingsPath:           userSettingsPath,
		applicationFiletypeMapping: applicationFiletypeMapping,
	}, nil

}

// TODO: Implement loading default settings from file
func SettingsDefaultNew() (*Settings, error) {
	s, err := SettingsNew()
	return s, err
}

func (s *Settings) AddApplicationFiletypeMapping(applicationPath string, fileType string) {
	s.applicationFiletypeMapping[applicationPath] = fileType
}

func (s *Settings) RemoveApplicationFiletypeMapping(applicationPathKey string) {
	delete(s.applicationFiletypeMapping, applicationPathKey)
}

func (s *Settings) GetApplicationFiletypeMapping() ApplicationFiletypeMapping {
	return s.applicationFiletypeMapping
}

func (s *Settings) WriteApplicationFiletypeMapping() error {

	f, err := os.OpenFile(
		path.Join(s.UserSettingsPath, APPLICATION_FILETYPE_MAPPING_FILE),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	for applicationPath, extension := range s.applicationFiletypeMapping {
		_, err := writer.WriteString(strings.Join([]string{applicationPath, extension}, ",") + "\n")

		if err != nil {
			return err
		}

	}

	return nil
}

func (s *Settings) ReadApplicationFiletypeMapping() error {

	f, err := os.OpenFile(
		path.Join(s.UserSettingsPath, APPLICATION_FILETYPE_MAPPING_FILE),
		os.O_RDONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	applicationFiletypeMapping := make(ApplicationFiletypeMapping)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		lineSplit := strings.Split(line, ",")
		applicationFiletypeMapping[lineSplit[0]] = lineSplit[1]

	}

	s.applicationFiletypeMapping = applicationFiletypeMapping

	return nil

}
