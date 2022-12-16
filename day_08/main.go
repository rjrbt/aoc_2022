package main

import (
	"bufio"
	"fmt"
	"os"
)

type ordinal int

const (
	up ordinal = iota
	down
	left
	right
)

type tree struct {
	id     string
	height int
	up     *tree
	down   *tree
	left   *tree
	right  *tree
}

func (t tree) isVisible() bool {
	// handle case of edge tree
	if t.up == nil || t.down == nil || t.left == nil || t.right == nil {
		return true
	}
	// loop through all ordinal directions and if any are shorter that
	// the current tree height, return true
	for _, o := range []ordinal{up, down, left, right} {
		if t.max(o, 0) < t.height {
			return true
		}
	}
	return false
}

func (t tree) max(o ordinal, localMax int) int {
	switch o {
	case up:
		if t.up != nil {
			if t.up.height > localMax {
				localMax = t.up.height
			}
			return t.up.max(up, localMax)
		}
	case down:
		if t.down != nil {
			if t.down.height > localMax {
				localMax = t.down.height
			}
			return t.down.max(down, localMax)
		}
	case left:
		if t.left != nil {
			if t.left.height > localMax {
				localMax = t.left.height
			}
			return t.left.max(left, localMax)
		}
	case right:
		if t.right != nil {
			if t.right.height > localMax {
				localMax = t.right.height
			}
			return t.right.max(right, localMax)
		}
	}
	return localMax
}

func addForestRow(line string, row int) {

}

type forest struct {
	trees map[string]*tree
}

func appendRow(f *forest, line string, row int) {
	// add all trees for the row
	width := len(line) - 1
	for col, char := range line {
		id := fmt.Sprintf("%v,%v", row, col)
		f.trees[id] = &tree{id: id, height: int(char - '0')}
	}
	// wire up neighbors of the trees
	for col := 0; col <= width; col++ {
		id := fmt.Sprintf("%v,%v", row, col)
		l := fmt.Sprintf("%v,%v", row, col-1)
		r := fmt.Sprintf("%v,%v", row, col+1)
		t := f.trees[id]
		t.left = f.trees[l]
		t.right = f.trees[r]

		if row != 0 {
			u := fmt.Sprintf("%v,%v", row-1, col)
			t.up = f.trees[u]
			t.up.down = t
		}
	}
}

func newForest() *forest {
	return &forest{
		trees: make(map[string]*tree),
	}
}

func main() {
	f := newForest()

	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	for scanner.Scan() {
		appendRow(f, scanner.Text(), row)
		row++
	}

	visibleCounter := 0

	for _, v := range f.trees {
		if v.isVisible() {
			visibleCounter++
		}
	}
	fmt.Println("part 1:", visibleCounter)
}
