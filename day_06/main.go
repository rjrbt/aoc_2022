package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rjrbt/aoc_2022/internal/set"
)

func bump(state []byte, b byte) []byte {
	l := len(state)
	for i := 0; i < l-1; i++ {
		state[i] = state[i+1]
	}
	state[l-1] = b
	return state
}

func makerPosition(l int, b []byte) (r int) {
	state := make([]byte, l)
	copy(state, b[:l])

	for i := l; i < len(b); i++ {
		stateSet := set.FromSlice(state)
		if len(stateSet) < l {
			state = bump(state, b[i])
		} else {
			return i
		}
	}
	return r
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	b := []byte(s)

	// part 1
	fmt.Println("packet position: ", makerPosition(4, b))

	// part 2
	fmt.Println("message position:", makerPosition(14, b))

}
