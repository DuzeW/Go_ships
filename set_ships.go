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
var counter = int8(-1)

func setShips() {
	fmt.Println("Ustaw swoje statki")
	fmt.Println("1 okręt 4-masztowy\n2 okręty 3-masztowe\n3 okręty 2-masztowe\n4 okręty 1-masztowe")
	fmt.Println("Podstawową zasadą konstrukcji okrętów bojowych (dłuższych niż 1) jest ich budowa z elementów sąsiadujących ze sobą bokami, a nie na skos.")
	fmt.Println("Od którego statku zaczniemy?")

	fmt.Println("Wybierz miejsce okrentu 4 masztowego")

	fmt.Println("Wybierz miejsca okrentu 3 masztowego")

	fmt.Println("Wybierz miejsca okrentu 2 masztowego")

	fmt.Println("Wybierz miejsca okrentu 1 masztowego")
	isCorrect := getPos()
	for isCorrect == false {
		fmt.Println("Spróbuj ponownie")
		isCorrect = getPos()
	}
}

func piclace() {

	//board.Set(gui.Left, pos, gui.Ship)
	//board.Display()
	//counter = counter + 1
	//coords[counter] = pos
	//return
}

func getPos() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz Pozycję:")
	pos, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu pozycji:", err)
		return false
	}
	pos = strings.TrimSpace(pos)
	fmt.Println(pos)
	if len(pos) < 2 || len(pos) > 3 {
		fmt.Println("Nieodpowiednia długość")
		return false
	}
	if len(pos) == 3 && (pos[1] != 1 || pos[2] != 0) {
		fmt.Println("Nieodpowiednia liczba")
		return false
	}
	if len(pos) == 2 && (pos[1] < 1 || pos[1] > 9) {
		fmt.Println("Nieodpowiednia liczba")
		return false
	}
	if pos[0] < 'A' || pos[0] > 'J' {
		fmt.Println("Nieodpowiednia litera")
		return false
	}
	return true
}
