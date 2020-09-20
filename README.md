# TicTacGo
Tic Tac Toe is a simple game that most of us grew up with. It is basically a game where two players take turns to complete three boxes either in a row, a column or a diagonal with X's or O's drawn within the boxes. This script allows users to play the game either with other users, with an easy AI or a harder AI. As this script was meant for me to practice coding in Go and to understand more about the language, it is not fully optimized and some parts may be redundant.

## So what does this repo do?
As mentioned earlier, users are able to choose between three game modes: either against another player locally, against an easy AI or a harder AI. If the user chooses to play against other people, each player will take turns inputting their choices on the board and the first to get three in a row will win. When against an easy AI, it is completely random who starts first. Afterwards, either the player or the computer will take turns to fill in the board. During this mode, choices made by the computer are completely random and thus, might be easier to win against. On the last mode, against a harder AI, who starts first is also completely random. However, the choices made by the computer are based on the minimax algorithm. The minimax algorithm is described by **Geeksforgeeks** as follows:

>Minimax is a kind of backtracking algorithm that is used in decision making and game theory to find the optimal move for a player, assuming that your opponent also plays optimally. In Minimax, the two players are called maximizer and minimizer. The maximizer tries to get the highest score possible while the minimizer tries to do the opposite and get the lowest score possible. Every board state has a value associated with it. In a given state if the maximizer has upper hand then, the score of the board will tend to be some positive value. If the minimizer has the upper hand in that board state then it will tend to be some negative value.

Thus, it might be hard for players to win against the computer as it will always make the optimal move to win the game. It would be interesting to see if someone manages to win against the computer.

## How do I run the script?
Simple! Just download the repo, run the script, and have fun! e.g. go run TicTacGo.go
