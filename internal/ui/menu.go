package ui

import (
	"github.com/gotk3/gotk3/gtk"
)

type RightClickContextMenu struct {
	*gtk.Menu
}

func (rccm *RightClickContextMenu) Popup(path string) {
	rccm.PopupAtPointer(nil)
}

func RightClickContextMenuNew() (*RightClickContextMenu, error) {
	rightClickContextMenu := RightClickContextMenu{}

	menu, err := gtk.MenuNew()
	if err != nil {
		return nil, err
	}

	menuItem, err := gtk.MenuItemNewWithLabel("Open...")
	if err != nil {
		return nil, err
	}
	menuItem.Show()

	menu.Append(menuItem)

	rightClickContextMenu.Menu = menu

	return &rightClickContextMenu, nil
}
