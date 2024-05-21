package main

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
)

var BASIC_URL = "https://go-pjatk-server.fly.dev/api"

func main() {
	ui := gui.NewGUI(true)

	board := gui.NewBoard(1, 1, nil)
	ui.Draw(board)

	states := [10][10]gui.State{}
	for i := range states {
		states[i] = [10]gui.State{}
	}
	board.SetStates(states)

	ctx := context.Background()
	ui.Start(ctx, nil)

	txt := gui.NewText(1, 1, "Press Ctrl+C to exit", nil)
	ui.Draw(txt)
	ui.Start(context.TODO(), nil)

	txt = gui.NewText(1, 1, "Press on any coordinate to log it.", nil)
	ui.Draw(txt)
	ui.Draw(gui.NewText(1, 2, "Press Ctrl+C to exit", nil))

	board = gui.NewBoard(1, 4, nil)
	ui.Draw(board)

	go func() {
		for {
			char := board.Listen(context.TODO())
			txt.SetText(fmt.Sprintf("Coordinate: %s", char))
			ui.Log("Coordinate: %s", char) // logs are displayed after the game exits
		}
	}()

	ui.Start(context.TODO(), nil)
	//	for {
	//		playerInfo()
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
