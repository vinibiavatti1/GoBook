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
	"1": [...]int{0, 0},
	"2": [...]int{0, 1},
	"3": [...]int{0, 2},
	"4": [...]int{1, 0},
	"5": [...]int{1, 1},
	"6": [...]int{1, 2},
	"7": [...]int{2, 0},
	"8": [...]int{2, 1},
	"9": [...]int{2, 2},
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
		[...]string{"1", "2", "3"},
		[...]string{"4", "5", "6"},
		[...]string{"7", "8", "9"},
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

func (g *Game) ValidateWinner() bool {
	won := false
	won = won || g.IsSamePlayerInPositions("1", "2", "3")
	won = won || g.IsSamePlayerInPositions("4", "5", "6")
	won = won || g.IsSamePlayerInPositions("7", "8", "9")
	won = won || g.IsSamePlayerInPositions("1", "4", "7")
	won = won || g.IsSamePlayerInPositions("2", "5", "8")
	won = won || g.IsSamePlayerInPositions("3", "6", "9")
	won = won || g.IsSamePlayerInPositions("1", "5", "9")
	won = won || g.IsSamePlayerInPositions("3", "5", "7")
	return won
}

func (g *Game) ValidateTie() bool {
	for i := range g.board {
		for j := range g.board {
			pos := g.board[i][j]
			if pos != Player1 && pos != Player2 {
				return false
			}
		}
	}
	return true
}

func (g *Game) Validate() GameResult {
	switch {
	case g.ValidateWinner():
		return Won
	case g.ValidateTie():
		return Tie
	default:
		return None
	}
}

func (g *Game) IsSamePlayerInPositions(x, y, z string) bool {
	xy1 := PositionMap[x]
	xy2 := PositionMap[y]
	xy3 := PositionMap[z]
	x1, y1 := xy1[0], xy1[1]
	x2, y2 := xy2[0], xy2[1]
	x3, y3 := xy3[0], xy3[1]
	return g.board[x1][y1] == g.board[x2][y2] && g.board[x2][y2] == g.board[x3][y3]
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
			fmt.Printf("Player %s Won!\n", g.currentPlayer)
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
