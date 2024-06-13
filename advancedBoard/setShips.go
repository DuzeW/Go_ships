package advancedBoard

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

var PlayerCoords [20]string
var selectedX []int
var selectedY []int
var ships = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
var lockX []int
var lockY []int

func ClearBoard() {
	PlayerCoords = [20]string{}
	selectedX = []int{}
	selectedY = []int{}
	lockX = []int{}
	lockY = []int{}
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
					for i := 0; i < len(lockX); i++ {
						if lockX[i] == x && lockY[i] == y {
							badClick = true
						}
					}
					for i := 0; i < len(selectedX); i++ {
						if selectedX[i] == x && selectedY[i] == y {
							badClick = true
						}
					}
					if badClick {
						j--
						continue
					}
					//zaznaczanie
					if len(selectedX) < 20 && !badClick {
						selectedX = append(selectedX, x)
						selectedY = append(selectedY, y)
					}
					if ships[k]-1 == j {
						lock()
					}
					//rysowanie
					for i := 0; i < len(lockX); i++ {
						states[lockX[i]][lockY[i]] = gui.Miss
					}
					for i := 0; i < len(selectedX); i++ {
						states[selectedX[i]][selectedY[i]] = gui.Ship
					}
					board.SetStates(states)
					ui.Log("Coordinate: %s", char)
				}
			}
			txt1.SetText(fmt.Sprintf("Statki ustawiono prawidłowo naciśnij Ctrl + C aby przejść dalej"))
			break
		}
		for i := 0; i < 20; i++ {
			PlayerCoords[i] = intsToCoord(selectedX[i], selectedY[i])
		}
	}()
	ui.Start(context.TODO(), nil)
}
func lock() {
	for l := 0; l < len(selectedX); l++ {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if selectedX[l]+i > 9 || selectedX[l]+i < 0 || selectedY[l]+j > 9 || selectedY[l]+j < 0 {
					continue
				}
				lockX = append(lockX, selectedX[l]+i)
				lockY = append(lockY, selectedY[l]+j)
			}
		}
	}
}
func whichNumberInSelected(k int) int {
	i := 0
	for k > 0 {
		i += ships[k-1]
		k--
	}
	return i
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

func setAnalyzer() bool {
	if len(selectedX) != 20 {
		return false
	}
	if isCorrectShip1() {
		return true
	}
	return false

}
func isCorrectShip2() bool {
	shipCounter := 0
	for i := 0; i < 20; i++ {
		result, _, _ := isShipAround(selectedX[i], selectedY[i])
		if result == true {
			shipCounter++
		}
	}
	if shipCounter == 6 {
		return true
	}
	return false
}

func isCorrectShip1() bool {
	shipCounter := 0
	for i := 0; i < 20; i++ {
		result, _, _ := isShipAround(selectedX[i], selectedY[i])
		if result == false {
			shipCounter++
		}
	}
	if shipCounter == 4 {
		return true
	}
	return false
}

func isShipAround(x int, y int) (bool, int, int) {
	for i := -1; i < 2; i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < 20; j++ {
			if selectedY[j] == y && selectedX[j] == x+i {
				return true, x + i, y
			}
		}
		for j := 0; j < 20; j++ {
			if selectedX[j] == x && selectedY[j] == y+i {
				return true, x, y + i
			}
		}
	}
	return false, 0, 0
}
