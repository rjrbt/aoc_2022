package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// This is a quick and dirty solution for part 1 and 2. I created a map of elves
// and exported a slice of the totals, and then reverse sorted the list to solve
// part 1 and 2

func elfMapFromStdIn() (map[int]int, error) {
	elfMap := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	currentElf := 1

	for scanner.Scan() {
		newLine := scanner.Text()
		if newLine == "" {
			currentElf++
		} else {
			v, err := strconv.Atoi(newLine)
			if err != nil {
				return elfMap, err
			}
			elfMap[currentElf] += v
		}
	}
	return elfMap, nil
}

func main() {
	elfMap, err := elfMapFromStdIn()
	if err != nil {
		log.Fatalln(err)
	}

	var t []int

	for _, v := range elfMap {
		t = append(t, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(t)))

	fmt.Println("Answer Part 1:", t[0])
	fmt.Println("Answer Part 2:", t[0]+t[1]+t[2])

}
