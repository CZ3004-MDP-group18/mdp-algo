package session

import (
	"fmt"
	"strings"
)

func parseArenaStr(arenaStr string, size int) [][]string {
	arenaStr = strings.ReplaceAll(arenaStr, " ", "")
	arenaStr = strings.ReplaceAll(arenaStr, "\n", "")
	arenaRune := []rune(arenaStr)
	if len(arenaRune) != size*size {
		panic(fmt.Sprintf("Arena string size is %v, should be %v", len(arenaRune), size*size))
	}

	res := make([][]string, size)
	for i := range res {
		res[i] = make([]string, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res[i][j] = string(arenaRune[i*size+j])
		}
	}
	return res
}
