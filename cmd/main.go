package main

import (
	"log"

	"gitlab.wige.one/wigeon/sage/internal/ui"
)

const APP_ID = "one.wige.gitlab.sage"
const APP_NAME = "Sage"
const APP_VERSION = "0.0.1"

const DEFAULT_LAYOUT_PATH = "internal/ui/layout/sage.xml"

func main() {

	// application, err := gtk.ApplicationNew(APP_ID, glib.APPLICATION_FLAGS_NONE)
	//
	// if err != nil {
	// 	log.Fatal("Could not create application:", err)
	// }
	//
	// application.Connect("activate", func() {
	//
	// 	appWindow.Connect("destroy", func() {
	// 		gtk.MainQuit()
	// 	})
	//
	// })
	//

	application, err := ui.ApplicationNew(APP_ID, APP_NAME, APP_VERSION)
	if err != nil {
		log.Fatal(err)
	}

	exitCode, err := application.Start()
	application.Exit(exitCode)

}
