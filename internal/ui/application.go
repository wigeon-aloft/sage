package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func SageApplicationStart(appWindow *gtk.ApplicationWindow) {

	fbui, err := FileBrowserUINew(appWindow)
	if err != nil {
		log.Fatal("Could not create FileBrowserUINew: ", err)
	}

	appWindow.Add(fbui.Layout)

	appWindow.ShowAll()
}
