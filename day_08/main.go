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

func (t tree) viewScore() int {
	u := t.viewCount(up)
	d := t.viewCount(down)
	l := t.viewCount(left)
	r := t.viewCount(right)
	return u * d * l * r
}

func (t tree) viewCount(o ordinal) int {
	count := 0
	switch o {
	case up:
		if t.up == nil {
			return count
		}
		return counter(t.up, o, count+1)
	case down:
		if t.down == nil {
			return count
		}
		return counter(t.down, o, count+1)
	case left:
		if t.left == nil {
			return count
		}
		return counter(t.left, o, count+1)
	case right:
		if t.right == nil {
			return count
		}
		return counter(t.right, o, count+1)
	}
	return count
}

func counter(t *tree, o ordinal, count int) int {
	switch o {
	case up:
		if t.up == nil {
			return count
		}
		if t.up.height >= t.height {
			return counter(t.up, o, count+1)
		}
	case down:
		if t.down == nil {
			return count
		}
		if t.down.height >= t.height {
			return counter(t.down, o, count+1)
		}
	case left:
		if t.left == nil {
			return count
		}
		if t.left.height >= t.height {
			return counter(t.left, o, count+1)
		}
	case right:
		if t.right == nil {
			return count
		}
		if t.right.height >= t.height {
			return counter(t.right, o, count+1)
		}
	}
	return count
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
	topViewScore := 0

	for _, v := range f.trees {
		// part 1
		if v.isVisible() {
			visibleCounter++
		}
		// part 2
		score := v.viewScore()
		if score > topViewScore {
			topViewScore = score
		}
	}
	fmt.Println("part 1:", visibleCounter)
	// note: currently the score is too low, I need to fix some logic
	fmt.Println("part 2:", topViewScore)
}
