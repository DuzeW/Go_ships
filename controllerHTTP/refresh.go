package controllerHTTP

import (
	"fmt"
	"net/http"
)

func Refresh() {
	URL := basicURL + "/game/refresh"
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
}
