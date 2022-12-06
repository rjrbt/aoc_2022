package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	set "github.com/rjrbt/aoc_2022/internal/set"
)

const (
	values string = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func toSet(input string) set.Set[string] {
	s := make(set.Set[string])
	for _, c := range input {
		s.Add(string(c))
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var items []string
	var keys []string

	i := 0
	l1, l2 := make(set.Set[string]), make(set.Set[string])

	for scanner.Scan() {
		// part 1
		line := scanner.Text()
		l, r := toSet(line[:len(line)/2]), toSet(line[len(line)/2:])
		items = append(items, set.Intersection(l, r).Slice()...)

		// part 2
		f := toSet(line)
		i++

		switch i {
		case 1:
			l1 = f
		case 2:
			l2 = f
		case 3:
			m1 := set.Intersection(l1, l2)
			m2 := set.Intersection(l2, f)
			r := set.Intersection(m1, m2)
			keys = append(keys, r.Slice()...)
			i = 0
		}
	}

	total := 0

	for _, k := range items {
		total += strings.Index(values, k)
	}
	fmt.Println("part 1 total:", total)

	totalPart2 := 0

	for _, k := range keys {
		totalPart2 += strings.Index(values, k)
	}
	fmt.Println("part 2 total:", totalPart2)
}
