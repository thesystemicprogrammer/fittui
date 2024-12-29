package state

import "github.com/gdamore/tcell/v2"

type InitialState struct {
}

func NewInitialState() *InitialState {
	return &InitialState{}
}

func (i *InitialState) renderGlobalMenu() {}

func (i *InitialState) renderContextMenu() {}

func (i *InitialState) renderContent() {}

func (i *InitialState) handleEvent(event *tcell.EventKey) *tcell.EventKey {
	return event
}
