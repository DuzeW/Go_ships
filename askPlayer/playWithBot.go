package askPlayer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var PlayWithBot bool

func PlayerPlayWithBot() {
	reader := bufio.NewReader(os.Stdin)
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
			PlayWithBot = true
			badAns = false
		}
		if ans == "N" || ans == "n" {
			PlayWithBot = false
			badAns = false
		}
	}
}
