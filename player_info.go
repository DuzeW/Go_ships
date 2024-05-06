package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nick string
var desc string
var playWithBot bool

func playerInfo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz nick:")
	nickInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu nicku:", err)
		return
	}
	nick = nickInput
	fmt.Println("Wpisz opis:")
	descInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu opisu:", err)
		return
	}
	desc = descInput
	badAns := true
	for badAns {
		fmt.Println("Czy chcesz zagrać z botem?(T/N)")
		ans, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Błąd podczas wpisywaniu wyboru gry z botem:", err)
			return
		}
		ans = strings.TrimSpace(ans)
		if ans == "T" || ans == "t" {
			playWithBot = true
			badAns = false
		}
		if ans == "N" || ans == "n" {
			playWithBot = false
			badAns = false
		}
	}
	nick = strings.TrimSpace(nick)
	desc = strings.TrimSpace(desc)
	return
}
