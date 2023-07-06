package main

func solve(board *[9][9]int) bool {
	if isSolved(board) {
		printBoard(board)
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for v := 1; v < 10; v++ {
					move := Move{
						Coordinates: Coordinates{
							x: j,
							y: i,
						},
						Value: v,
					}
					if isValidMove(board, move) {
						board[i][j] = v

						if solve(board) {
							return true
						} else {
							board[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isSolved(board *[9][9]int) bool {
	for i := range board {
		for _, val := range board[i] {
			if val == 0 {
				return false
			}
		}
	}

	return true
}

func isValidMove(board *[9][9]int, move Move) bool {
	return isValidMoveForRow(board, move) && isValidMoveForColumn(board, move) && isValidMoveForRegion(board, move)
}

func isValidMoveForRow(board *[9][9]int, move Move) bool {
	for _, val := range board[move.Coordinates.y] {
		if val == move.Value {
			return false
		}
	}
	return true
}

func isValidMoveForColumn(board *[9][9]int, move Move) bool {
	for i := range board {
		if board[i][move.Coordinates.x] == move.Value {
			return false
		}
	}
	return true
}

func isValidMoveForRegion(board *[9][9]int, move Move) bool {
	startRow := move.Coordinates.y / 3 * 3
	startCol := move.Coordinates.x / 3 * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == move.Value {
				return false
			}
		}
	}
	return true
}
