package gomisclib_test

import (
	"testing"

	gomisclib "github.com/OZoneGuy/go-misc-lib"
	"github.com/stretchr/testify/assert"
)

func MapTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{2, 4, 6}, {8, 10, 12}}

	for i, input := range inputs {
		output := gomisclib.Map(func(x int) int { return x * 2 }, input)
		assert.Equal(t, outputs[i], output, "Map(%v) = %v, want %v", input, output, outputs[i])
	}
}

func MutMapTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{2, 4, 6}, {8, 10, 12}}

	for i, input := range inputs {
		gomisclib.MutMap(func(x int) int { return x * 2 }, input)
		assert.Equal(t, outputs[i], input, "MutMap(%v) = %v, want %v", input, input, outputs[i])
	}
}

func FilterTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := [][]int{{1, 3}, {5}}

	for i, input := range inputs {
		output := gomisclib.Filter(func(x int) bool { return x%2 == 1 }, input)
		assert.Equal(t, outputs[i], output, "Filter(%v) = %v, want %v", input, output, outputs[i])
	}
}

func ContainsTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.Contains(input, 2)
		assert.Equal(t, outputs[i], output, "Contains(%v) = %v, want %v", input, output, outputs[i])
	}
}

func AllTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.All(func(x int) bool { return x < 4 }, input)
		assert.Equal(t, outputs[i], output, "All(%v) = %v, want %v", input, output, outputs[i])
	}
}

func AnyTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []bool{true, false}

	for i, input := range inputs {
		output := gomisclib.Any(func(x int) bool { return x < 4 }, input)
		assert.Equal(t, outputs[i], output, "Any(%v) = %v, want %v", input, output, outputs[i])
	}
}

func FlattenTest(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}}
	expected := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, gomisclib.Flatten(input), "Should flatten the list of lists")
}

func ApplyNTest(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{3, 4, 5}
	f := func(x int) int { return x + 1 }
	for i, n := range input {
		assert.Equal(t, expected[i], gomisclib.ApplyN(f, n, 2), "Should apply f n times")
	}
}

func FoldLTest(t *testing.T) {
	inputs := [][]int{{1, 2, 3}, {4, 5, 6}}
	outputs := []int{6, 15}
	f := func(x, y int) int { return x + y }
	for i, ns := range inputs {
		assert.Equal(t, outputs[i], gomisclib.FoldL(f, ns, 0), "Should fold left")
	}
}

func FoldRTest(t *testing.T) {
	inputs := [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}}
	outputs := []string{"cba", "fed"}
	f := func(x rune, y string) string { return string(x) + y }
	for i, cs := range inputs {
		assert.Equal(t, outputs[i], gomisclib.FoldR(f, cs, ""), "Should fold right")
	}
}
