package main

import (
	"fmt"
	"ships/advancedBoard"
	"ships/askPlayer"
	"ships/controllerHTTP"
	"ships/game"
	"time"
)

func main() {

	askPlayer.PlayerNick()
	askPlayer.PlayerDesc()
	controllerHTTP.WaitingList()
	askPlayer.ChooseOp()
	advancedBoard.SetShips()
	controllerHTTP.Connect()
	controllerHTTP.GameStatus()
	for controllerHTTP.Status != "game_in_progress" {
		time.Sleep(1 * time.Second)
		controllerHTTP.GameStatus()
		fmt.Println("Oczekiwanie na grę")
	}
	game.ShowBoards()

	//	for {
	//		setShips()
	//		connect()
	//		gameStatusFail := gameStatus()
	//		for gameStatusFail == true {
	//			time.Sleep(1 * time.Second)
	//			gameStatusFail = gameStatus()
	//		}
	//		for !isEnded {
	//			fire()
	//			checkBoard()
	//			gameStatus()
	//			board.Display()
	//		}
	//		fmt.Println("Gra dobiegła końca wynik to:", lastGameStatus)
	//		playAgain()
	//	}
	//
	//}
	//func playAgain() {
	//	reader := bufio.NewReader(os.Stdin)
	//	for {
	//		fmt.Println("Czy chcesz zagrać ponownie?(T/N)")
	//		again, err := reader.ReadString('\n')
	//		if err != nil {
	//			fmt.Println("Błąd podczas wpisywania:", err)
	//		}
	//		again = strings.TrimSpace(again)
	//		if again == "T" || again == "t" {
	//		}
	//		if again == "N" || again == "n" {
	//			os.Exit(0)
	//		}
	//	}
}
