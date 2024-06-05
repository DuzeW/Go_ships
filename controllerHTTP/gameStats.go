package controllerHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

var LastGameStatus string
var Status string
var OppShots []string
var ShouldFire bool
var Timer int

func GameStatus() {
	for {
		URL := basicURL + "/game"
		req, err := http.NewRequest("GET", URL, nil)
		if err != nil {
			fmt.Println("Błąd tworzenia zapytania:", err)
			return
		}
		req.Header.Set("X-Auth-Token", tokenAPI)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Błąd wysyłania zapytania:", err)
			time.Sleep(250 * time.Millisecond)
			continue
		}
		defer resp.Body.Close()
		var data GameResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			fmt.Println("Błąd odbierania zapytania:", err)
			time.Sleep(250 * time.Millisecond)
			continue
		}
		Status = data.GameStatus
		LastGameStatus = data.LastGameStatus
		OppShots = data.OppShots
		ShouldFire = data.ShouldFire
		Timer = data.Timer
		return
	}
}
