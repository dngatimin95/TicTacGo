package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func startMenu() int {
	for {
		fmt.Println("\nPlease select a numerical choice from below:")
		fmt.Println("1. Play Tic-Tac-Go vs Easy AI \n2. Play Tic-Tac-Go vs Hard AI \n3. Play Tic-Tac-Go vs Another Player \n4. Exit\n")
		var choice int
		fmt.Scanln(&choice)

		if choice < 5 && choice > 0 {
			return choice
		} else {
			fmt.Println("\nInvalid choice. Please try again")
		}
	}
}

func endMenu() bool {
	for {
		fmt.Println("Do you want to play again? Y/N")
		var playAgain string
		fmt.Scanln(&playAgain)
		playAgain = strings.TrimSpace(playAgain)
		playAgain = strings.ToLower(playAgain)

		if playAgain == "y" {
			return true
		} else if playAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye")
			os.Exit(0)
		} else {
			fmt.Println("Invalid choice. Please try again")
		}
	}
}

func startGame(choice int) bool {
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	if choice == 3 {
		playManualGame(board)
	} else {
		playAIGame(board, choice)
	}
	return endMenu()
}

func playManualGame(board [9]int) {
	for turn := 0; turn < 9; turn++ {
		if analyzeBoard(board) != 0 || turn >= 9 {
			break
		}
		if ((turn) % 2) == 0 {
			printBoard(board)
			fmt.Println("Player 1's turn.")
			board = human1Turn(board)
		} else {
			printBoard(board)
			fmt.Println("Player 2's turn.")
			board = human2Turn(board)
		}
	}

	printBoard(board)
	switch analyzeBoard(board) {
	case 0:
		fmt.Println("\nThe game's a tie! How boring.")
		break
	case 1:
		printBoard(board)
		fmt.Println("\nPlayer 2 wins!")
		break
	case -1:
		fmt.Println("\nPlayer 1 wins!")
		break
	}
}

func playAIGame(board [9]int, choice int) {
	player := rand.Intn(2) + 1

	if player == 1 {
		fmt.Println("\nPlayer starts first!")
	} else {
		fmt.Println("\nComputer starts first!")
	}

	if choice == 1 {
		playEasy(board, player)
	} else {
		playHard(board, player)
	}

}

func human1Turn(board [9]int) [9]int {
	move := 0
	fmt.Println("Please refer to the legend and select a spot on the board. Otherwise, select 0 to quit:")
	printLegend()

	for {
		fmt.Scanln(&move)
		if move >= 0 && move < 10 && board[move-1] == 0 {
			if move == 0 {
				fmt.Println("Thanks for playing! Goodbye")
				os.Exit(0)
			}
			board[move-1] = -1
			break
		} else {
			fmt.Println("Invalid choice. Please try again")
		}
	}
	return board
}

func human2Turn(board [9]int) [9]int {
	move := 0
	fmt.Println("Please refer to the legend and select a spot on the board. Otherwise, select 0 to quit:")
	printLegend()

	for {
		fmt.Scanln(&move)
		if move >= 0 && move < 10 && board[move-1] == 0 {
			if move == 0 {
				fmt.Println("Thanks for playing! Goodbye")
				os.Exit(0)
			}
			board[move-1] = 1
			break
		} else {
			fmt.Println("Invalid choice. Please try again")
		}
	}
	return board
}

func easyAITurn(board [9]int) [9]int {
	aiMove := rand.Intn(9)
	for {
		if board[aiMove] != 0 {
			aiMove = rand.Intn(9)
		} else {
			break
		}
	}
	board[aiMove] = 1
	return board
}

func hardAITurn(board [9]int) [9]int {
	move := -1
	score := -2
	for i := 0; i < 9; i++ {
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
	return board
}

func analyzeBoard(board [9]int) int {
	wins := [8][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
	for i := 0; i < 8; i++ {
		if board[wins[i][0]] != 0 &&
			board[wins[i][0]] == board[wins[i][1]] &&
			board[wins[i][0]] == board[wins[i][2]] {
			return board[wins[i][2]]
		}
	}
	return 0
}

func minimax(board [9]int, player int) int {
	winner := analyzeBoard(board)
	if winner != 0 {
		return winner * player
	}

	move := -1
	score := -2

	for i := 0; i < 9; i++ {
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

func playEasy(board [9]int, player int) {
	for turn := 0; turn < 9; turn++ {
		if analyzeBoard(board) != 0 || turn >= 9 {
			break
		}
		if ((turn + player) % 2) == 0 {
			board = easyAITurn(board)
			printBoard(board)
		} else {
			printBoard(board)
			board = human1Turn(board)
		}
	}

	endPrint(board)
}

func playHard(board [9]int, player int) {
	for turn := 0; turn < 9; turn++ {
		if analyzeBoard(board) != 0 || turn >= 9 {
			break
		}
		if ((turn + player) % 2) == 0 {
			board = hardAITurn(board)
			printBoard(board)
		} else {
			printBoard(board)
			board = human1Turn(board)
		}
	}

	endPrint(board)
}

func endPrint(board [9]int) {
	printBoard(board)
	switch analyzeBoard(board) {
	case 0:
		fmt.Println("\nThe game's a tie! How boring.")
		break
	case 1:
		printBoard(board)
		fmt.Println("\nComputer wins!")
		break
	case -1:
		fmt.Println("\nHuman wins!")
		break
	}
}

func boardSym(i int) string {
	switch i {
	case -1:
		return "X"
	case 0:
		return " "
	case 1:
		return "O"
	}
	return ""
}

func printBoard(board [9]int) {
	fmt.Println("\nThis is what the board looks like now:")
	fmt.Println("-------------")
	fmt.Println("|", boardSym(board[0]), "|", boardSym(board[1]), "|", boardSym(board[2]), "|", "\n-------------")
	fmt.Println("|", boardSym(board[3]), "|", boardSym(board[4]), "|", boardSym(board[5]), "|", "\n-------------")
	fmt.Println("|", boardSym(board[6]), "|", boardSym(board[7]), "|", boardSym(board[8]), "|", "\n-------------\n")
}

func printLegend() {
	fmt.Println("\nLegend: \n-------------")
	fmt.Println("| 1 | 2 | 3 | \n-------------")
	fmt.Println("| 4 | 5 | 6 | \n-------------")
	fmt.Println("| 7 | 8 | 9 | \n-------------\n")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	restart := true

	for restart != false {
		opt := startMenu()
		if opt == 4 {
			fmt.Println("Thanks for playing! Goodbye")
			os.Exit(0)
		} else {
			restart = startGame(opt)
		}
	}
}
