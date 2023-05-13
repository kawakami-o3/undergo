package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInts(t *testing.T) {

	assert.EqualValues(t, []int{1, 2, 3}, Ints("[1, 2, 3]"), "")

}
