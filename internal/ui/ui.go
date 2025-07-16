package ui

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func BuildDefaultLayout(appWindow *gtk.ApplicationWindow) {
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	if err != nil {
		log.Fatal("Could not create VBox: ", err)
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Could not create entry widget:", err)
	}

	vbox.Add(entry)
	vbox.Add(setupFileTreeView())

	appWindow.Add(vbox)

	appWindow.ShowAll()

	// TODO - implement gtk Builder to build layout from file
}
