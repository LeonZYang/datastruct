package bisect_test

import (
	"testing"

	"github.com/LeonZYang/datastruct/bisect"
	"github.com/stretchr/testify/assert"
)

type IntArr []int

func (ir IntArr) Equal(i int, val int) int {
	if ir[i] == val {
		return 0
	}

	if ir[i] > val {
		return 1
	}
	return -1
}

func (ir IntArr) Len() int {
	return len(ir)
}

func TestBiSect(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{
		1, 1, 1, 2, 2, 3, 4, 5, 6, 7,
	}

	assert.Equal(-1, bisect.Bisect(arr, 0))

	assert.Equal(2, bisect.Bisect(arr, 1))

	assert.Equal(4, bisect.Bisect(arr, 2))

	assert.Equal(9, bisect.Bisect(arr, 7))

	assert.Equal(-1, bisect.Bisect(arr, 10))
}

func TestBiSectLeft(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{
		1, 1, 1, 2, 2, 3, 4, 5, 6, 7,
	}

	assert.Equal(-1, bisect.BisectLeft(arr, 0))

	assert.Equal(0, bisect.BisectLeft(arr, 1))

	assert.Equal(3, bisect.BisectLeft(arr, 2))

	assert.Equal(9, bisect.BisectLeft(arr, 7))

	assert.Equal(-1, bisect.BisectLeft(arr, 10))
}

func TestBiSectRight(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{
		1, 1, 1, 2, 2, 3, 4, 5, 6, 7,
	}

	assert.Equal(-1, bisect.BisectRight(arr, 0))

	assert.Equal(3, bisect.BisectRight(arr, 1))

	assert.Equal(5, bisect.BisectRight(arr, 2))

	assert.Equal(10, bisect.BisectRight(arr, 7))

	assert.Equal(-1, bisect.BisectRight(arr, 10))
}
