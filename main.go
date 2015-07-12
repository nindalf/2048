package main

import (
	"bytes"
	"fmt"
)

// sentinel
var sen = 0

type grid [][]int

func (g grid) MoveRight() {
	for i := range g {
		g.collapseRowRight(i)
		// combine
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == g[i][j-1] {
				g[i][j-1] = sen
				g[i][j] *= 2
				j--
			}
		}
		g.collapseRowRight(i)
	}
}

func (g grid) collapseRowRight(ri int) {
	row := g[ri]
	si := -1
	for j := len(row) - 1; j >= 0; j-- {
		if row[j] != sen && si >= 0 {
			// fmt.Println(j, si)
			row[j], row[si] = row[si], row[j]
			si--
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (g grid) MoveLeft() {
	for i := range g {
		g.collapseRowLeft(i)
		// combine
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == g[i][j+1] {
				g[i][j+1] = sen
				g[i][j] *= 2
				j++
			}
		}
		g.collapseRowLeft(i)
	}
}

func (g grid) collapseRowLeft(ri int) {
	row := g[ri]
	si := -1
	for j := 0; j < len(row); j++ {
		if row[j] != sen && si >= 0 {
			row[j], row[si] = row[si], row[j]
			si++
		}
		if si == -1 && row[j] == sen {
			si = j
		}
	}
}

func (g grid) MoveDown() {
	for j := range g[0] {
		g.collapseColDown(j)
		for i := len(g) - 1; i > 0; i-- {
			if g[i][j] == g[i-1][j] {
				g[i][j] *= 2
				g[i-1][j] = sen
				i--
			}
		}
		g.collapseColDown(j)
	}
}

func (g grid) collapseColDown(ci int) {
	si := -1
	for i := len(g) - 1; i >= 0; i-- {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si--
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (g grid) MoveUp() {
	for j := range g[0] {
		g.collapseColUp(j)
		for i := 0; i < len(g)-1; i++ {
			if g[i][j] == g[i+1][j] {
				g[i][j] *= 2
				g[i+1][j] = sen
				i++
			}
		}
		g.collapseColUp(j)
	}
}

func (g grid) collapseColUp(ci int) {
	si := -1
	for i := 0; i < len(g); i++ {
		if g[i][ci] != sen && si >= 0 {
			g[si][ci], g[i][ci] = g[i][ci], g[si][ci]
			si++
		}
		if si == -1 && g[i][ci] == sen {
			si = i
		}
	}
}

func (g grid) String() string {
	var b bytes.Buffer
	for i := range g {
		b.Write([]byte(fmt.Sprintf("%v\n", g[i])))
	}
	return b.String()
}

func main() {
	g := grid{
		[]int{2, 2, 4, 8},
		[]int{2, 4, 4, 4},
		[]int{0, 8, 0, 4},
		[]int{0, 4, 0, 4}}
	fmt.Println(g)
	g.MoveUp()
	fmt.Println(g)
}
