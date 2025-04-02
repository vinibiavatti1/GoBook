package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"time"
)

type Cell string

const (
	Live Cell = "*"
	Dead Cell = " "
)

type Rule struct {
	Born    []int
	Survive []int
}

var (
	Gol          = &Rule{Born: []int{3}, Survive: []int{2, 3}}
	HighLife     = &Rule{Born: []int{2, 3}, Survive: []int{3, 6}}
	Assimilation = &Rule{Born: []int{3, 4, 5}, Survive: []int{4, 5, 6, 7}}
	TwoXTwo      = &Rule{Born: []int{3, 6}, Survive: []int{1, 2, 5}}
	DayAndNight  = &Rule{Born: []int{3, 6, 7, 8}, Survive: []int{3, 4, 5, 7, 8}}
	Amoeba       = &Rule{Born: []int{3, 5, 7}, Survive: []int{1, 3, 5, 8}}
	Move         = &Rule{Born: []int{3, 6, 8}, Survive: []int{2, 4, 5}}
	PseudoLife   = &Rule{Born: []int{3, 5, 7}, Survive: []int{2, 3, 8}}
	Diamoeba     = &Rule{Born: []int{3, 5, 6, 7, 8}, Survive: []int{5, 6, 7, 8}}
	Rule34       = &Rule{Born: []int{3, 4}, Survive: []int{3, 4}}
	LongLife     = &Rule{Born: []int{3, 4, 5}, Survive: []int{5}}
	Stains       = &Rule{Born: []int{3, 6, 7, 8}, Survive: []int{2, 3, 5, 6, 7, 8}}
	Seeds        = &Rule{Born: []int{2}, Survive: []int{}}
	Maze         = &Rule{Born: []int{3}, Survive: []int{1, 2, 3, 4, 5}}
	Coagulations = &Rule{Born: []int{3, 7, 8}, Survive: []int{2, 3, 5, 6, 7, 8}}
	WalledCities = &Rule{Born: []int{4, 5, 6, 7, 8}, Survive: []int{2, 3, 4, 5}}
	Gnarl        = &Rule{Born: []int{1}, Survive: []int{1}}
	Replicator   = &Rule{Born: []int{1, 3, 5, 7}, Survive: []int{1, 3, 5, 7}}
	Mystery      = &Rule{Born: []int{3, 4, 5, 8}, Survive: []int{0, 5, 6, 7, 8}}
)

type update struct {
	x    int
	y    int
	cell Cell
}

type GameOfLife struct {
	Width  int
	Height int
	Rule   *Rule
	screen [][]Cell
}

func NewGameOfLife(w int, h int, rule *Rule) *GameOfLife {
	gol := &GameOfLife{
		Width:  w,
		Height: h,
		Rule:   rule,
	}
	gol.Reset()
	return gol
}

func (g *GameOfLife) Reset() {
	g.screen = make([][]Cell, g.Height)
	for y := range g.Height {
		g.screen[y] = make([]Cell, g.Width)
		for x := range g.Width {
			g.screen[y][x] = Dead
		}
	}
}

func (g *GameOfLife) Get(x, y int) Cell {
	if x < 0 || y < 0 {
		return Dead
	}
	if x >= g.Width || y >= g.Height {
		return Dead
	}
	return g.screen[y][x]
}

func (g *GameOfLife) Set(c Cell, x, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= g.Width || y >= g.Height {
		return false
	}
	g.screen[y][x] = c
	return true
}

func (g *GameOfLife) Render() {
	for y := range g.Height {
		for x := range g.Width {
			c := g.Get(x, y)
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func (g *GameOfLife) Next() {
	updates := []*update{}
	for y := range g.Height {
		for x := range g.Width {
			cell := g.Get(x, y)
			lives := g.countLives(x, y)
			if cell == Live && !slices.Contains(g.Rule.Survive, lives) {
				updates = append(updates, &update{x: x, y: y, cell: Dead})
			} else if slices.Contains(g.Rule.Born, lives) {
				updates = append(updates, &update{x: x, y: y, cell: Live})
			}
		}
	}
	for _, v := range updates {
		g.Set(v.cell, v.x, v.y)
	}
}

func (g *GameOfLife) countLives(x, y int) int {
	cells := [8]Cell{
		g.Get(x-1, y-1),
		g.Get(x-0, y-1),
		g.Get(x+1, y-1),
		g.Get(x-1, y-0),
		g.Get(x+1, y-0),
		g.Get(x-1, y+1),
		g.Get(x-0, y+1),
		g.Get(x+1, y+1),
	}
	count := 0
	for _, c := range cells {
		if c == Live {
			count++
		}
	}
	return count
}

func run(cmd string) {
	c := exec.Command("cmd", "/c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

func main() {
	gol := NewGameOfLife(60, 20, Gol)
	gol.Set(Live, 30, 8)
	gol.Set(Live, 30, 9)
	gol.Set(Live, 30, 10)
	gol.Set(Live, 29, 9)
	gol.Set(Live, 31, 10)
	run("@echo off && cls")
	for {
		gol.Next()
		gol.Render()
		time.Sleep(time.Millisecond * 100)
		run("cls")
	}
}
