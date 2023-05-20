package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {

	assert.EqualValues(t, []int{1, 2, 3}, Ints("[1, 2, 3]"), "")

}

func TestInts2d(t *testing.T) {

	assert.EqualValues(t, [][]int{[]int{1, 2}, []int{3}}, Ints2d("[[1, 2], [3]]"), "")

}
