package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// initialState returns a slice of string slices initialized
// with the creates in their initial state as outlined in the
// diagram:
// [N]             [R]             [C]
// [T] [J]         [S] [J]         [N]
// [B] [Z]     [H] [M] [Z]         [D]
// [S] [P]     [G] [L] [H] [Z]     [T]
// [Q] [D]     [F] [D] [V] [L] [S] [M]
// [H] [F] [V] [J] [C] [W] [P] [W] [L]
// [G] [S] [H] [Z] [Z] [T] [F] [V] [H]
// [R] [H] [Z] [M] [T] [M] [T] [Q] [W]
//
//	1   2   3   4   5   6   7   8   9
func initialState() [][]string {
	return [][]string{
		{}, // holding the zero index
		{"R", "G", "H", "Q", "S", "B", "T", "N"},
		{"H", "S", "F", "D", "P", "Z", "J"},
		{"Z", "H", "V"},
		{"M", "Z", "J", "F", "G", "H"},
		{"T", "Z", "C", "D", "L", "M", "S", "R"},
		{"M", "T", "W", "V", "H", "Z", "J"},
		{"T", "F", "P", "L", "Z"},
		{"Q", "V", "W", "S"},
		{"W", "H", "L", "M", "T", "D", "N", "C"},
	}
}

// parseMove takes a string like "move 3 from 9 to 7"
// and returns the count, from and to values 3,9,7
func parseMove(s string) (count, from, to int) {
	p := strings.Split(s, " ")
	count, _ = strconv.Atoi(p[1])
	from, _ = strconv.Atoi(p[3])
	to, _ = strconv.Atoi(p[5])
	return
}

func move(c [][]string, count, from, to int) {
	for i := 0; i < count; i++ {
		c[to] = append(c[to], c[from][len(c[from])-1])
		c[from] = c[from][:len(c[from])-1]
	}
}

func movePart2(c [][]string, count, from, to int) {
	c[to] = append(c[to], c[from][len(c[from])-count:]...)
	c[from] = c[from][:len(c[from])-count]
}

func formatOutput(label string, c [][]string) {
	fmt.Print(label, ": ")
	for i := 1; i < 10; i++ {
		fmt.Print(c[i][len(c[i])-1])
	}
	fmt.Print("\n")
}

func main() {
	c := initialState()
	cp2 := initialState()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n := scanner.Text()
		count, from, to := parseMove(n)
		move(c, count, from, to)
		movePart2(cp2, count, from, to)
	}

	formatOutput("Part 1", c)
	formatOutput("Part 2", cp2)
}
