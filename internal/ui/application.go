package ui

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gotk3/gotk3/gtk"
	"gitlab.wige.one/wigeon/sage/internal/logic"
	"gitlab.wige.one/wigeon/sage/internal/ui/dialogs"
)

func SageApplicationStart(appWindow *gtk.ApplicationWindow) {

	settings, err := logic.SettingsNew()
	if err != nil {
		errDialog, err := dialogs.FatalErrorDialogNew(appWindow, err)
		if err != nil {
			log.Fatal(err)
		}
		errDialog.ShowAll()
	}
	err = settings.ReadApplicationFiletypeMapping()
	if err != nil {
		errDialog, err := dialogs.ErrorDialogNew(appWindow, err)
		if err != nil {
			log.Fatal(err)
		}
		errDialog.ShowAll()

	}

	StartShutdownListener(settings)

	fbui, err := FileBrowserUINew(appWindow, settings)
	if err != nil {
		errDialog, err := dialogs.FatalErrorDialogNew(appWindow, err)
		if err != nil {
			log.Fatal(err)
		}
		errDialog.ShowAll()
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
	// FIX: return exitCode, not 0
	os.Exit(0)
}
