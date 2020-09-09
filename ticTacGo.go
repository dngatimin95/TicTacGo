package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func startMenu() int {
	for {
		fmt.Println("\nPlease select a numerical choice from below:")
		fmt.Println("1. Play Tic-Tac-Go vs Easy AI \n2. Play Tic-Tac-Go vs Hard AI <Still WIP> \n3. Play Tic-Tac-Go vs Another Player \n4. Exit\n")
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

func playAIGame(board [9]int, choice int) {
	turn := 1
	stopGame := false
	ai := true
	player := rand.Intn(2) + 1

	if player == 1 {
		fmt.Println("\nPlayer starts first!")
	} else {
		fmt.Println("\nComputer starts first!")
	}

	if choice == 1 { //easyAI
		if player == 1 { //human first
			for stopGame != true {
				player = 1
				printBoard(board)
				turn, stopGame, board = humanTurn(turn, stopGame, player, board, ai)

				if stopGame == true {
					break
				}

				player = 2
				turn, stopGame, board = easyAITurn(turn, stopGame, player, board)
			}
		} else { //Ai first
			for stopGame != true {
				player = 1
				turn, stopGame, board = easyAITurn(turn, stopGame, player, board)

				if stopGame == true {
					break
				}
				printBoard(board)

				player = 2
				turn, stopGame, board = humanTurn(turn, stopGame, player, board, ai)
			}
		}
	} else { //hardAI
		goPlay(board, player)
	}
	printBoard(board)
}

func humanTurn(turn int, stopGame bool, player int, board [9]int, ai bool) (int, bool, [9]int) {
	move := userInput()
	if move == 0 {
		fmt.Println("Thanks for playing! Goodbye")
		os.Exit(0)
	}
	board = executeInput(move-1, player, board)
	res := checkWin(board)

	if ai == true {
		if player == res {
			fmt.Println("\nHuman wins!")
			stopGame = true
		} else if res > 0 {
			fmt.Println("\nComputer wins!")
			stopGame = true
		} else if turn >= 9 {
			fmt.Println("\nThe game's a tie! How boring.")
			stopGame = true
		} else {
			turn++
		}
	} else if res > 0 {
		fmt.Println("\nPlayer", res, "wins!")
		stopGame = true
	} else if turn >= 9 {
		fmt.Println("\nThe game's a tie! How boring.")
		stopGame = true
	} else {
		turn++
	}

	return turn, stopGame, board
}

func easyAITurn(turn int, stopGame bool, player int, board [9]int) (int, bool, [9]int) {
	aiMove := rand.Intn(9)
	occupied := false
	for occupied != true {
		if board[aiMove] != 0 {
			aiMove = rand.Intn(9)
		} else {
			occupied = true
		}
	}

	if player == 1 {
		board[aiMove] = 1
	} else if player == 2 {
		board[aiMove] = 5
	}
	res := checkWin(board)

	if player == res {
		fmt.Println("\nComputer wins!")
		stopGame = true
	} else if res > 0 {
		fmt.Println("\nHuman wins!")
		stopGame = true
	} else if turn >= 9 {
		fmt.Println("\nThe game's a tie! How boring.")
		stopGame = true
	} else {
		turn++
	}
	return turn, stopGame, board
}

func playManualGame(board [9]int) {
	var player int
	turn := 1
	stopGame := false
	ai := false

	for stopGame != true {
		printBoard(board)
		if turn%2 == 1 {
			fmt.Println("Player 1's turn.")
			player = 1
		} else {
			fmt.Println("Player 2's turn.")
			player = 2
		}

		turn, stopGame, board = humanTurn(turn, stopGame, player, board, ai)
	}
	printBoard(board)
}

func userInput() int {
	var move int
	fmt.Println("Please refer to the legend and select a spot on the board. Otherwise, select 0 to quit:")
	printLegend()

	for {
		fmt.Scanln(&move)
		if move >= 0 && move < 10 {
			return move
		} else {
			fmt.Println("Invalid choice. Please try again")
		}
	}
}

func executeInput(move int, player int, board [9]int) [9]int {
	if board[move] != 0 {
		fmt.Println("This space is occupied. Try again.")
		move = userInput()
		if move == 0 {
			fmt.Println("Thanks for playing! Goodbye")
			os.Exit(0)
		}
		board = executeInput(move-1, player, board)
	} else {
		if player == 1 {
			board[move] = 1
		} else if player == 2 {
			board[move] = 5
		}
	}
	return board
}

func checkWin(board [9]int) int {
	var x [8]int
	x[0] = board[0] + board[3] + board[6]
	x[1] = board[1] + board[4] + board[7]
	x[2] = board[2] + board[5] + board[8]
	x[3] = board[0] + board[1] + board[2]
	x[4] = board[3] + board[4] + board[5]
	x[5] = board[6] + board[7] + board[8]
	x[6] = board[0] + board[4] + board[8]
	x[7] = board[2] + board[4] + board[6]

	for _, v := range x {
		if v == 3 {
			return 1
		} else if v == 15 {
			return 2
		}
	}
	return 0
}

func printBoard(board [9]int) {
	var display [9]string

	for i := 0; i < len(board); i++ {
		switch board[i] {
		case 0:
			display[i] = " "
		case 1:
			display[i] = "X"
		case 5:
			display[i] = "O"
		}
	}

	fmt.Println("\nThis is what the board looks like now:")
	fmt.Println("-------------")
	fmt.Println("|", display[0], "|", display[1], "|", display[2], "|", "\n-------------")
	fmt.Println("|", display[3], "|", display[4], "|", display[5], "|", "\n-------------")
	fmt.Println("|", display[6], "|", display[7], "|", display[8], "|", "\n-------------\n")
}

func printLegend() {
	fmt.Println("\nLegend: \n-------------")
	fmt.Println("| 1 | 2 | 3 | \n-------------")
	fmt.Println("| 4 | 5 | 6 | \n-------------")
	fmt.Println("| 7 | 8 | 9 | \n-------------\n")
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	restart := true

	for restart != false {
		opt := startMenu()
		if opt == 4 {
			os.Exit(0)
		} else {
			restart = startGame(opt)
		}
	}
}
