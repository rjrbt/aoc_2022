package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type set map[int]struct{}

func (s set) add(x int) {
	s[x] = struct{}{}
}

// contains compares a set to another set and returns true is a is contained by the set
// otherwise it returns false
func (s set) contains(a set) bool {
	for k := range s {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

func (s set) intersects(a set) bool {
	for k := range s {
		if _, ok := a[k]; ok {
			return true
		}
	}
	return false
}

// genSet takes a number range string such as "22-88" and returns a map of keys from 22 to 88
func genSet(s string) set {
	m := make(set)
	r := strings.Split(s, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])
	for i := start; i <= end; i++ {
		m.add(i)
	}
	return m
}

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	matchCounter := 0
	overlapCounter := 0

	for scanner.Scan() {
		newLine := scanner.Text()
		setStrings := strings.Split(newLine, ",")
		a, b := genSet(setStrings[0]), genSet(setStrings[1])

		if a.contains(b) {
			matchCounter++
		} else if b.contains(a) {
			matchCounter++
		}
		if a.intersects(b) {
			overlapCounter++
		}
	}
	fmt.Println("contained:", matchCounter)
	fmt.Println("overlap:", overlapCounter)
}
