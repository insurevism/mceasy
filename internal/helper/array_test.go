package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayChunk(t *testing.T) {

	{
		origin := []int{1}
		assert.Equal(t, ArrayChunk(origin, 2), [][]int{{1}})
		assert.Equal(t, origin, []int{1})
	}
	{
		origin := []int{1, 2}
		assert.Equal(t, ArrayChunk(origin, 2), [][]int{{1, 2}})
		assert.Equal(t, origin, []int{1, 2})
	}
	{
		origin := []int{1, 2, 3, 4, 5}
		assert.Equal(t, ArrayChunk(origin, 2), [][]int{{1, 2}, {3, 4}, {5}})
		assert.Equal(t, origin, []int{1, 2, 3, 4, 5})
	}
	{
		origin := []int{1, 2, 3, 4, 5}
		assert.Equal(t, ArrayChunk(origin, 3), [][]int{{1, 2, 3}, {4, 5}})
		assert.Equal(t, origin, []int{1, 2, 3, 4, 5})
	}

}
