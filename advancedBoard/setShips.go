package advancedBoard

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
)

func SetShips() {
	ui := gui.NewGUI(true)

	txt := gui.NewText(1, 1, "Press on any coordinate to log it.", nil)
	ui.Draw(txt)
	ui.Draw(gui.NewText(1, 2, "Press Ctrl+C to exit", nil))

	board := gui.NewBoard(1, 4, nil)
	ui.Draw(board)

	go func() {
		x := 1
		y := 1
		for {
			char := board.Listen(context.TODO())
			x, y = coordToInts(char)
			txt.SetText(fmt.Sprintf("Coordinate: %s", char))

			states := [10][10]gui.State{}
			states[x][y] = gui.Ship
			board.SetStates(states)
			ui.Log("Coordinate: %s", char) // logs are displayed after the game exits
		}
	}()
	ui.Start(context.TODO(), nil)
}
func coordToInts(coord string) (int, int) {
	if len(coord) == 3 {
		return int(coord[0] - 65), 9
	}
	return int(coord[0] - 65), int(coord[1] - 49)
}
func intsToCoord(x, y int) string {
	if y == 10 {
		return string(x) + "10"
	}
	return string(x) + string(y+'0')
}
