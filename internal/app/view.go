package app

func (v *App) View() string {
	if v.Page == 0 {
		return v.Menu.View()
	} else {
		return v.GameView.View()
	}
}
