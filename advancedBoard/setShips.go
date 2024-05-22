package advancedBoard

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

var PlayerCoords [20]string

func SetShips() {
	fmt.Println("Czas ustawić statki...")
	time.Sleep(4 * time.Second)

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
					selectedX = append(selectedX[:i], selectedX[i+1:]...)
					selectedY = append(selectedY[:i], selectedY[i+1:]...)
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
			txt.SetText(fmt.Sprintf("Coordinate: %s ", char))
			board.SetStates(states)
			ui.Log("Coordinate: %s", char) // logs are displayed after the game exits
			if len(selectedX) == 20 {
				break
			}
		}
		for i := 0; i < 20; i++ {
			PlayerCoords[i] = intsToCoord(selectedX[i], selectedY[i])
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
	if y == 9 {
		return string(x+65) + "10"
	}
	return string(x+65) + string(y+49)
}
