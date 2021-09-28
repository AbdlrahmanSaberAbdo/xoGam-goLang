package main

import "fmt"

func main() {
	xoBoard := [3][3]string{}
	player  := "X"

	for {
		fmt.Println("Player: ", player)

		var row int
		fmt.Println("please enter row number, 0, 1 or 2 ")
		fmt.Scanln(&row)  // expect input from user

		if invalidCell(row)  {
			fmt.Println("invalid input")
			continue
		}

		var column int
		fmt.Println("please enter column number, 0, 1 or 2")
		fmt.Scanln(&column)  // expect input from user

		if invalidCell(column) {
			fmt.Println("invalid input")
			continue
		}
		fmt.Println("row: ", row, "column: ", column)

		if isEmpty(xoBoard[row][column]) {
			xoBoard[row][column] = player
		} else {
			fmt.Println("Please choose other cell, this cell is field before")
			continue
		}

		fmt.Println(xoBoard[0])
		fmt.Println(xoBoard[1])
		fmt.Println(xoBoard[2])

		// Did someone win
		for i := 0; i < 3; i++ {
			// check rows || columns
			if isEqualInRowCase(xoBoard, i, player) || isEqualInColumnCase(xoBoard, i, player) {
				fmt.Println("Game ended, winner is player:", player)
				// we use return to exit both for loops and the app
				return
			}
		}
		if isEqualInThirdCase(xoBoard, player)|| isEqualInFourthCase(xoBoard, player) {
			fmt.Println("Game ended, winner is player:", player)
			// end the game and exit the app
			// we can use "break" or "return"
			return
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}

}

/*
	* if user enter number isn't in our xo-board
 */
func invalidCell(number int) bool {
	return number > 2
}

/*
* Check if cell is empty or not
 */
func isEmpty(cell string) bool {
	return cell == ""
}

/*
* Check if is equal in row case or not
 */
func isEqualInRowCase(xoBoard[3][3] string, i int, player string) bool {
	return xoBoard[i][0] == xoBoard[i][1] && xoBoard[i][1] == xoBoard[i][2] && xoBoard[i][2] == player
}

/*
* Check if is equal in column case or not
 */
func isEqualInColumnCase(xoBoard[3][3] string, i int, player string) bool {
	return xoBoard[0][i] == xoBoard[1][i] && xoBoard[1][i] == xoBoard[2][i] && xoBoard[2][i] == player
}

/*
* Check if is equal in third case, ex:
	x
		x
			x
 */
func isEqualInThirdCase(xoBoard[3][3] string, player string) bool {
	return xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == player

}

/*
* Check if is equal in third case, ex:
			x
		x
	x
*/
func isEqualInFourthCase(xoBoard[3][3] string, player string) bool {
	return xoBoard[0][0] == xoBoard[1][1] && xoBoard[1][1] == xoBoard[2][2] && xoBoard[2][2] == player

}