package dialogs

import (
	"github.com/gotk3/gotk3/gtk"
)

type OpenFileDialog struct {
	*gtk.Dialog
	binaryPathEntry          *gtk.Entry
	saveSelectionCheckButton *gtk.CheckButton
}

func OpenFileDialogNew(parent gtk.IWindow) (*OpenFileDialog, error) {

	var fileDialog OpenFileDialog
	var dialog *gtk.Dialog

	flags := gtk.DialogFlags(gtk.DIALOG_DESTROY_WITH_PARENT)

	dialog, err := gtk.DialogNewWithButtons(
		"Open file...",
		parent,
		flags,
		[]any{"Close", gtk.RESPONSE_CLOSE},
		[]any{"Open", gtk.RESPONSE_NONE},
	)
	if err != nil {
		return nil, err
	}
	dialog.SetModal(true)

	promptLabel, err := gtk.LabelNew("Specify a path to an application to open this file:")
	if err != nil {
		return nil, err
	}

	binaryPathEntry, err := gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	saveSelectionCheckButton, err := gtk.CheckButtonNew()
	if err != nil {
		return nil, err
	}

	saveSelectionLabel, err := gtk.LabelNew("Save selected path for this file type?")
	if err != nil {
		return nil, err
	}

	checkBoxLayout, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 8)
	if err != nil {
		return nil, err
	}
	checkBoxLayout.Add(saveSelectionCheckButton)
	checkBoxLayout.Add(saveSelectionLabel)

	contentArea, err := dialog.GetContentArea()
	if err != nil {
		return nil, err
	}

	contentArea.Add(promptLabel)
	contentArea.Add(binaryPathEntry)
	contentArea.Add(checkBoxLayout)

	fileDialog.Dialog = dialog
	fileDialog.binaryPathEntry = binaryPathEntry
	fileDialog.saveSelectionCheckButton = saveSelectionCheckButton

	fileDialog.Connect("response", fileDialog.dialogResponse)

	return &fileDialog, nil
}

func (ofd *OpenFileDialog) dialogResponse(dialog *gtk.Dialog) {

	// TODO: write dialog response
	ofd.Destroy()

}
