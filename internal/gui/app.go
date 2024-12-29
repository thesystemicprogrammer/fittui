package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/tomberch/fittui/internal/gui/state"
)

type App struct {
	app      *tview.Application
	mainGrid *tview.Grid
	state    state.Stater
}

func NewApp() *App {
	app := &App{
		app: tview.NewApplication(),
	}

	app.renderMain()
	app.run()
	return app
}

func (a *App) renderMain() {

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignLeft).
			SetTextStyle(tcell.StyleDefault.Foreground(tcell.ColorDeepSkyBlue)).
			SetText(text)
	}

	logo :=
		` ___ ___ _____
| __|_ _|_   _|
| _| | |  | |
|_| |___| |_|`

	box := tview.NewBox().
		SetBorder(true)

	infoGrid := tview.NewGrid().
		SetRows(-1, -1, -1, -1, -1).
		AddItem(newPrimitive("Info1"), 0, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info2"), 1, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info3"), 2, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info3"), 3, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info4"), 4, 0, 1, 1, 0, 0, false)

	menuGrid := tview.NewGrid().
		SetRows(0, 0, 0, 0, 0).
		AddItem(newPrimitive("Info1"), 0, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info2"), 1, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info3"), 2, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Info4"), 3, 0, 1, 1, 0, 0, false)

	contextGrid := tview.NewGrid().
		SetColumns(-36, -17, -17, -8).
		AddItem(infoGrid, 0, 0, 1, 1, 0, 0, false).
		AddItem(menuGrid, 0, 1, 1, 1, 0, 0, false).
		AddItem(newPrimitive("SubMenu"), 0, 2, 1, 1, 0, 0, false).
		AddItem(newPrimitive(logo), 0, 3, 1, 1, 0, 0, false)

	a.mainGrid = tview.NewGrid().
		SetRows(6, 0, 2).
		SetBorders(false).
		AddItem(contextGrid, 0, 0, 1, 1, 0, 0, false).
		AddItem(box, 1, 0, 1, 1, 0, 0, false).
		AddItem(newPrimitive("Footer"), 2, 0, 1, 1, 0, 0, false)
}

func (a *App) run() {
	if err := a.app.SetRoot(a.mainGrid, true).Run(); err != nil {
		panic(err)
	}
}
