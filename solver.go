package godoku

import (
	"errors"
	"math/rand"
	"time"
)

//Solve a given puzzle, return an error if there is no solution
func Solve(board *[]byte, size byte) (*[]byte, error) {
	total := size * size
	tracker := NewTracker(int(size))
	noSol := false
	rand.Seed(time.Now().Unix())

	//Attempting to avoid a recursive algorithm.
	for i := byte(0); i < total; i++ {
		//there are no more unique numbers left to check
		if byte(len(tracker.attempts[i])) == size {
			if i == 0 {
				noSol = true
				break
			}
			tracker.Clear(int(i))
			(*board)[i] = 0
			i -= 2 //need to offset the post increment
			continue
		}

		newNum := availableNum(int(size), tracker.attempts[i])
		tracker.attempts[i][newNum] = true
		(*board)[i] = newNum

		if Validate(board, int(size)) {
			//we have found a valid number and can continue.
			continue
		}

		//Not a valid number, and we need to try again
		i--
	}
	if noSol {
		return nil, errors.New("ERR: No solution")
	}
	return board, nil
}
