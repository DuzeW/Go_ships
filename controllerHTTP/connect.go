package controllerHTTP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"ships/advancedBoard"
	"ships/askPlayer"
)

var basicURL = "https://go-pjatk-server.fly.dev/api"
var tokenAPI string

func Connect() {
	URL := basicURL + "/game"

	data := map[string]interface{}{
		"coords":      advancedBoard.PlayerCoords,
		"desc":        askPlayer.Desc,
		"nick":        askPlayer.Nick,
		"target_nick": askPlayer.OpNick,
		"wpbot":       askPlayer.PlayWithBot,
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
	//time.Sleep(40 * time.Second)
}
func IsConnect() bool {
	if len(tokenAPI) < 5 {
		return false
	}
	return true
}
