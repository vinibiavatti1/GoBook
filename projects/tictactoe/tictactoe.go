// Tic Tac Toe Game
// This is a simple implementation of the Tic Tac Toe game in Go.
// The game is played on a 3x3 grid, where two players take turns placing their marks (X or O) in the empty cells.
// The first player to get three marks in a row (horizontally, vertically, or diagonally) wins the game.

package main

import (
	"fmt"
	"strings"
)

const (
	Player1 string = "X"
	Player2 string = "O"
	Empty   string = " "
	Len     int    = 3
	Div     string = " | "
)

type GameResult int

const (
	None GameResult = iota
	Won
	Tie
)

var PositionMap = map[string][2]int{
	"a": [...]int{0, 0},
	"b": [...]int{0, 1},
	"c": [...]int{0, 2},
	"d": [...]int{1, 0},
	"e": [...]int{1, 1},
	"f": [...]int{1, 2},
	"g": [...]int{2, 0},
	"h": [...]int{2, 1},
	"i": [...]int{2, 2},
}

type Game struct {
	board         [Len][Len]string
	currentPlayer string
}

func NewGame() *Game {
	game := &Game{
		currentPlayer: Player1,
	}
	game.Reset()
	return game
}

func (g *Game) Reset() {
	g.board = [...][Len]string{
		[...]string{"a", "b", "c"},
		[...]string{"d", "e", "f"},
		[...]string{"g", "h", "i"},
	}
}

func (g *Game) Play(pos string) error {
	coords, ok := PositionMap[pos]
	if !ok {
		return fmt.Errorf("Invalid Position: %s", pos)
	}
	x, y := coords[0], coords[1]
	boardpos := g.board[x][y]
	if boardpos == Player1 || boardpos == Player2 {
		return fmt.Errorf("Position is not empty")
	}
	g.board[x][y] = g.currentPlayer
	return nil
}

func (g *Game) SwapCurrentPlayer() {
	if g.currentPlayer == Player1 {
		g.currentPlayer = Player2
	} else {
		g.currentPlayer = Player1
	}
}

func (g *Game) Validate() GameResult {
	// Validate Rows
	won := false
	for i := range g.board {
		won = g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2]
		if won {
			return Won
		}
	}
	// Validate Cols
	for i := range g.board {
		won = g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i]
		if won {
			return Won
		}
	}
	// Validate Diagonal
	won = g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2]
	if won {
		return Won
	}
	// Validate Inverse Diagonal
	won = g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0]
	if won {
		return Won
	}
	// Validate Tie
	for i := range g.board {
		for j := range g.board {
			pos := g.board[i][j]
			if pos != Player1 && pos != Player2 {
				return None
			}
		}
	}
	return Tie
}

func (g *Game) PrintBoard() {
	for i := range g.board {
		line := strings.Join(g.board[i][:], Div)
		fmt.Println(line)
	}
}

func (g *Game) Run() {
	running := true
	for running {
		fmt.Println("Current Player:", g.currentPlayer)
		g.PrintBoard()
		for {
			pos := Scan()
			err := g.Play(pos)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		result := g.Validate()
		switch result {
		case None:
			g.SwapCurrentPlayer()
		case Tie:
			fmt.Println("Tie!!!")
			running = false
		case Won:
			fmt.Printf("Player %s won!\n", g.currentPlayer)
			running = false
		}
	}
	fmt.Println("Game finished")
	g.PrintBoard()
}

func Scan() string {
	for {
		var pos string
		fmt.Print("Position: ")
		_, err := fmt.Scanln(&pos)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return pos
	}
}

func main() {
	game := NewGame()
	game.Run()
}
