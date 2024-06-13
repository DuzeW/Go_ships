package advancedBoard

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

var PlayerCoords [20]string
var selectedXY [][]int
var ships = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
var lockXY [][]int

func ClearBoard() {
	PlayerCoords = [20]string{}
	selectedXY = [][]int{}
	lockXY = [][]int{}
}

func SetShips() {
	fmt.Println("Czas ustawić statki...")
	time.Sleep(3 * time.Second)

	ui := gui.NewGUI(true)

	txt1 := gui.NewText(1, 1, "Naciśnij pola gdzie chcesz ustawić statki", nil)
	ui.Draw(txt1)
	txt2 := gui.NewText(1, 2, "", nil)
	ui.Draw(txt2)

	board := gui.NewBoard(1, 4, nil)
	ui.Draw(board)

	go func() {
		for {
			states := [10][10]gui.State{}
			for k := 0; k < len(ships); k++ {
				for j := 0; j < ships[k]; j++ {
					states = [10][10]gui.State{}
					txt1.SetText(fmt.Sprintf("Ustawiasz statek długości %d część %d", ships[k], j+1))
					char := board.Listen(context.TODO())
					x, y := coordToInts(char)
					badClick := false
					for i := 0; i < len(lockXY); i++ {
						if lockXY[i][0] == x && lockXY[i][1] == y {
							badClick = true
						}
					}
					for i := 0; i < len(selectedXY); i++ {
						if selectedXY[i][0] == x && selectedXY[i][1] == y {
							badClick = true
						}
					}
					if badClick {
						j--
						continue
					}
					//zaznaczanie
					if len(selectedXY) < 20 && !badClick {
						selectedXY = append(selectedXY, []int{x, y})
					}
					if ships[k]-1 == j {
						lock()
					}
					//rysowanie
					for i := 0; i < len(lockXY); i++ {
						states[lockXY[i][0]][lockXY[i][1]] = gui.Miss
					}
					for i := 0; i < len(selectedXY); i++ {
						states[selectedXY[i][0]][selectedXY[i][1]] = gui.Ship
					}
					board.SetStates(states)
					ui.Log("Coordinate: %s", char)
				}
			}
			txt1.SetText(fmt.Sprintf("Statki ustawiono prawidłowo naciśnij Ctrl + C aby przejść dalej"))
			break
		}
		for i := 0; i < 20; i++ {
			PlayerCoords[i] = intsToCoord(selectedXY[i][0], selectedXY[i][1])
		}
	}()
	ui.Start(context.TODO(), nil)
}

func lock() {
	for l := 0; l < len(selectedXY); l++ {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if selectedXY[l][0]+i > 9 || selectedXY[l][0]+i < 0 || selectedXY[l][1]+j > 9 || selectedXY[l][1]+j < 0 {
					continue
				}
				lockXY = append(lockXY, []int{selectedXY[l][0] + i, selectedXY[l][1] + j})
			}
		}
	}
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
