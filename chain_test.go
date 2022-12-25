package chain

import (
	"reflect"
	"testing"
)

func TestChain_Bounds(t *testing.T) {

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	expected := []string{
		"abcd",
		"efgh",
		"ijkl",
		"mnop",
		"qrst",
		"uvwx",
		"yz",
	}

	actual := make([]string, 0, len(expected))

	c := New(len(alphabet), OptionStep(4))
	for c.Next() {
		left, right := c.Bounds()
		segment := alphabet[left:right]
		actual = append(actual, segment)
	}

	assertEqual(t, expected, actual)
}

func TestChain_Next(t *testing.T) {

	testCases := []struct {
		start    int
		step     int
		stop     int
		expected [][]int
	}{
		{
			start:    0,
			step:     1,
			stop:     5,
			expected: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}},
		},
		{
			start:    0,
			step:     2,
			stop:     5,
			expected: [][]int{{0, 2}, {2, 4}, {4, 5}},
		},
		{
			start:    0,
			step:     3,
			stop:     5,
			expected: [][]int{{0, 3}, {3, 5}},
		},
		{
			start:    0,
			step:     4,
			stop:     5,
			expected: [][]int{{0, 4}, {4, 5}},
		},
		{
			start:    0,
			step:     5,
			stop:     5,
			expected: [][]int{{0, 5}},
		},
	}

	for _, testCase := range testCases {
		c := New(testCase.stop, OptionStart(testCase.start), OptionStep(testCase.step))
		actual := make([][]int, 0)
		for c.Next() {
			left, right := c.Bounds()
			actual = append(actual, []int{left, right})
		}

		assertEqual(t, testCase.expected, actual)
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
