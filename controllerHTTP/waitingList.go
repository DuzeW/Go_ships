package controllerHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WaitingPlayer struct {
	GameStatus string `json:"game_status"`
	Nick       string `json:"nick"`
}

func WaitingList() {
	URL := basicURL + "/lobby"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Błąd tworzenia zapytania:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Błąd wysyłania zapytania:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Błąd odpowiedzi serwera:", resp.Status)
		return
	}

	var players []WaitingPlayer
	err = json.NewDecoder(resp.Body).Decode(&players)
	if err != nil {
		fmt.Println("Błąd dekodowania odpowiedzi JSON:", err)
		return
	}
	fmt.Println("Lista graczy oczekujących na grę:")
	for _, player := range players {
		fmt.Printf("Nick: %s, Status: %s\n", player.Nick, player.GameStatus)
	}
}
