package checkers

import (
	"log"
)

func handleError(err error) {
	log.Println(err)
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
