package main

import (
	"bytes"
	"fmt"
)

type grid [][]int

func (g grid) moveRight() {
	for i := range g {
		// combine
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == g[i][j+1] {
				g[i][j] = 0
				g[i][j+1] *= 2
				j++
			} else if g[i][j+1] == 0 {
				g[i][j], g[i][j+1] = g[i][j+1], g[i][j]
			}
		}
		// collapse
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == 0 {
				g[i][j], g[i][j-1] = g[i][j-1], g[i][j]
			}
		}
	}
}

func (g grid) moveLeft() {
	for i := range g {
		// combine
		for j := len(g[i]) - 1; j > 0; j-- {
			if g[i][j] == g[i][j-1] {
				g[i][j] = 0
				g[i][j-1] *= 2
				j--
			} else if g[i][j-1] == 0 {
				g[i][j], g[i][j-1] = g[i][j-1], g[i][j]
			}
		}
		// collapse
		for j := 0; j < len(g[i])-1; j++ {
			if g[i][j] == 0 {
				g[i][j], g[i][j+1] = g[i][j+1], g[i][j]
			}
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
	g := grid{[]int{2, 2, 4, 8}, []int{2, 4, 4, 4}, []int{0, 8, 0, 2}, []int{0, 4, 0, 4}}
	fmt.Println(g)
	g.moveLeft()
	fmt.Println(g)
}
