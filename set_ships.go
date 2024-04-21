package main

import (
	"bufio"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"os"
	"strings"
)

var board = gui.New(gui.NewConfig())
var coords [10]string
var counter = int8(-1)

func set_ships() {
	fmt.Println("Ustaw swoje statki")
	fmt.Println("1 okręt 4-masztowy\n2 okręty 3-masztowe\n3 okręty 2-masztowe\n4 okręty 1-masztowe")
	fmt.Println("Podstawową zasadą konstrukcji okrętów bojowych (dłuższych niż 1) jest ich budowa z elementów sąsiadujących ze sobą bokami, a nie na skos.")
	fmt.Println("Od którego statku zaczniemy?")

	fmt.Println("Wybierz miejsce okrentu 4 masztowego")
	for range 4 {
		pick_place()
	}
	board.Display()
	fmt.Println("Wybierz miejsca okrentu 3 masztowego")
	for range 2 {
		for range 3 {
			pick_place()
		}
	}
	board.Display()
	fmt.Println("Wybierz miejsca okrentu 2 masztowego")
	for range 3 {
		for range 2 {
			pick_place()
		}
	}
	board.Display()
	fmt.Println("Wybierz miejsca okrentu 1 masztowego")
	for range 4 {
		pick_place()
	}
	board.Display()
}
func pick_place() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz Pozycję:")
	pos, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu pozycji:", err)
		return
	}
	pos = strings.TrimSpace(pos)
	fmt.Println(pos)
	board.Set(gui.Left, pos, gui.Ship)
	board.Display()
	counter += counter
	coords[counter] = pos
	return
}
func get_board() *gui.Board {
	return board
}
func get_coords() [10]string {
	return coords
}
