package game

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"ships/controllerHTTP"
	"time"
)

func ShowBoards() {
	ui := gui.NewGUI(true)

	txt := gui.NewText(1, 1, "Press on any coordinate to log it.", nil)
	takenCoords := gui.NewText(1, 3, "Press on any coordinate to log it.", nil)
	ui.Draw(txt)
	ui.Draw(takenCoords)
	ui.Draw(gui.NewText(1, 2, "Naciśnij Ctrl+C żeby się poddać", nil))

	board := gui.NewBoard(1, 5, nil)
	myCoords := controllerHTTP.GetMyBoard()
	//myCoords := []string{"A1", "A2", "A3", "A4", "A6", "A7", "A8", "C6", "C7", "C8", "D1", "D2", "F1", "F2", "H1", "H2", "H10", "J5", "J7", "J9"}

	states := [10][10]gui.State{}
	for i := 0; i < len(myCoords); i++ {
		x, y := coordToInts(myCoords[i])
		states[x][y] = gui.Ship
	}
	board.SetStates(states)
	ui.Draw(board)

	opStates := [10][10]gui.State{}
	opBoard := gui.NewBoard(50, 5, nil)
	opBoard.SetStates(opStates)
	ui.Draw(opBoard)

	//var miss []string
	//var hit []string
	go func() {
		i := 60
		for {
			i--
			txt.SetText(fmt.Sprintf("Czas na ruch %d", i))
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			char := opBoard.Listen(context.TODO())
			takenCoords.SetText(fmt.Sprintf("Czas na ruch %s", char))
		}
	}()
	go func() {
		for {
			controllerHTTP.GameStatus()
			time.Sleep(1 * time.Second)
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
