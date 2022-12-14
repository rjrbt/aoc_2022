package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type session struct {
	root    *dir
	current *dir
}

func (s *session) cd(d string) {
	switch d {
	case "/":
		s.current = s.root
	case "..":
		s.current = s.current.parent
	default:
		s.current = s.current.children[d]
	}
}

type dir struct {
	name     string
	parent   *dir
	children map[string]*dir
	files    map[string]int
}

func (d *dir) addFile(size int, name string) {
	d.files[name] = size
}

func (d *dir) addDir(name string, parent *dir) {
	if _, ok := d.children[name]; ok {
		return
	}
	d.children[name] = newDir(name, parent)
}

func (d *dir) total() int {
	total := 0
	for _, size := range d.files {
		total += size
	}
	return total
}

func (d *dir) recursiveTotal() int {
	total := 0
	for _, v := range d.children {
		total += v.recursiveTotal()
	}
	total += d.total()
	return total
}

func newDir(name string, parent *dir) *dir {
	return &dir{
		name:     name,
		parent:   parent,
		children: make(map[string]*dir),
		files:    make(map[string]int),
	}
}

func newSession() *session {
	rootDir := newDir("/", nil)
	return &session{
		root:    rootDir,
		current: rootDir,
	}
}

func process(n string, s *session) {
	tokens := strings.Split(n, " ")
	switch tokens[0] {
	case "$":
		if tokens[1] == "cd" {
			s.cd(tokens[2])
		}
	case "dir":
		s.current.addDir(tokens[1], s.current)
	default: // add file
		size, _ := strconv.Atoi(tokens[0])
		s.current.addFile(size, tokens[1])
	}
}

func main() {
	s := newSession()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		process(scanner.Text(), s)
	}

	maxDirSize := 100000
	// TODO: make a channel queue for all dirs

}
