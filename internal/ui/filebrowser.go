package ui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

const (
	COLUMN_FILENAME int = iota
	COLUMN_SIZE
	COLUMN_MODIFIED
	COLUMN_EXTENSION
)

func createFileTreeViewColumn(name string, index int) (*gtk.TreeViewColumn, error) {
	cellRenderer, err := gtk.CellRendererTextNew()
	if err != nil {
		return nil, err
	}

	column, err := gtk.TreeViewColumnNewWithAttribute(name, cellRenderer, "text", int(index))
	if err != nil {
		return nil, err
	}

	return column, nil
}

func addRow(liststore *gtk.ListStore, name, size, modified, extension string) error {
	iter := liststore.Append()

	err := liststore.Set(
		iter,
		[]int{COLUMN_FILENAME, COLUMN_SIZE, COLUMN_MODIFIED, COLUMN_EXTENSION},
		[]any{name, size, modified, extension},
	)

	if err != nil {
		return err
	}

	return nil
}

func setupFileTreeView() *gtk.TreeView {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		log.Fatal("Could not create treeview: ", err)
	}

	column1, err := createFileTreeViewColumn("Name", COLUMN_FILENAME)
	column2, err := createFileTreeViewColumn("Size", COLUMN_SIZE)
	column3, err := createFileTreeViewColumn("Modified", COLUMN_MODIFIED)
	column4, err := createFileTreeViewColumn("Type", COLUMN_EXTENSION)

	if err != nil {
		log.Fatal("Error adding column: ", err)
	}

	treeView.AppendColumn(column1)
	treeView.AppendColumn(column2)
	treeView.AppendColumn(column3)
	treeView.AppendColumn(column4)

	listStore, err := gtk.ListStoreNew(
		glib.TYPE_STRING,
		glib.TYPE_STRING,
		glib.TYPE_STRING,
		glib.TYPE_STRING,
	)

	if err != nil {
		log.Fatal("Could not create liststore: ", err)
	}
	treeView.SetModel(listStore)

	updateFileTreeView(listStore, "/")

	return treeView
}

func updateFileTreeView(liststore *gtk.ListStore, path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	dirEntry, err := os.ReadDir(path)

	if err != nil {
		return err
	}

	liststore.Clear()

	for _, item := range dirEntry {
		itemInfo, err := item.Info()
		if err != nil {
			return err
		}

		addRow(
			liststore,
			item.Name(),
			fmt.Sprint(itemInfo.Size()),
			itemInfo.ModTime().Format(time.RFC822),
			filepath.Ext(item.Name()),
		)
	}

	return nil
}
