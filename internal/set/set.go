package set

import "fmt"

// Set is a simple key-only map that conforms to the Hashable interface
type Set[T comparable] map[T]struct{}

// Add adds an item of type T to the Set[T]
func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

// Has takes a value and returns a boolean based on
// whether the item is present or not
func (s Set[T]) Has(x T) bool {
	_, ok := s[x]
	return ok
}

// Contains takes a set and returns a boolean value if the
// set is completely contained inside of the compared set.
func (s Set[T]) Contains(a Set[T]) bool {
	for k := range s {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

// Intersects compares a Set[T] with itself to see if there is any intersection
// it returns true on any intersection otherwise it returns false
func (s Set[T]) Intersects(a Set[T]) bool {
	for k := range s {
		if _, ok := a[k]; ok {
			return true
		}
	}
	return false
}

// Intersection takes two Sets and returns a set of the
// intersection between the two.
func Intersection[T comparable](a, b Set[T]) Set[T] {
	s := make(Set[T])
	for k := range a {
		if b.Has(k) {
			s.Add(k)
		}
	}
	return s
}

// Merge takes two Set[T] and merges them together into a new Set[T]
func Merge[T comparable](a, b Set[T]) Set[T] {
	s := make(Set[T])
	for k := range a {
		s.Add(k)
	}
	for k := range b {
		s.Add(k)
	}
	return s
}

// Slice returns a Slice T of Set T
func (s Set[T]) Slice() []T {
	var r []T
	for k := range s {
		r = append(r, k)
	}
	return r
}

// String returns a string representation of a Set, similar to map[] but
// without the values
func (s Set[T]) String() string {
	var r string
	r += "Set["
	for k := range s {
		r += fmt.Sprint(k, " ")
	}
	r = r[:len(r)-1]
	r += "]"
	return r
}
