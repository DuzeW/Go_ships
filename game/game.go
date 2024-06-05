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

	txt := gui.NewText(1, 1, "", nil)
	takenCoords := gui.NewText(1, 3, "Wybierz miejsce gdzie chcesz oddać strzał", nil)
	ui.Draw(txt)
	ui.Draw(takenCoords)
	ui.Draw(gui.NewText(1, 2, "Naciśnij Ctrl+C aby wyjść z gry", nil))

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

	var miss []string
	var hit []string
	var checked []string
	go func() {
		for {
			if controllerHTTP.Timer > 0 {
				txt.SetText(fmt.Sprintf("Czas na ruch %d", controllerHTTP.Timer))
			}
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			if controllerHTTP.ShouldFire {
				char := opBoard.Listen(context.TODO())
				ui.Log("Coordinate: %s", char)
				isAlreadyChecked := false
				for i := 0; i < len(checked); i++ {
					if checked[i] == char {
						isAlreadyChecked = true
					}
				}
				if isAlreadyChecked == false {
					checked = append(checked, char)
				} else {
					continue
				}
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
			}
			opBoard.SetStates(opStates)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for {
			if controllerHTTP.ShouldFire {
				takenCoords.SetText(fmt.Sprintf("Wybierz miejsce gdzie chcesz oddać strzał"))
			}
			if !controllerHTTP.ShouldFire && controllerHTTP.Status != "ended" {
				takenCoords.SetText(fmt.Sprintf("Ruch przeciwnika"))
			}
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
			time.Sleep(200 * time.Millisecond)
		}
	}()
	go func() {
		for {
			controllerHTTP.GameStatus()
			if controllerHTTP.Status == "ended" {
				txt.SetText(fmt.Sprintf(""))
				if controllerHTTP.LastGameStatus == "win" {
					takenCoords.SetText(fmt.Sprintf("ZWYCIĘSTWO"))
				}
				if controllerHTTP.LastGameStatus == "lose" {
					takenCoords.SetText(fmt.Sprintf("PORAŻKA"))
				}
				ui.Log("End of game LastGameStatus: %s", controllerHTTP.LastGameStatus)
				break
			}
			time.Sleep(500 * time.Millisecond)
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
