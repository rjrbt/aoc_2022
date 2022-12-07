package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	set "github.com/rjrbt/aoc_2022/internal/set"
)

// genSet takes a number range string such as "22-88" and returns a map of keys from 22 to 88
func genSet(s string) set.Set[int] {
	m := make(set.Set[int])
	r := strings.Split(s, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])
	for i := start; i <= end; i++ {
		m.Add(i)
	}
	return m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	matchCounter := 0
	overlapCounter := 0

	for scanner.Scan() {
		n := scanner.Text()
		setStrings := strings.Split(n, ",")
		a, b := genSet(setStrings[0]), genSet(setStrings[1])

		if a.Contains(b) {
			matchCounter++
		} else if b.Contains(a) {
			matchCounter++
		}
		if a.Intersects(b) {
			overlapCounter++
		}
	}
	fmt.Println("contained:", matchCounter)
	fmt.Println("overlap:", overlapCounter)
}
