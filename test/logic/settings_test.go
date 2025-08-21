package logic_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"testing"

	"gitlab.wige.one/wigeon/sage/internal/logic"
)

func TestMain(m *testing.M) {
	s, err := logic.SettingsNew()
	if err != nil {
		log.Fatal("Could not create initial settings struct: ", err)
	}

	os.RemoveAll(s.UserSettingsPath)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestSettingsAddApplicationFiletypeMapping(t *testing.T) {

	targetMap := logic.ApplicationFiletypeMapping{
		"sh":  "/usr/bin/bash",
		"pdf": "atril",
	}
	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("sh", "/usr/bin/bash")
	s.AddApplicationFiletypeMapping("pdf", "atril")

	fmt.Printf("%+v\n", s.GetApplicationFiletypeMapping())

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
		"pdf": "atril",
	}
	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("sh", "/usr/bin/bash")
	s.AddApplicationFiletypeMapping("pdf", "atril")

	s.RemoveApplicationFiletypeMapping("sh")

	fmt.Printf("%+v\n", s.GetApplicationFiletypeMapping())

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

func TestSettingsLookupApplicationFiletypeMapping(t *testing.T) {

	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("sh", "/usr/bin/bash")

	if s.LookupApplication("sh") != "/usr/bin/bash" {
		t.Fatalf(
			"s.LookupApplication(%q) == %q, expected %q",
			"sh",
			s.LookupApplication("sh"),
			"/usr/bin/bash",
		)
	}

}

func TestSettingsWriteApplicationFiletypeMapping(t *testing.T) {

	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	s.AddApplicationFiletypeMapping("sh", "/usr/bin/bash")
	s.AddApplicationFiletypeMapping("pdf", "atril")

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

	expectedContent := "sh,/usr/bin/bash\npdf,atril\n"
	if string(fileContent) != expectedContent {
		t.Fatalf("fileContent == %q, expected %q", string(fileContent), expectedContent)
	}

}

func TestSettingsReadApplicationFiletypeMapping(t *testing.T) {

	targetMap := logic.ApplicationFiletypeMapping{
		"sh":  "/usr/bin/bash",
		"pdf": "atril",
	}

	s, err := logic.SettingsNew()
	if err != nil {
		t.Fatal(err)
	}

	err = s.ReadApplicationFiletypeMapping()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", s.GetApplicationFiletypeMapping())

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
