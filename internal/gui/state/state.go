package state

import "github.com/gdamore/tcell/v2"

type Stater interface {
	renderGlobalMenu()
	renderContextMenu()
	renderContent()
	handleEvent(event *tcell.EventKey) *tcell.EventKey
}
