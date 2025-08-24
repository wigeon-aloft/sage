package dialogs

import (
	"os"

	"github.com/gotk3/gotk3/gtk"
)

func errorDialogNew(title string, parent gtk.IWindow, errorMessage string) (*gtk.Dialog, error) {

	flags := gtk.DialogFlags(gtk.DIALOG_DESTROY_WITH_PARENT)

	dialog, err := gtk.DialogNewWithButtons(
		title,
		parent,
		flags,
		[]any{"Close", gtk.RESPONSE_CLOSE},
	)

	if err != nil {
		return nil, err
	}

	label, err := gtk.LabelNew(errorMessage)
	if err != nil {
		return nil, err
	}

	contentArea, err := dialog.GetContentArea()
	if err != nil {
		return nil, err
	}

	contentArea.Add(label)

	return dialog, nil

}

func ErrorDialogNew(parent gtk.IWindow, err error) (*gtk.Dialog, error) {
	dialog, err := errorDialogNew("Error", parent, err.Error())

	if err != nil {
		return nil, err
	}

	dialog.Connect("response", func(dialog *gtk.Dialog) {
		dialog.Destroy()
	})

	return dialog, nil
}

// TODO: any err in this function should trigger an immediate call to application.Exit(), no error should be returned
func FatalErrorDialogNew(parent gtk.IWindow, err error) (*gtk.Dialog, error) {

	dialog, err := errorDialogNew("Fatal Error", parent, err.Error())

	if err != nil {
		return nil, err
	}

	dialog.SetModal(true)

	dialog.Connect("response", func(dialog *gtk.Dialog) {
		os.Exit(1)
	})

	return dialog, nil

}
