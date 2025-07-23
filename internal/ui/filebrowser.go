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

var (
	fileBrowser = logic.FileBrowserNew()
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

	column.SetClickable(true)
	column.SetReorderable(true)
	column.SetResizable(true)
	column.SetMinWidth(16)

	return column, nil
}

// TODO: use map[int(column index)]string for file info (name, size, etc.)
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

func setupFileTreeView() (*gtk.TreeView, error) {
	treeView, err := gtk.TreeViewNew()
	if err != nil {
		return nil, err
	}

	column1, err := createFileTreeViewColumn("Name", COLUMN_FILENAME)
	column2, err := createFileTreeViewColumn("Size", COLUMN_SIZE)
	column3, err := createFileTreeViewColumn("Modified", COLUMN_MODIFIED)
	column4, err := createFileTreeViewColumn("Type", COLUMN_EXTENSION)

	if err != nil {
		return nil, err
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

	err = fileBrowser.ChangeDirectory("/")
	if err != nil {
		return nil, err
	}
	updateFileTreeView(listStore, fileBrowser.CurrentDirectory())

	treeView.Connect("row-activated", func(tv *gtk.TreeView, tp *gtk.TreePath, tvc *gtk.TreeViewColumn) {
		model, err := tv.GetModel()
		if err != nil {
			log.Fatal(err)
		}

		listStore := model.(*gtk.ListStore)

		iter, err := listStore.GetIter(tp)
		if err != nil {
			log.Fatal(err)
		}

		columnValue, err := listStore.GetValue(iter, COLUMN_FILENAME)
		if err != nil {
			log.Fatal(err)
		}

		path, err := columnValue.GetString()
		if err != nil {
			log.Fatal(err)
		}

		err = fileBrowser.ChangeDirectory(
			filepath.Join(fileBrowser.CurrentDirectory(), path),
		)

		if err != nil {
			log.Fatal(err)
		}

		updateFileTreeView(listStore, fileBrowser.CurrentDirectory())
	})

	return treeView, nil
}

func updateFileTreeView(liststore *gtk.ListStore, path string) error {

	err := fileBrowser.ChangeDirectory(path)

	if err != nil {
		return err
	}

	contents, err := fileBrowser.GetCurrentDirContents()

	if err != nil {
		return err
	}

	liststore.Clear()

	var ext string

	for _, item := range contents {

		ext = filepath.Ext(item.Name())
		if ext == "" {
			ext = "dir"
		}

		addRow(
			liststore,
			item.Name(),
			fmt.Sprint(item.Size()),
			item.ModTime().Format(time.RFC822),
			ext,
		)
	}

	return nil
}
