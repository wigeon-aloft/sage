package ui

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"gitlab.wige.one/wigeon/sage/internal/logic"
)

type Application struct {
	id      string
	name    string
	version string

	*gtk.Application
	*gtk.ApplicationWindow

	settings *logic.Settings

	fileBrowserUI *FileBrowserUI
}

func ApplicationNew(id, name, version string) (*Application, error) {

	var applicationWindow *gtk.ApplicationWindow

	application := Application{}

	application.id = id
	application.name = name
	application.version = version

	gtkApplication, err := gtk.ApplicationNew(
		id,
		glib.APPLICATION_FLAGS_NONE,
	)
	if err != nil {
		log.Fatal("Unable to create GTK application:", err)
	}
	application.Application = gtkApplication

	settings, err := logic.SettingsNew()
	if err != nil {
		log.Fatal(err)
	}
	err = settings.ReadApplicationFiletypeMapping()
	if err != nil {
		log.Fatal(err)
	}
	application.settings = settings

	application.Connect("activate", func() {

		applicationWindow, err = gtk.ApplicationWindowNew(gtkApplication)
		if err != nil {
			log.Fatal("Unable to create GTK application window:", err)
		}
		applicationWindow.SetTitle(fmt.Sprintf("%s - %s", name, version))
		applicationWindow.SetDefaultSize(800, 600)
		applicationWindow.SetSizeRequest(800, 600)
		applicationWindow.Connect("destroy", application.GTKDestroyHandler)

		fileBrowserUI, err := FileBrowserUINew(
			applicationWindow,
			settings,
		)
		if err != nil {
			log.Fatal(err)
		}
		application.fileBrowserUI = fileBrowserUI

		applicationWindow.Add(application.fileBrowserUI.Layout)
		application.ApplicationWindow = applicationWindow

		application.Show()
		application.fileBrowserUI.Layout.ShowAll()
	})

	return &application, nil

}

func (a *Application) shutdownSignalListenerStart() error {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		oscall := <-c
		fmt.Println("Received signal", oscall)
		a.Exit(0)
	}()

	return nil
}

func (a *Application) Start() (int, error) {

	err := a.shutdownSignalListenerStart()
	if err != nil {
		// TODO: define error codes, below '1' is just a placeholder
		return 1, err
	}

	// We call a.Run() last, as it blocks until GTK exits
	exitCode := a.Run(os.Args)

	return exitCode, nil
}

func (a *Application) GTKDestroyHandler(applicationWindow *gtk.ApplicationWindow) {
	// TODO: pass a more relevant exit code once exit codes are implemented
	a.Exit(0)
}

func (a *Application) Exit(exitCode int) {

	err := a.settings.WriteApplicationFiletypeMapping()
	if err != nil {
		log.Println(err)
	}

	gtk.MainQuit()
	os.Exit(exitCode)
}
