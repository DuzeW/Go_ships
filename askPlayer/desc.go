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
	fmt.Println("\nWpisz swój opis:")
	descInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Błąd podczas wpisywaniu nicku:", err)
		return
	}
	Desc = strings.TrimSpace(descInput)
	if len(Desc) > 100 {
		fmt.Println("Zła długość opisu")
		PlayerDesc()
	}
}
