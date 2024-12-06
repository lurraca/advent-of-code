package main

import "testing"

func TestDay1Part2(t *testing.T) {

	tests := []struct {
		left, right []int
		expected    int
	}{
		{
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 31,
		},
	}

	t.Run("part 2", func(t *testing.T) {
		for _, test := range tests {
			actual := Distance(test.left, test.right)
			if actual != test.expected {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		}
	})

}
