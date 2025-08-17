package logic_test

import (
	"io"
	"os"
	"path"
	"reflect"
	"testing"

	"gitlab.wige.one/wigeon/sage/internal/logic"
)

func TestSettingsAddApplicationFiletypeMapping(t *testing.T) {

	targetMap := logic.ApplicationFiletypeMapping{
		"/usr/bin/bash": "sh",
		"atril":         "pdf",
	}
	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("/usr/bin/bash", "sh")
	s.AddApplicationFiletypeMapping("atril", "pdf")

	if !reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap) {
		t.Fatalf(
			"reflect.DeepEqual(%+v, %+v) == %t, expected %t",
			s.GetApplicationFiletypeMapping(),
			targetMap,
			reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap),
			true,
		)
	}
}

func TestSettingsRemoveApplicationFiletypeMapping(t *testing.T) {

	targetMap := logic.ApplicationFiletypeMapping{
		"atril": "pdf",
	}
	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("/usr/bin/bash", "sh")
	s.AddApplicationFiletypeMapping("atril", "pdf")

	s.RemoveApplicationFiletypeMapping("/usr/bin/bash")

	if !reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap) {
		t.Fatalf(
			"reflect.DeepEqual(%+v, %+v) == %t, expected %t",
			s.GetApplicationFiletypeMapping(),
			targetMap,
			reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap),
			true,
		)
	}

}

func TestSettingsWriteApplicationFiletypeMapping(t *testing.T) {

	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("/usr/bin/bash", "sh")
	s.AddApplicationFiletypeMapping("atril", "pdf")

	err = s.WriteApplicationFiletypeMapping()
	if err != nil {
		t.Fatal(err)
	}
	mappingFilePath := path.Join(s.UserSettingsPath, logic.APPLICATION_FILETYPE_MAPPING_FILE)
	f, err := os.Open(mappingFilePath)
	if err != nil {
		t.Fatal(err)
	}

	fileContent, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	expectedContent := "/usr/bin/bash,sh\natril,pdf\n"
	if string(fileContent) != expectedContent {
		t.Fatalf("fileContent == %q, expected %q", string(fileContent), expectedContent)
	}

}

func TestSettingsReadApplicationFiletypeMapping(t *testing.T) {

	targetMap := logic.ApplicationFiletypeMapping{
		"/usr/bin/bash": "sh",
		"atril":         "pdf",
	}

	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	err = s.ReadApplicationFiletypeMapping()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap) {
		t.Fatalf(
			"reflect.DeepEqual(%+v, %+v) == %t, expected %t",
			s.GetApplicationFiletypeMapping(),
			targetMap,
			reflect.DeepEqual(s.GetApplicationFiletypeMapping(), targetMap),
			true,
		)
	}
}
