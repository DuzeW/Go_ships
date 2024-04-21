package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func connect() {
	URL := BASIC_URL + "/game"
	data := map[string]interface{}{
		"coords":      coords,
		"desc":        desc,
		"nick":        nick,
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
}
