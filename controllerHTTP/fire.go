package controllerHTTP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fire(coord string) string {
	URL := basicURL + "/game/fire"

	data := map[string]string{"coord": coord}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Błąd podczas konwersji danych na JSON:", err)
		return ""
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Błąd tworzenia żądania:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", tokenAPI)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Błąd wykonania żądania:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Błąd odczytu odpowiedzi:", err)
		return ""
	}
	var result map[string]string
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Błąd parsowania odpowiedzi:", err)
		return ""
	}
	return result["result"]
}
