package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tallyMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func tally(a, b int) int {
	r := b - a
	switch r {
	case 0: // handles a tie
		return b + 3
	case -2, 1: // handles winning
		return b + 6
	default: // handles losses
		return b
	}
}

// win, lose, or draw
func wld(a, b int) int {
	switch b {
	case 1: // lose
		return ((a + 4) % 3) + 1
	case 2: // draw
		return a + 3
	default: // win
		return ((a + 3) % 3) + 1 + 6
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalScore := 0
	totalPitch := 0
	for scanner.Scan() {
		newLine := scanner.Text()
		v := strings.Split(newLine, " ")
		l, r := tallyMap[v[0]], tallyMap[v[1]]
		totalScore += tally(l, r)
		totalPitch += wld(l, r)
	}
	fmt.Println("total score:", totalScore)
	fmt.Println("total pitch:", totalPitch)
}
