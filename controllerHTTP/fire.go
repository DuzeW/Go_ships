package controllerHTTP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Fire(coord string) string {
	URL := basicURL + "/game/fire"
	data := map[string]string{"coord": coord}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Błąd podczas konwersji danych na JSON:", err)
		return ""
	}

	allowedResponses := map[string]bool{
		"miss": true,
		"sunk": true,
		"hit":  true,
	}

	for i := 0; i < 3; i++ {
		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
		if err != nil {
			return "Błąd tworzenia żądania"
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Auth-Token", tokenAPI)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "Błąd wykonania żądania:"
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return "Błąd odczytu odpowiedzi"
			} else {
				var result map[string]string
				err = json.Unmarshal(body, &result)
				if err != nil {
					return "Błąd parsowania odpowiedzi"
				} else {
					if msg, exists := result["result"]; exists {
						if allowedResponses[msg] {
							return msg
						} else {
							return "Nieoczekiwana odpowiedź"
						}
					}
				}
			}
		}
		time.Sleep(1 * time.Second)
	}

	return "błąd strzału"
}
