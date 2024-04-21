package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func player_info() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz nick:")
	nick, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu nicku:", err)
		return "", ""
	}
	fmt.Println("Wpisz opis:")
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu opisu:", err)
		return "", ""
	}
	nick = strings.TrimSpace(nick)
	desc = strings.TrimSpace(desc)
	fmt.Println("Twój nick to: ", nick)
	fmt.Println("Twój opis to: ", desc)
	return nick, desc
}
