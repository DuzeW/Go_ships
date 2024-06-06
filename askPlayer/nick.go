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
}
