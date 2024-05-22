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
		var selectedX []int
		var selectedY []int
		for {

			states := [10][10]gui.State{}
			char := board.Listen(context.TODO())
			x, y = coordToInts(char)
			isSelected := false
			for i := 0; i < len(selectedX); i++ {
				if selectedX[i] == x && selectedY[i] == y {
					selectedX = append(selectedX[:i], selectedX[i:]...)
					selectedY = append(selectedY[:i], selectedY[i:]...)
					states[selectedX[i]][selectedY[i]] = gui.Empty
					isSelected = true
				}
			}
			if len(selectedX) < 20 && isSelected == false {
				selectedX = append(selectedX, x)
				selectedY = append(selectedY, y)
			}

			for i := 0; i < len(selectedX); i++ {
				states[selectedX[i]][selectedY[i]] = gui.Ship
			}
			txt.SetText(fmt.Sprintf("Coordinate: %s %t\n", char, isSelected))
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