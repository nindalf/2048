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

func (g grid) collapseRowRight(rowno int) {
	row := g[rowno]
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

func (g grid) collapseRowLeft(rowno int) {
	row := g[rowno]
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
	g.MoveLeft()
	fmt.Println(g)
}
