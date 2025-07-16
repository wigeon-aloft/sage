package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"gitlab.wige.one/wigeon/sage/internal/ui"
)

const APP_ID = "one.wige.gitlab.sage"
const APP_NAME = "Sage"
const APP_VERSION = "0.0.1"

const DEFAULT_LAYOUT_PATH = "internal/ui/layout/sage.xml"

func main() {

	application, err := gtk.ApplicationNew(APP_ID, glib.APPLICATION_FLAGS_NONE)

	if err != nil {
		log.Fatal("Could not create application:", err)
	}

	application.Connect("activate", func() {

		appWindow, err := gtk.ApplicationWindowNew(application)
		if err != nil {
			log.Fatal("Could not create application window: ", err)
		}

		ui.BuildDefaultLayout(appWindow)

		appWindow.SetTitle(fmt.Sprintf("%s - %s", APP_NAME, APP_VERSION))
		appWindow.SetDefaultSize(400, 400)
		appWindow.Show()

		appWindow.Connect("destroy", func() {
			gtk.MainQuit()
		})

	})

	application.Run(os.Args)
}
