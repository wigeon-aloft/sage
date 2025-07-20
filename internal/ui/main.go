package ui

import (
	"log"

	"gitlab.wige.one/wigeon/sage/internal/ui/dialogs"

	"github.com/gotk3/gotk3/gtk"
)

func BuildDefaultLayout(appWindow *gtk.ApplicationWindow) {
	vbox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	if err != nil {
		log.Fatal("Could not create VBox: ", err)
	}

	treeView, err := setupFileTreeView()
	if err != nil {
		log.Fatal(err)
	}

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
			errDialog, err := dialogs.ErrorDialogNew(
				"Error",
				appWindow,
				err.Error(),
			)
			if err != nil {
				log.Fatal("Unable to create error dialog: ", err)
			}

			errDialog.ShowAll()
		}
	})

	scrollableWindow.Add(treeView)

	upButton, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal()
	}
	upButton.SetLabel("UP")

	upButton.Connect("clicked", func(button *gtk.Button) {

		treeViewModel, err := treeView.GetModel()
		if err != nil {
			log.Fatal("Could not get treeView model: ", err)
		}

		fileBrowser.NavigateUp()

		err = updateFileTreeView(treeViewModel.(*gtk.ListStore), fileBrowser.CurrentDirectory())

		if err != nil {
			log.Fatal("Unable to update file treeview: ", err)
		}

		entry.SetText(fileBrowser.CurrentDirectory())
	})

	vbox.Add(entry)
	vbox.Add(upButton)
	vbox.Add(scrollableWindow)

	appWindow.Add(vbox)

	appWindow.ShowAll()

	// TODO: implement gtk Builder to build layout from file
}
