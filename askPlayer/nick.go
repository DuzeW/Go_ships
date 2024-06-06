package askPlayer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Nick string

func PlayerNick() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nWpisz swój nick:")
	nickInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu nicku:", err)
		return
	}
	Nick = strings.TrimSpace(nickInput)
	if len(Nick) > 20 || len(Nick) < 1 {
		fmt.Println("Zła długość nicku")
		PlayerNick()
	}
}
