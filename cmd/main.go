package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"gitlab.wige.one/wigeon/sage/internal/ui"
)

const APP_ID = "one.wige.gitlab.sage"
const APP_NAME = "Sage"
const APP_VERSION = "0.0.1"

const DEFAULT_LAYOUT_PATH = "internal/ui/layout/sage.xml"

func main() {
	gtk.Init(nil)

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
	
	// Create your application via your custom UI wrapper
	application, err := ui.ApplicationNew(APP_ID, APP_NAME, APP_VERSION)
	if err != nil {
		log.Fatal(err)
	}

	// Load the layout via GTK Builder
	builder, err := gtk.BuilderNewFromFile(DEFAULT_LAYOUT_PATH)
	if err != nil {
		log.Fatalf("Unable to load UI file: %v", err)
	}

	// Get the main window defined in your sage.xml
	obj, err := builder.GetObject("main_window") // <-- ID must match sage.xml
	if err != nil {
		log.Fatalf("Could not get main_window: %v", err)
	}

	win, ok := obj.(*gtk.Window)
	if !ok {
		log.Fatal("main_window is not a *gtk.Window")
	}

	// Attach window close event
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.ShowAll()

	// Start your application
	exitCode, err := application.Start()
	if err != nil {
		log.Fatalf("Application failed: %v", err)
	}
	application.Exit(exitCode)
}
