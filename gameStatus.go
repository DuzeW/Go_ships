package main

import (
	"encoding/json"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
)

type GameResponse struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}

var isEnded bool = false
var lastGameStatus string

func gameStatus() bool {
	URL := "/game"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Błąd tworzenia zapytania:", err)
		return true
	}
	req.Header.Set("X-Auth-Token", tokenAPI)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Błąd wysyłania zapytania:", err)
		return true
	}
	defer resp.Body.Close()
	var data GameResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Błąd odbierania zapytania:", err)
		return true
	}
	if data.GameStatus == "ended" {
		isEnded = true
	}
	lastGameStatus = data.LastGameStatus
	fmt.Println("Nick:", data.Nick)
	fmt.Println("OppShots:")
	for _, shot := range data.OppShots {

		board.Set(gui.Left, shot, gui.Miss)
		//ZROBIĆ SPRAWDZENIE CZY POLE JEST PUSTE

		board.Set(gui.Left, shot, gui.Hit)
	}
	fmt.Println("Opponent:", data.Opponent)
	fmt.Println("ShouldFire:", data.ShouldFire)
	fmt.Println("Timer:", data.Timer)
	return false
}
