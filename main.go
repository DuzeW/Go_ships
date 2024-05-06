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
		fmt.Println("Twój nick to: ", nick)
		fmt.Println("Twój opis to: ", desc)
		setShips()
		connect()
		for !isEnded {
			fire()
			checkBoard()
			gameStatus()
			board.Display()
		}
		fmt.Println("Gra dobiegła końca wynik to:", lastGameStatus)
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
