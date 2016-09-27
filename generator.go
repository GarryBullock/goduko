package godoku

import (
	"fmt"
	"math/rand"
	"time"
)

//Generate a game board
func Generate(size int) {
	start := time.Now()
	total := byte(size * size)
	board := make([]byte, total)
	fmt.Println(board)
	rand.Seed(time.Now().Unix())

	for i := 0; i < int(total); i++ {
		//loop and use random integers until we find a valid solution
		maxattempts := 20
		for {
			entry := rand.Intn(size-1) + 1
			//fmt.Println(i, entry)
			board[i] = byte(entry)
			if Validate(&board, size) || maxattempts == 0 {
				break
			}
			maxattempts--
			//time.Sleep(100 * time.Millisecond)
		}
		if maxattempts == 0 {
			board = make([]byte, total)
			i = 0
		}
	}

	fmt.Println(board)
	fmt.Printf("Elapsed time to generate: %v\n", time.Since(start))
}
