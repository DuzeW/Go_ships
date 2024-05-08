package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var tokenAPI string

func connect() {

	URL := BASIC_URL + "/game"
	var coords = []string{"A1", "A2", "A3", "A4", "C1", "C2", "C3", "E1", "E2", "E3", "G1", "G2", "I1", "I2", "A10", "B10", "J10", "I8", "A8", "E6"}
	data := map[string]interface{}{
		"coords":      coords,
		"desc":        desc,
		"nick":        nick,
		"target_nick": "",
		"wpbot":       playWithBot,
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
	tokenAPI = resp.Header.Get("X-Auth-Token")
	fmt.Println(tokenAPI)
}
