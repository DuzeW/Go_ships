package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
)

var URL = "https://go-pjatk-server.fly.dev/api/game"

func main() {
	fmt.Println("pętla!")
	data := map[string]interface{}{
		"coords":      []string{"A1", "A3", "B9", "C7", "D1", "D2", "D3", "D4", "D7", "E7", "F1", "F2", "F3", "F5", "G5", "G8", "G9", "I4", "J4", "J8"},
		"desc":        "mydesc",
		"nick":        "custom123",
		"target_nick": "",
		"wpbot":       true,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Błąd:", err)
		return
	}

	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Błąd:", err)
	}
	defer resp.Body.Close()

	fmt.Println("APIkey")
	var tokenAPI = resp.Header.Get("X-Auth-Token")
	tokenAPI = tokenAPI[1 : len(tokenAPI)-1]
	fmt.Println(tokenAPI)

	board := gui.New(gui.NewConfig())
	board.Display()
	board.Set(gui.Left, "A1", gui.Ship)
	board.Set(gui.Left, "A2", gui.Ship)
	board.Display()
}
