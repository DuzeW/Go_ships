package advancedBoard

import (
	"context"
	gui "github.com/grupawp/warships-gui/v2"
)

func NewAdvancedBoard() {
	ui := gui.NewGUI(true)
	txt := gui.NewText(1, 1, "Press Ctrl+C to exit", nil)
	ui.Draw(txt)

	board := gui.NewBoard(1, 3, nil)
	ui.Draw(board)

	states := [10][10]gui.State{}
	for i := range states {
		states[i] = [10]gui.State{}
	}
	board.SetStates(states)

	ctx := context.Background()
	ui.Start(ctx, nil)
}
