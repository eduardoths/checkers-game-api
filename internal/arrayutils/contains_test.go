package arrayutils_test

import (
	"testing"

	"github.com/eduardoths/checkers-game/internal/arrayutils"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		array []T
		value T
		want  bool
	}
	testCases := map[string]testCase[int]{
		"Should return false if int can't be found": {
			array: []int{0, 1, 2, 3, 4, 5},
			value: 100,
			want:  false,
		},
		"Should return true if int is found": {
			array: []int{0, 1, 2, 3, 4, 5},
			value: 0,
			want:  true,
		},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			actual := arrayutils.Contains(tc.array, tc.value)
			assert.Equal(t, tc.want, actual)
		})
	}
}
