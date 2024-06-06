package askPlayer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ShowRanking bool

func ShouldShowRanking() {
	reader := bufio.NewReader(os.Stdin)
	badAns := true
	for badAns {
		fmt.Println("\nCzy chcesz zobaczyć ranking graczy?(T/N)")
		ans, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Błąd podczas wpisywaniu:", err)
			return
		}
		ans = strings.TrimSpace(ans)
		if ans == "T" || ans == "t" {
			ShowRanking = true
			badAns = false
		}
		if ans == "N" || ans == "n" {
			ShowRanking = false
			badAns = false
		}
	}
}
