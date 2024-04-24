package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"io/ioutil"
	"net/http"
)

func fire() {
	URL := BASIC_URL + "/game/fire"
	isCorrect, coord := getCoord()
	for !isCorrect {
		isCorrect, coord = getCoord()
	}
	data := map[string]string{"coord": coord}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Błąd podczas konwersji danych na JSON:", err)
		return
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Błąd tworzenia żądania:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", tokenAPI)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Błąd wykonania żądania:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Błąd odczytu odpowiedzi:", err)
		return
	}
	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Błąd parsowania odpowiedzi:", err)
		return
	}
	fmt.Println("Result:", result["result"])
	if result["result"] == "miss" {
		board.Set(gui.Right, coord, gui.Miss)
	}
	if result["result"] == "hit" || result["result"] == "sunk" {
		board.Set(gui.Right, coord, gui.Hit)
	}

}