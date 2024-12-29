package menu

import "github.com/gdamore/tcell/v2"

type MenuItem struct {
	title string
	key   *tcell.Key
}

func NewMenuItem(title string, key *tcell.Key) *MenuItem {
	return &MenuItem{
		title: title,
		key:   key,
	}
}
