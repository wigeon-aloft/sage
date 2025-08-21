package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"gitlab.wige.one/wigeon/sage/internal/logic"
)

func SageApplicationStart(appWindow *gtk.ApplicationWindow) {

	settings, err := logic.SettingsNew()
	if err != nil {
		log.Fatal("Could not create settings: ", err)
	}
	err = settings.ReadApplicationFiletypeMapping()
	if err != nil {
		log.Fatal("Could not read application-filetype mapping from disk: ", err)
	}

	fbui, err := FileBrowserUINew(appWindow, settings)
	if err != nil {
		log.Fatal("Could not create FileBrowserUINew: ", err)
	}

	appWindow.Add(fbui.Layout)

	appWindow.ShowAll()
}
