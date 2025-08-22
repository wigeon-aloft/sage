package ui

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	StartShutdownListener(settings)

	fbui, err := FileBrowserUINew(appWindow, settings)
	if err != nil {
		log.Fatal("Could not create FileBrowserUINew: ", err)
	}

	appWindow.Add(fbui.Layout)

	appWindow.ShowAll()
}

func StartShutdownListener(settings *logic.Settings) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		oscall := <-c
		fmt.Println("Received signal", oscall)
		Exit(0, settings)
	}()

}

func Exit(exitCode int, settings *logic.Settings) {
	err := settings.WriteApplicationFiletypeMapping()
	if err != nil {
		fmt.Println("Unable to write application-filetype map to file:", err)
	}
	os.Exit(0)
}
