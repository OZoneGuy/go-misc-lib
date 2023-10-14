package gomisclib_test

import (
	"testing"

	gomisclib "github.com/OZoneGuy/go-misc-lib"
)

func MapTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{2, 4, 6}, {8, 10, 12}}

	for i, input := range inputs {
		output := gomisclib.Map(func(x int) int { return x * 2 }, input)
		if gomisclib.Equal(output, outputs[i]) {
			t.Errorf("Map(%v) = %v, want %v", input, output, outputs[i])
		}
	}
}

func MutMapTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{2, 4, 6}, {8, 10, 12}}

	for i, input := range inputs {
		gomisclib.MutMap(func(x *int) { *x *= 2 }, input)
		if gomisclib.Equal(input, outputs[i]) {
			t.Errorf("MutMap(%v) = %v, want %v", input, input, outputs[i])
		}
	}
}

func FilterTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{1, 3}, {5}}

	for i, input := range inputs {
		output := gomisclib.Filter(func(x int) bool { return x%2 == 1 }, input)
		if gomisclib.Equal(output, outputs[i]) {
			t.Errorf("Filter(%v) = %v, want %v", input, output, outputs[i])
		}
	}
}

func ContainsTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.Contains(input, 2)
		if output != outputs[i] {
			t.Errorf("Contains(%v) = %v, want %v", input, output, outputs[i])
		}
	}
}

func AllTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.All(func(x int) bool { return x < 4 }, input)
		if output != outputs[i] {
			t.Errorf("All(%v) = %v, want %v", input, output, outputs[i])
		}
	}
}

func AnyTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.Any(func(x int) bool { return x < 4 }, input)
		if output != outputs[i] {
			t.Errorf("Any(%v) = %v, want %v", input, output, outputs[i])
		}
	}
}
