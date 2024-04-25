package main

import (
	"bufio"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"os"
	"strings"
)

var board = gui.New(gui.NewConfig())
var coords [20]string

func setShips() {
	fmt.Println("Ustaw swoje statki")
	fmt.Println("1 okręt 4-masztowy\n2 okręty 3-masztowe\n3 okręty 2-masztowe\n4 okręty 1-masztowe")
	fmt.Println("Podstawową zasadą konstrukcji okrętów bojowych (dłuższych niż 1) jest ich budowa z elementów sąsiadujących ze sobą bokami, a nie na skos.")
	fmt.Println("Od którego statku zaczniemy?")

	for i := 0; i < 20; i++ {
		switch i {
		case 0, 1, 2, 3:
			fmt.Println("Wybierz miejsce okrentu 4 masztowego")
		case 4, 5, 6, 7, 8, 9:
			fmt.Println("Wybierz miejsca okrentu 3 masztowego")
		case 10, 11, 12, 13, 14, 15:
			fmt.Println("Wybierz miejsca okrentu 2 masztowego")
		case 16, 17, 18, 19:
			fmt.Println("Wybierz miejsca okrentu 1 masztowego")
		default:
			fmt.Println("Błąd podczas ustawiania statku")
		}
		coord := getCorrectCoord()
		coords[i] = coord
		board.Set(gui.Left, coord, gui.Ship)
		board.Display()
		fmt.Println(coords)
	}
}

func getCoord() (bool, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz Koordynaty:")
	coord, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu koordynatów:", err)
		return false, ""
	}
	coord = strings.TrimSpace(coord)
	if len(coord) < 2 || len(coord) > 3 {
		fmt.Println("Nieodpowiednia długość")
		return false, ""
	}
	if len(coord) == 3 && (coord[1] != '1' || coord[2] != '0') {
		fmt.Println("Nieodpowiednia liczba")
		return false, ""
	}
	if len(coord) == 2 && (coord[1] < '1' || coord[1] > '9') {
		fmt.Println("Nieodpowiednia liczba tu")
		return false, ""
	}
	if coord[0] < 'A' || coord[0] > 'J' {
		fmt.Println("Nieodpowiednia litera")
		return false, ""
	}
	return true, coord
}

func isEmptyCoord(coord string) bool {
	for i := 0; i < 20; i++ {
		if coord == coords[i] {
			return false
		}
	}
	return true
}
func getCorrectCoord() string {
	isCorrect, coord := getCoord()
	isCorrect = isEmptyCoord(coord)
	for !isCorrect {
		fmt.Println("Spróbuj ponownie")
		isCorrect = isEmptyCoord(coord)
		isCorrect, coord = getCoord()
	}
	return coord
}
