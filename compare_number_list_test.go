package compare_test

import (
	"compare"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainerContains(t *testing.T) {

	for _, tt := range []struct {
		numbers   []int
		candidate int
		expected  bool
	}{
		{
			numbers:   []int{0, 3, 5, 234, 53},
			candidate: 5,
			expected:  true,
		},
		{
			numbers:   []int{0, 3, 5, 234, 53},
			candidate: 7,
			expected:  false,
		},
		{
			numbers:   []int{1, 15, 20, 45, 60},
			candidate: 60,
			expected:  true,
		},
		{
			numbers:   []int{40, 200, 2000},
			candidate: 201,
			expected:  false,
		},
	} {
		c := compare.NewContainer(tt.numbers)
		assert.Equal(t, tt.expected, c.Contains(tt.candidate))
	}

}
func BenchmarkContainerContains(b *testing.B) {
	rand.Seed(1)
	randomNumbers := rand.Perm(1000000)
	cBig := compare.NewContainer(randomNumbers)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cBig.Contains(i)
	}
}
