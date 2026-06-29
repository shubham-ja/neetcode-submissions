func isValidSudoku(board [][]byte) bool {
    rows := make([]int, 9)
    cols := make([]int, 9)
    squares := make([]int, 9)
    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
                continue
            }

            val := board[r][c] - '1'
			
            bit := 1 << val
            squareIdx := (r/3)*3 + c/3
			fmt.Println(r,c, squareIdx)
            if rows[r]&bit != 0 || cols[c]&bit != 0 ||
               squares[squareIdx]&bit != 0 {
                return false
            }

            rows[r] |= bit
            cols[c] |= bit
            squares[squareIdx] |= bit
        }
    }

    return true
}