package controllerHTTP

import (
	"fmt"
	"net/http"
)

func AbandonGame() {
	URL := basicURL + "/game/abandon"
	req, err := http.NewRequest("DELETE", URL, nil)
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

	fmt.Println("Gra została porzucona pomyślnie.")
}
