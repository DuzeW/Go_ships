package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var BASIC_URL = "https://go-pjatk-server.fly.dev/api"

func main() {
	for {
		playerInfo()
		setShips()
		connect()
		for len(shotsR) != 20 || len(reqBoard) == 0 {
			fire()
			checkBoard()
			board.Display()
		}
		if len(shotsR) != 20 {
			fmt.Println("Zwycięstwo")
		}
		if len(reqBoard) == 0 {
			fmt.Println("Przegrana")
		}
		playAgain()
	}

}
func playAgain() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Czy chcesz zagrać ponownie?(T/N)")
		again, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Błąd podczas wpisywania:", err)
		}
		again = strings.TrimSpace(again)
		if again == "T" || again == "t" {
		}
		if again == "N" || again == "n" {
			os.Exit(0)
		}
	}
}
