package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {
	assert.EqualValues(t, []int{1, 2, 3}, Ints("[1, 2, 3]"), "")
	assert.EqualValues(t, []int{-1, 2, 3}, Ints("[-1, 2, 3]"), "")
}

func TestInts2d(t *testing.T) {
	assert.EqualValues(t, [][]int{{1, 2}, {3}}, Ints2d("[[1, 2], [3]]"), "")
	assert.EqualValues(t, [][]int{{1, -2}, {3}}, Ints2d("[[1, -2], [3]]"), "")
	assert.EqualValues(t, [][]int{{}, {}}, Ints2d("[[], []]"), "")
}

func TestFloats(t *testing.T) {
	assert.EqualValues(t, []float64{1.0, 2.0, 3.0}, Floats("[1.00, 2.0, 3.0]"), "")
	assert.EqualValues(t, []float64{1.0, -2.0, 3.0}, Floats("[1.00, -2.0, 3.0]"), "")
}

func TestFloats2d(t *testing.T) {
	assert.EqualValues(t, [][]float64{{1.0, 2.0}, {3.0}}, Floats2d("[[1.0, 2], [3.00]]"), "")
	assert.EqualValues(t, [][]float64{{1.0, -2.0}, {3.0}}, Floats2d("[[1.0, -2.0], [3.00]]"), "")
}

func TestStrings(t *testing.T) {
	assert.EqualValues(t, []string{"a", "bb", "ccc"}, Strings(`["a", "bb", "ccc"]`), "")
}

func TestStrings2d(t *testing.T) {
	assert.EqualValues(t, [][]string{{"a", "bb"}, {"ccc"}}, Strings2d(`[["a", "bb"], ["ccc"]]`), "")
}
