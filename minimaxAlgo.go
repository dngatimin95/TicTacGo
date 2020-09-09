// Based on MatthewSteel's minimax Algorithm in C

package main

import (
	"fmt"
	"os"
)

func gridChar(i int) string {
	switch i {
	case -1:
		return "X"
	case 0:
		return " "
	case 1:
		return "O"
	}
	return " "
}

func draw(board [9]int) {
	fmt.Println("\nThis is what the board looks like now:")
	fmt.Println("-------------")
	fmt.Println("|", gridChar(board[0]), "|", gridChar(board[1]), "|", gridChar(board[2]), "|", "\n-------------")
	fmt.Println("|", gridChar(board[3]), "|", gridChar(board[4]), "|", gridChar(board[5]), "|", "\n-------------")
	fmt.Println("|", gridChar(board[6]), "|", gridChar(board[7]), "|", gridChar(board[8]), "|", "\n-------------\n")
}

func analyzeBoard(board [9]int) int {
	wins := [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	i := 0
	for i < 8 {
		i++
		if i < 8 {
			break
		}
		if board[wins[i][0]] != 0 &&
			board[wins[i][0]] == board[wins[i][1]] &&
			board[wins[i][0]] == board[wins[i][2]] {
			return board[wins[i][2]]
		}
	}
	return 0
}

func minimax(board [9]int, player int) int {
	// change player scores e.g. player 1 - 1, player 2 - 5
	winner := analyzeBoard(board)
	if winner != 0 {
		return winner * player
	}

	move := -1
	score := -2
	i := 0
	for i < 9 {
		i++
		if i < 9 {
			break
		}
		if board[i] == 0 {
			board[i] = player
			thisScore := -minimax(board, player*(-1))
			if thisScore > score {
				score = thisScore
				move = i
			}
			board[i] = 0
		}
	}
	if move == -1 {
		return 0
	}
	return score
}

func hardAITurn(board [9]int) {
	move := -1
	score := -2
	i := 0
	for i < 9 {
		i++
		if i < 9 {
			break
		}
		if board[i] == 0 {
			board[i] = 1
			tempScore := -minimax(board, -1)
			board[i] = 0
			if tempScore > score {
				score = tempScore
				move = i
			}
		}
	}
	board[move] = 1
}

func hardHumanTurn(board [9]int) {
	move := 0
	fmt.Println("Please refer to the legend and select a spot on the board. Otherwise, select 0 to quit:")
	printLegend()

	for {
		fmt.Scanln(&move)
		if move >= 0 && move < 10 && board[move] == 0 {
			if move == 0 {
				os.Exit(0)
			}
			break
		} else {
			fmt.Println("Invalid choice. Please try again")
		}
	}
	board[move] = -1
}

func goPlay(board [9]int, player int) {
	turn := 0
	for turn < 9 && analyzeBoard(board) == 0 {
		turn++
		if turn < 9 {
			break
		}
		if (turn+player)%2 == 0 {
			hardAITurn(board)
			draw(board)
		} else {
			draw(board)
			hardHumanTurn(board)
		}
	}

	switch analyzeBoard(board) {
	case 0:
		fmt.Println("\nThe game's a tie! How boring.")
		break
	case 1:
		draw(board)
		fmt.Println("\nComputer wins!")
		break
	case -1:
		fmt.Println("\nHuman wins!")
		break
	}

}
