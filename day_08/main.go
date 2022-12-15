package main

type ordinal int

const (
	up ordinal = iota
	down
	left
	right
)

type tree struct {
	height int
	up     *tree
	down   *tree
	left   *tree
	right  *tree
}

func (t tree) isVisible() bool {
	// handle case of edge
	if t.up == nil || t.down == nil || t.left == nil || t.right == nil {
		return true
	}
	if t.max(up, 0) < t.height {
		return true
	}
	if t.max(down, 0) < t.height {
		return true
	}
	if t.max(left, 0) < t.height {
		return true
	}
	if t.max(right, 0) < t.height {
		return true
	}
	return false
}

func (t tree) max(direction ordinal, val int) int {
	localMax := 0
	if t.height > val {
		localMax = t.height
	}

	switch direction {
	case up:
		if t.up != nil {
			return t.up.max(up, localMax)
		}
	case down:
		if t.down != nil {
			return t.down.max(down, localMax)
		}
	case left:
		if t.left != nil {
			return t.left.max(left, localMax)
		}
	case right:
		if t.right != nil {
			return t.right.max(right, localMax)
		}
	}
	return localMax
}

func main() {

}
