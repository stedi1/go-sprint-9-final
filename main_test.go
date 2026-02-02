package main

import (
	"math/rand/v2"
	"testing"

	"github.com/attic-labs/testify/assert"
)

// Пишите тесты в этом файле

func TestGeneratorWhen0(t *testing.T) {
	testSlice := generateRandomElements(0)
	assert.Empty(t, testSlice)
}

func TestGeneratorWhenOk(t *testing.T) {
	anySize := 300
	goodSlice := generateRandomElements(anySize)
	assert.NotEmpty(t, goodSlice)
	assert.Equal(t, anySize, len(goodSlice))
}

func TestMonoMax(t *testing.T) {
	slice := make([]int, 0)
	for range 100 {
		slice = append(slice, rand.IntN(100))
	}

	// case 1
	slice[32] = 100500
	max := maximum(slice)
	assert.Equal(t, 100500, max)

	// case 2
	slice[82] = 200500
	max = maximum(slice)
	assert.Equal(t, 200500, max)

	// case 3
	slice[32] = 500500
	max = maximum(slice)
	assert.Equal(t, 500500, max)
}

func TestMaxChunks(t *testing.T) {
	slice := make([]int, 0)
	for range 1000 {
		slice = append(slice, rand.IntN(100))
	}

	// case 1
	slice[32] = 100500
	max := maxChunks(slice)
	assert.Equal(t, 100500, max)

	// case 2
	slice[820] = 200500
	max = maxChunks(slice)
	assert.Equal(t, 200500, max)

	// case 3
	slice[32] = 500500
	max = maxChunks(slice)
	assert.Equal(t, 500500, max)

	// case 4
	slice[999] = 6_000_000
	max = maxChunks(slice)
	assert.Equal(t, 6000000, max)
}
