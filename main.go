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
	//askPlayer.ShouldShowRanking()
	//if askPlayer.ShowRanking {
	//	controllerHTTP.GetPlayerStats()
	//}
	//askPlayer.PlayerNick()
	askPlayer.Nick = "t"
	//askPlayer.PlayerDesc()
	askPlayer.Desc = "t"
	//controllerHTTP.WaitingList()
	//askPlayer.ChooseOp()
	askPlayer.OpNick = ""
	askPlayer.PlayWithBot = true
	//advancedBoard.SetShips()
	advancedBoard.PlayerCoords = [20]string{"A1", "A2", "A3", "A4", "A6", "A7", "A8", "C6", "C7", "C8", "D1", "D2", "F1", "F2", "H1", "H2", "H10", "J5", "J7", "J9"}
	controllerHTTP.Connect()
	controllerHTTP.GameStatus()
	for controllerHTTP.Status != "game_in_progress" {
		controllerHTTP.Refresh()
		time.Sleep(1 * time.Second)
		controllerHTTP.GameStatus()
		fmt.Println("Oczekiwanie na grę")
	}
	game.ShowBoards()
	controllerHTTP.AbandonGame()
	fmt.Println(controllerHTTP.LastGameStatus)
	time.Sleep(5 * time.Second)

}
