package menu

type MainMenu struct {
	items []*MenuItem
}

func NewMainMenu(items []*MenuItem) *MainMenu {
	return &MainMenu{
		items,
	}
}
