package set

import "fmt"

// Hashable are the supported types for the set
type Hashable interface {
	~int | ~string | ~uint64 | ~byte
}

// Set is a simple key-only map that conforms to the Hashable interface
type Set[T Hashable] map[T]struct{}

func (s Set[T]) Add(x T) {
	s[x] = struct{}{}
}

// Has takes a value and returns a boolean based on
// whether the item is present or not
func (s Set[T]) Has(x T) bool {
	if _, ok := s[x]; ok {
		return true
	}
	return false
}

func (s Set[T]) Contains(a Set[T]) bool {
	for k := range s {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}

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
func Intersection[T Hashable](a, b Set[T]) Set[T] {
	s := make(Set[T])
	for k := range a {
		if b.Has(k) {
			s.Add(k)
		}
	}
	return s
}

func Merge[T Hashable](a, b Set[T]) Set[T] {
	s := make(Set[T])
	for k := range a {
		s.Add(k)
	}
	for k := range b {
		s.Add(k)
	}
	return s
}

func (s Set[T]) Slice() []T {
	var r []T
	for k := range s {
		r = append(r, k)
	}
	return r
}

func (s Set[T]) String() string {
	var r string
	r += "Set[ "
	for k := range s {
		r += fmt.Sprint(k, " ")
	}
	r += "]"
	return r
}
