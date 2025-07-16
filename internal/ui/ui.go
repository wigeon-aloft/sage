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

	treeView := setupFileTreeView()

	scrollableWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("Could not create scrolled window: ", err)
	}
	scrollableWindow.SetPropagateNaturalHeight(true)

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Could not create entry widget:", err)
	}

	entry.Connect("activate", func(entry *gtk.Entry) {
		query, err := entry.GetText()
		if err != nil {
			log.Fatal("Unable to get text from Entry widget: ", err)
		}

		listStore, err := treeView.GetModel()
		if err != nil {
			log.Fatal("Unable to get model from treeview: ", err)
		}

		err = updateFileTreeView(listStore.(*gtk.ListStore), query)
		if err != nil {
			log.Print("Unable to update file treeview: ", err)
		}
	})

	scrollableWindow.Add(treeView)

	vbox.Add(entry)
	vbox.Add(scrollableWindow)

	appWindow.Add(vbox)

	appWindow.ShowAll()

	// TODO - implement gtk Builder to build layout from file
}
