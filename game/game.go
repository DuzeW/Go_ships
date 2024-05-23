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
	ui.Draw(txt)
	ui.Draw(gui.NewText(1, 2, "Press Ctrl+C to exit", nil))

	board := gui.NewBoard(1, 4, nil)
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
	opBoard := gui.NewBoard(50, 4, nil)
	opBoard.SetStates(opStates)
	ui.Draw(opBoard)

	var miss []string
	var hit []string
	go func() {
		controllerHTTP.GameStatus()
		for {
			controllerHTTP.GameStatus()
			if controllerHTTP.ShouldFire {
				char := opBoard.Listen(context.TODO())
				result := controllerHTTP.Fire(char)
				if result == "miss" {
					miss = append(miss, char)
				}
				if result == "hit" || result == "sunk" {
					hit = append(hit, char)
				}
				if result == "sunk" {
					for i := 0; i < len(hit); i++ {
						x, y := coordToInts(hit[i])
						for j := -1; j < 2; j++ {
							for k := -1; k < 2; k++ {
								if x+j <= 9 && x+j >= 0 && y+k <= 9 && y+k >= 0 {
									miss = append(miss, intsToCoord(x+j, y+k))
								}
							}
						}
					}
				}
				for i := 0; i < len(miss); i++ {
					x, y := coordToInts(miss[i])
					opStates[x][y] = gui.Miss
				}
				for i := 0; i < len(hit); i++ {
					x, y := coordToInts(hit[i])
					opStates[x][y] = gui.Hit
				}
				txt.SetText(fmt.Sprintf("Twój ruch Coordinate: %s %s Pozostały czas %d", char, result, controllerHTTP.Timer))
				opBoard.SetStates(opStates)
				ui.Log("Coordinate: %s", char)
			} else {
				txt.SetText(fmt.Sprintf("Ruch przeciwnika"))
				for i := 0; i < len(controllerHTTP.OppShots); i++ {
					x, y := coordToInts(controllerHTTP.OppShots[i])
					states[x][y] = gui.Miss
				}
				for i := 0; i < len(controllerHTTP.OppShots); i++ {
					x, y := coordToInts(controllerHTTP.OppShots[i])
					for j := 0; j < len(myCoords); j++ {
						if controllerHTTP.OppShots[i] == myCoords[j] {
							states[x][y] = gui.Hit
						}
					}
				}
				board.SetStates(states)
			}
			controllerHTTP.GameStatus()
			if controllerHTTP.Status != "game_in_progress" {
				break
			}
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
