package main

import "testing"

func TestDay1Part1(t *testing.T) {

	tests := []struct {
		left, right []int
		expected    int
	}{
		{
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
	}

	t.Run("part 1", func(t *testing.T) {
		for _, test := range tests {
			actual := Distance(test.right, test.left)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		}
	})

}
