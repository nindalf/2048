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

type grid [][]int

func (g grid) MoveRight() (moved bool) {
	for i := range g {
		m := g.collapseRowRight(i)
		moved = m || moved
		// combine
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == g[i][j-1] && g[i][j] != sen {
				g[i][j-1] = sen
				g[i][j] *= 2
				j--
				moved = true
			}
		}
		m = g.collapseRowRight(i)
		moved = m || moved
	}
	return
}

func (g grid) collapseRowRight(ri int) (moved bool) {
	row := g[ri]
	si := -1
	for j := len(row) - 1; j >= 0; j-- {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si--
			moved = true
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
	return
}

func (g grid) MoveLeft() (moved bool) {
	for i := range g {
		m := g.collapseRowLeft(i)
		moved = m || moved
		// combine
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == g[i][j+1] && g[i][j] != sen {
				g[i][j+1] = sen
				g[i][j] *= 2
				j++
				moved = true
			}
		}
		m = g.collapseRowLeft(i)
		moved = m || moved
	}
	return
}

func (g grid) collapseRowLeft(ri int) (moved bool) {
	row := g[ri]
	si := -1
	for j := 0; j < len(row); j++ {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si++
			moved = true
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
	return
}

func (g grid) MoveDown() (moved bool) {
	for j := range g[0] {
		m := g.collapseColDown(j)
		moved = m || moved
		for i := len(g) - 1; i > 0; i-- {
			if g[i][j] == g[i-1][j] && g[i][j] != sen {
				g[i][j] *= 2
				g[i-1][j] = sen
				i--
				moved = true
			}
		}
		m = g.collapseColDown(j)
		moved = m || moved
	}
	return
}

func (g grid) collapseColDown(ci int) (moved bool) {
	si := -1
	for i := len(g) - 1; i >= 0; i-- {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si--
			moved = true
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
	return
}

func (g grid) MoveUp() (moved bool) {
	for j := range g[0] {
		m := g.collapseColUp(j)
		moved = m || moved
		for i := 0; i < len(g)-1; i++ {
			if g[i][j] == g[i+1][j] && g[i][j] != sen {
				g[i][j] *= 2
				g[i+1][j] = sen
				i++
				moved = true
			}
		}
		m = g.collapseColUp(j)
		moved = m || moved
	}
	return
}

func (g grid) collapseColUp(ci int) (moved bool) {
	si := -1
	for i := 0; i < len(g); i++ {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si++
			moved = true
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
	return
}

func (g grid) AddNumber() {
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

func (g grid) Win() bool {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == winningScore {
				return true
			}
		}
	}
	return false
}

func (g grid) Full() bool {
	for i := range g {
		for j := range g[i] {
			if g[i][j] == sen {
				return false
			}
		}
	}
	return true
}

func (g grid) Cheatcode() bool {
	var x [2]int // store the max value
	for i := range g {
		for j := range g[i] {
			if g[x[0]][x[1]] < g[i][j] {
				x[0], x[1] = i, j
			}
		}
	}
	g[x[0]][x[1]] *= 2
	return true
}

func (g grid) String() string {
	var b bytes.Buffer
	for i := range g {
		b.Write([]byte(fmt.Sprintf("%v\n", g[i])))
	}
	return b.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cheatsOn, _ := strconv.ParseBool(os.Getenv("CHEATS_ENABLED"))
	reader := bufio.NewReader(os.Stdin)

	g := grid{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0}}
	g.AddNumber()
	g.AddNumber()

	fmt.Printf("Try to score %d!\n", winningScore)
	for {
		fmt.Printf("----------\n%v", g)
		fmt.Println("1 - Up, 2 - Down, 3 - Left, 4 - Right")

		// Windows workaround for reading Stdin
		t, errread := reader.ReadString('\n')
		i, errconv := strconv.Atoi(t[0:1])
		if errread != nil || errconv != nil || !((i > 0 && i < 5) || i == 9) {
			fmt.Println("Exiting the game. See you soon!")
			break
		}
		var moved bool
		switch i {
		case 1:
			moved = g.MoveUp()
		case 2:
			moved = g.MoveDown()
		case 3:
			moved = g.MoveLeft()
		case 4:
			moved = g.MoveRight()
		case 9:
			if cheatsOn {
				moved = g.Cheatcode()
			}
		}
		if !moved {
			continue
		}
		if g.Win() {
			fmt.Printf("----------\n%v", g)
			fmt.Printf("You won! You reached %d!\n", winningScore)
			break
		}
		g.AddNumber()
		if g.Full() {
			fmt.Println("Game over :(")
			break
		}
	}
	// To keep the screen visible on windows
	time.Sleep(2 * time.Second)
}
