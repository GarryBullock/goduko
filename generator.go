package godoku

//Generate a game board
func Generate(size int) *[]byte {
	total := byte(size * size)
	board := make([]byte, total)

	soln, err := Solve(&board, byte(size))
	if err != nil {
		panic("Should be impossible to have an unsolveable puzzle when creating a new puzzle.")
	}

	return soln
}
