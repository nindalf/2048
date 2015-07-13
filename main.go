package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// sentinel
var sen = 0
var winningScore = 64

type board struct {
	grid  [][]int
	moved bool
}

func (b board) MoveRight() {
	g := b.grid
	for i := range g {
		b.collapseRowRight(i)
		// combine
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == g[i][j-1] && g[i][j] != sen {
				g[i][j-1] = sen
				g[i][j] *= 2
				j--
				b.moved = true
			}
		}
		b.collapseRowRight(i)
	}
}

func (b board) collapseRowRight(ri int) {
	row := b.grid[ri]
	si := -1
	for j := len(row) - 1; j >= 0; j-- {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si--
			b.moved = true
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (b board) MoveLeft() {
	g := b.grid
	for i := range g {
		b.collapseRowLeft(i)
		// combine
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == g[i][j+1] && g[i][j] != sen {
				g[i][j+1] = sen
				g[i][j] *= 2
				j++
				b.moved = true
			}
		}
		b.collapseRowLeft(i)
	}
}

func (b board) collapseRowLeft(ri int) {
	row := b.grid[ri]
	si := -1
	for j := 0; j < len(row); j++ {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si++
			b.moved = true
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (b board) MoveDown() {
	g := b.grid
	for j := range g[0] {
		b.collapseColDown(j)
		for i := len(g) - 1; i > 0; i-- {
			if g[i][j] == g[i-1][j] && g[i][j] != sen {
				g[i][j] *= 2
				g[i-1][j] = sen
				i--
				b.moved = true
			}
		}
		b.collapseColDown(j)
	}
}

func (b board) collapseColDown(ci int) {
	g := b.grid
	si := -1
	for i := len(g) - 1; i >= 0; i-- {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si--
			b.moved = true
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (b board) MoveUp() {
	g := b.grid
	for j := range g[0] {
		b.collapseColUp(j)
		for i := 0; i < len(g)-1; i++ {
			if g[i][j] == g[i+1][j] && g[i][j] != sen {
				g[i][j] *= 2
				g[i+1][j] = sen
				i++
				b.moved = true
			}
		}
		b.collapseColUp(j)
	}
}

func (b board) collapseColUp(ci int) {
	g := b.grid
	si := -1
	for i := 0; i < len(g); i++ {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si++
			b.moved = true
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (b board) AddNumber() {
	g := b.grid
	possibles := make([][2]int, 0, len(g)*len(g[0]))
	for i := range g {
		for j := range g[i] {
			if g[i][j] == sen {
				possibles = append(possibles, [2]int{i, j})
			}
		}
	}

	x := possibles[rand.Intn(len(possibles))]
	g[x[0]][x[1]] = 2
}

func (b board) Win() bool {
	g := b.grid
	for i := range g {
		for j := range g[i] {
			if g[i][j] == winningScore {
				return true
			}
		}
	}
	return false
}

func (b board) Full() bool {
	g := b.grid
	for i := range g {
		for j := range g[i] {
			if g[i][j] == sen {
				return false
			}
		}
	}
	return true
}

func (b board) Cheatcode() {
	g := b.grid
	var x [2]int // store the max value
	for i := range g {
		for j := range g[i] {
			if g[x[0]][x[1]] < g[i][j] {
				x[0], x[1] = i, j
			}
		}
	}
	g[x[0]][x[1]] *= 2
	b.moved = true
}

func (b board) String() string {
	var buf bytes.Buffer
	for i := range b.grid {
		buf.Write([]byte(fmt.Sprintf("%v\n", b.grid[i])))
	}
	return buf.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cheatsOn, _ := strconv.ParseBool(os.Getenv("CHEATS_ENABLED"))
	reader := bufio.NewReader(os.Stdin)

	grid := [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0}}
	b := board{grid: grid, moved: false}
	b.AddNumber()
	b.AddNumber()

	fmt.Printf("Try to score %d!\n", winningScore)
	for {
		fmt.Printf("----------\n%v", b)
		fmt.Println("1 - Up, 2 - Down, 3 - Left, 4 - Right")

		// Windows workaround for reading Stdin
		t, errread := reader.ReadString('\n')
		i, errconv := strconv.Atoi(t[0:1])
		if errread != nil || errconv != nil || !((i > 0 && i < 5) || i == 9) {
			fmt.Println("Exiting the game. See you soon!")
			break
		}
		b.moved = false
		switch i {
		case 1:
			b.MoveUp()
		case 2:
			b.MoveDown()
		case 3:
			b.MoveLeft()
		case 4:
			b.MoveRight()
		case 9:
			if cheatsOn {
				b.Cheatcode()
			}
		}
		if !b.moved {
			continue
		}
		if b.Win() {
			fmt.Printf("----------\n%v", b)
			fmt.Printf("You won! You reached %d!\n", winningScore)
			break
		}
		b.AddNumber()
		if b.Full() {
			fmt.Println("Game over :(")
			break
		}
	}
	// To keep the screen visible on windows
	time.Sleep(2 * time.Second)
}
