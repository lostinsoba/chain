package chain

import (
	"reflect"
	"testing"
)

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
		var c Chain
		c.SetStart(testCase.start)
		c.SetStep(testCase.step)
		c.SetStop(testCase.stop)
		actual := make([][]int, 0)
		for c.Next() {
			left, right := c.Bounds()
			actual = append(actual, []int{left, right})
		}

		assertEqual(t, t.Name(), testCase.expected, actual)
	}
}

func TestChain_Bounds(t *testing.T) {
	var (
		alphabet = "abcdefghijklmnopqrstuvwxyz"
		step     = 4
	)
	testCases := []struct {
		name     string
		loops    []int
		expected []string
	}{
		{
			name:  "forward",
			loops: []int{1},
			expected: []string{
				"abcd",
				"efgh",
				"ijkl",
				"mnop",
				"qrst",
				"uvwx",
				"yz",
			},
		},
		{
			name:  "backward",
			loops: []int{-1},
			expected: []string{
				"wxyz",
				"stuv",
				"opqr",
				"klmn",
				"ghij",
				"cdef",
				"ab",
			},
		},
		{
			name:  "forward then backward",
			loops: []int{1, -1},
			expected: []string{
				"abcd",
				"efgh",
				"ijkl",
				"mnop",
				"qrst",
				"uvwx",
				"yz",
				"wxyz",
				"stuv",
				"opqr",
				"klmn",
				"ghij",
				"cdef",
				"ab",
			},
		},
		{
			name:  "backward then forward",
			loops: []int{-1, 1},
			expected: []string{
				"wxyz",
				"stuv",
				"opqr",
				"klmn",
				"ghij",
				"cdef",
				"ab",
				"abcd",
				"efgh",
				"ijkl",
				"mnop",
				"qrst",
				"uvwx",
				"yz",
			},
		},
	}

	for _, testCase := range testCases {
		var c Chain
		c.SetStop(len(alphabet))
		c.SetStep(step)

		actual := make([]string, 0, len(testCase.expected))

		var (
			prevDirection int
		)
		for _, direction := range testCase.loops {
			if direction == -1 || direction == -prevDirection {
				c.Reverse()
			}
			prevDirection = direction

			for c.Next() {
				left, right := c.Bounds()
				segment := alphabet[left:right]
				actual = append(actual, segment)
			}
		}

		assertEqual(t, testCase.name, testCase.expected, actual)
	}
}

func TestChain_Reset(t *testing.T) {
	var (
		stop    = 15
		step    = 5
		resetAt = 10
	)
	testCases := []struct {
		name     string
		loops    []int
		expected [][]int
	}{
		{
			name:     "forward",
			loops:    []int{1},
			expected: [][]int{{0, 5}, {5, 10}, {10, 15}, {0, 5}, {5, 10}, {10, 15}},
		},
		{
			name:     "backward",
			loops:    []int{-1},
			expected: [][]int{{10, 15}, {10, 15}, {5, 10}, {0, 5}},
		},
	}

	for _, testCase := range testCases {
		var c Chain
		c.SetStop(stop)
		c.SetStep(step)

		actual := make([][]int, 0, len(testCase.expected))

		var (
			prevDirection int
			alreadyReset  bool
		)
		for _, direction := range testCase.loops {
			if direction == -1 || direction == -prevDirection {
				c.Reverse()
			}
			prevDirection = direction

			for c.Next() {
				left, right := c.Bounds()
				subInterval := []int{left, right}
				actual = append(actual, subInterval)
				if left == resetAt && !alreadyReset {
					alreadyReset = true
					c.Reset()
				}
			}
		}

		assertEqual(t, testCase.name, testCase.expected, actual)
	}
}

func assertEqual(t *testing.T, name string, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("[%s] expected: %v, actual: %v", name, expected, actual)
	}
}
