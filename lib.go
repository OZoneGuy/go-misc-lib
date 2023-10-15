package gomisclib

// Map applies f to each element of a, returning a new slice containing the results.
func Map[T1, T2 any](f func(T1) T2, a []T1) []T2 {
	b := make([]T2, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}

// Equal returns true if a and b have the same length and contain the same elements in the same order.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

// MutMap applies f to each element of a mutating a
func MutMap[T any](f func(T) T, a []T) {
	for i := range a {
		a[i] = f(a[i])
	}
}

// Filter returns a slice containing all elements of a for which f returns true.
func Filter[T any](f func(T) bool, a []T) []T {
	b := make([]T, 0, len(a))
	for _, v := range a {
		if f(v) {
			b = append(b, v)
		}
	}
	return b
}

// Contains returns true if v is in a.
func Contains[T comparable](a []T, v T) bool {
	for _, x := range a {
		if x == v {
			return true
		}
	}
	return false
}

// All returns true if f returns true for all elements of a.
func All[T any](f func(T) bool, a []T) bool {
	for _, v := range a {
		if !f(v) {
			return false
		}
	}
	return true
}

// Any returns true if f returns true for any element of a.
func Any[T any](f func(T) bool, a []T) bool {
	for _, v := range a {
		if f(v) {
			return true
		}
	}
	return false
}

// Count returns the number of elements of a for which f returns true.
func Count[T any](f func(T) bool, a []T) int {
	n := 0
	for _, v := range a {
		if f(v) {
			n++
		}
	}
	return n
}

// Find returns the first element of a for which f returns true, or nil if there is no such element.
func Find[T any](f func(T) bool, a []T) *T {
	for _, v := range a {
		if f(v) {
			return &v
		}
	}
	return nil
}

// FindIndex returns the index of the first element of a for which f returns true, or -1 if there is no such element.
func FindIndex[T any](f func(T) bool, a []T) int {
	for i, v := range a {
		if f(v) {
			return i
		}
	}
	return -1
}

// ForEeach calls f on each element of a.
func ForEach[T any](f func(T), a []T) {
	for _, v := range a {
		f(v)
	}
}

// Flatten flattens a list of lists of A to a list of A
func Flatten[T any](a [][]T) []T {
	var out []T
	for _, i := range a {
		out = append(out, i...)
	}
	return out
}

// ApplyN applies f to x N times, returning the result.
func ApplyN[T any](f func(T) T, x T, n int) T {
	for i := 0; i < n; i++ {
		x = f(x)
	}
	return x
}

// FoldL folds a from the left with f, returning the result.
func FoldL[T1, T2 any](f func(T1, T2) T1, a []T2, init T1) T1 {
	for _, v := range a {
		init = f(init, v)
	}
	return init
}

// FoldR folds a from the right with f, returning the result.
func FoldR[T1, T2 any](f func(T2, T1) T1, a []T2, init T1) T1 {
	for i := len(a) - 1; i >= 0; i-- {
		init = f(a[i], init)
	}
	return init
}
