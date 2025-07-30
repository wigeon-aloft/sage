package ui

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"gitlab.wige.one/wigeon/sage/internal/logic"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	COLUMN_FILENAME int = iota
	COLUMN_SIZE
	COLUMN_MODIFIED
	COLUMN_EXTENSION
)

type Column struct {
	Index int
	Name  string
	Type  glib.Type
}

var DEFAULT_COLUMNS = []Column{
	Column{Index: COLUMN_FILENAME, Name: "Name", Type: glib.TYPE_STRING},
	Column{Index: COLUMN_SIZE, Name: "Size", Type: glib.TYPE_STRING},
	Column{Index: COLUMN_MODIFIED, Name: "Modified", Type: glib.TYPE_STRING},
	Column{Index: COLUMN_EXTENSION, Name: "Type", Type: glib.TYPE_STRING},
}

type FileBrowserUI struct {
	fileBrowser *logic.FileBrowser

	fileListStore *gtk.ListStore
	fileTreeView  *gtk.TreeView
	Layout        *gtk.Box
	pathEntry     *gtk.Entry
}

func FileBrowserUINew() (*FileBrowserUI, error) {
	var fbui FileBrowserUI

	fbui.fileBrowser = logic.FileBrowserNew()

	treeView, listStore, err := setupFileTreeView()
	if err != nil {
		return nil, err
	}
	treeView.Connect("row-activated", fbui.treeViewRowActivatedConnection)
	treeView.SetModel(listStore)
	fbui.fileTreeView = treeView
	fbui.fileListStore = listStore

	layout, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	if err != nil {
		return nil, err
	}

	scrollableWindow, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}
	scrollableWindow.SetPropagateNaturalHeight(true)

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Could not create entry widget:", err)
	}
	entry.Connect("activate", fbui.pathEntryActivatedConnection)
	entry.SetText(fbui.fileBrowser.CurrentDirectory())
	fbui.pathEntry = entry

	scrollableWindow.Add(treeView)

	upButton, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal()
	}
	upButton.SetLabel("↑")
	upButton.Connect("clicked", fbui.upButtonClickedConnection)

	backButton, err := gtk.ButtonNew()
	if err != nil {
		log.Fatal()
	}
	backButton.SetLabel("←")
	backButton.Connect("clicked", fbui.backButtonClickedConnection)

	buttonBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 8)

	buttonBox.Add(backButton)
	buttonBox.Add(upButton)

	layout.Add(entry)
	layout.Add(buttonBox)
	layout.Add(scrollableWindow)

	fbui.Layout = layout

	return &fbui, nil

}

func createFileTreeViewColumn(column Column) (*gtk.TreeViewColumn, error) {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, err
	}

	columnGtk, err := gtk.TreeViewColumnNewWithAttribute(column.Name, cellRenderer, "text", column.Index)
	if err != nil {
		return nil, err
	}

	columnGtk.SetClickable(true)
	columnGtk.SetReorderable(true)
	columnGtk.SetResizable(true)
	columnGtk.SetMinWidth(16)

	return columnGtk, nil
}

// TODO: use map[int(column index)]string for file info (name, size, etc.)
func (fbui *FileBrowserUI) addRow(name, size, modified, extension string) error {
	iter := fbui.fileListStore.Append()

	err := fbui.fileListStore.Set(
		iter,
		[]int{COLUMN_FILENAME, COLUMN_SIZE, COLUMN_MODIFIED, COLUMN_EXTENSION},
		[]any{name, size, modified, extension},
	)

	if err != nil {
		return err
	}

	return nil
}

func setupFileTreeView() (*gtk.TreeView, *gtk.ListStore, error) {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		return nil, nil, err
	}

	var listStoreTypes []glib.Type
	for _, defaultColumn := range DEFAULT_COLUMNS {
		column, err := createFileTreeViewColumn(defaultColumn)
		if err != nil {
			return nil, nil, err
		}
		treeView.AppendColumn(column)
		listStoreTypes = append(listStoreTypes, defaultColumn.Type)
	}

	listStore, err := gtk.ListStoreNew(listStoreTypes...)

	if err != nil {
		return nil, nil, err
	}
	treeView.SetModel(listStore)

	return treeView, listStore, nil
}

func (fbui *FileBrowserUI) updateFileTreeView() error {

	contents, err := fbui.fileBrowser.GetCurrentDirContents()

	if err != nil {
		return err
	}

	fbui.fileListStore.Clear()

	var ext string

	for _, item := range contents {

		ext = filepath.Ext(item.Name())
		if ext == "" {
			ext = "dir"
		}

		fbui.addRow(
			item.Name(),
			fmt.Sprint(item.Size()),
			item.ModTime().Format(time.RFC822),
			ext,
		)
	}

	return nil
}

func (fbui *FileBrowserUI) treeViewRowActivatedConnection(tv *gtk.TreeView, tp *gtk.TreePath, tvc *gtk.TreeViewColumn) error {

	iter, err := fbui.fileListStore.GetIter(tp)
	if err != nil {
		return err
	}

	columnValue, err := fbui.fileListStore.GetValue(iter, COLUMN_FILENAME)
	if err != nil {
		return err
	}

	path, err := columnValue.GetString()
	if err != nil {
		return err
	}

	err = fbui.fileBrowser.ChangeDirectory(
		filepath.Join(fbui.fileBrowser.CurrentDirectory(), path),
	)
	if err != nil {
		return err
	}

	err = fbui.updateFileTreeView()
	if err != nil {
		return err
	}

	fbui.pathEntry.SetText(fbui.fileBrowser.CurrentDirectory())

	return nil
}

func (fbui *FileBrowserUI) backButtonClickedConnection(_ *gtk.Button) {

	previousDirectory := fbui.fileBrowser.NavigateBack()
	fbui.pathEntry.SetText(previousDirectory)

	err := fbui.updateFileTreeView()

	if err != nil {
		log.Fatal("Unable to update file treeview: ", err)
	}

}

func (fbui *FileBrowserUI) upButtonClickedConnection(_ *gtk.Button) {

	newPath := fbui.fileBrowser.NavigateUp()
	fbui.pathEntry.SetText(newPath)

	err := fbui.updateFileTreeView()

	if err != nil {
		log.Fatal("Unable to update file treeview: ", err)
	}

}

func (fbui *FileBrowserUI) pathEntryActivatedConnection(entry *gtk.Entry) {
	query, err := entry.GetText()
	if err != nil {
		log.Fatal("Unable to get text from Entry widget: ", err)
	}

	err = fbui.fileBrowser.ChangeDirectory(query)
	if err != nil {
		log.Fatal("Unable to change directory: ", err)
	}

	err = fbui.updateFileTreeView()
	if err != nil {
		log.Fatal("Unable to update file treeview: ", err)
	}
}
