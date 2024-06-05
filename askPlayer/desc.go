package askPlayer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Desc string

func PlayerDesc() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Wpisz swój opis:")
	descInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu nicku:", err)
		return
	}
	Desc = strings.TrimSpace(descInput)
}
