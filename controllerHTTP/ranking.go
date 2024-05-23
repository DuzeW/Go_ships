package controllerHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PlayerStats struct {
	Games  int    `json:"games"`
	Nick   string `json:"nick"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Wins   int    `json:"wins"`
}

type StatsResponse struct {
	Stats []PlayerStats `json:"stats"`
}

var PlayerStatsList []PlayerStats

func GetPlayerStats() {
	URL := basicURL + "/stats"
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
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Błąd odpowiedzi serwera:", resp.Status)
		return
	}

	var statsResp StatsResponse
	err = json.NewDecoder(resp.Body).Decode(&statsResp)
	if err != nil {
		fmt.Println("Błąd dekodowania odpowiedzi JSON:", err)
		return
	}
	PlayerStatsList = statsResp.Stats
	for _, stat := range PlayerStatsList {
		fmt.Printf("Nick: %s, Gry: %d, Punkty: %d, Ranga: %d, Wygrane: %d\n", stat.Nick, stat.Games, stat.Points, stat.Rank, stat.Wins)
	}
}
