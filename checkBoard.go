package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var reqBoard []string

func checkBoard() {
	URL := BASIC_URL + "/game/board"
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
	var data map[string][]string
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Błąd odbierania zapytania:", err)
		return
	}
	reqBoard := data["board"]
	err = board.Import(reqBoard)
	if err != nil {
		fmt.Println("Błąd podczas pobierania tablicy błąd:", err)
		return
	}
	fmt.Println(reqBoard)
}
