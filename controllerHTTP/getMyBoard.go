package controllerHTTP

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BoardResponse struct {
	Board []string `json:"board"`
}

var MyBoard []string

func GetMyBoard() []string {
	URL := basicURL + "/game/board"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Błąd tworzenia zapytania:", err)
		return []string{}
	}
	req.Header.Set("X-Auth-Token", tokenAPI)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Błąd wysyłania zapytania:", err)
		return []string{}
	}
	defer resp.Body.Close()
	var boardResp BoardResponse
	err = json.NewDecoder(resp.Body).Decode(&boardResp)
	if err != nil {
		fmt.Println("Błąd dekodowania odpowiedzi JSON:", err)
		return []string{}
	}
	return boardResp.Board
}
