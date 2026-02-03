package main

import (
	"testing"

	"github.com/attic-labs/testify/assert"
)

// Пишите тесты в этом файле
// табилчные тесты?
func TestRandomGenerator(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
		isNil   bool
	}{
		{
			name:    "Пустой слайс",
			size:    0,
			wantLen: 0,
			isNil:   true,
		},
		{
			name:    "Отрицательное число",
			size:    -10,
			wantLen: 0,
			isNil:   true,
		},
		{
			name:    "Малый слайс",
			size:    5,
			wantLen: 5,
			isNil:   false,
		},
		{
			name:    "Средний слайс",
			size:    500,
			wantLen: 500,
			isNil:   false,
		},
		{
			name:    "Большой слайс",
			size:    500_000,
			wantLen: 500_000,
			isNil:   false,
		},
	}

	// тут собственно встраиваем данные из структур для теста
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			randomSlice := generateRandomElements(test.size)

			if test.isNil {
				assert.Nil(t, randomSlice, "generateRandomElements(%d), должен вернуть nil", test.size)
				return
			}

			assert.NotNil(t, randomSlice, "generateRandomElements(%d), должен вернуть слайс 1+", test.size)
			assert.Equal(t, test.wantLen, len(randomSlice), "слайсы не равны", test.size)

		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент слайс",
			data: []int{30},
			want: 30,
		},
		{
			name: "простой слайс",
			data: []int{3, 23, 45, 22, 1, 100, 13},
			want: 100,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := maximum(test.data)
			assert.Equal(t, test.want, result, "maxiumum() ожидание: %d, реальность: %d", test.want, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "пустой слайс",
			data: []int{},
			want: 0,
		},
		{
			name: "один элемент слайс",
			data: []int{30},
			want: 30,
		},
		{
			name: "простой слайс",
			data: []int{3, 23, 45, 22, 1, 100, 13},
			want: 100,
		},
		{
			name: "большой слайс",
			data: []int{3, 23, 45, 22, 1, 100, 13, 100_500, 1234, 23, 85, 500, 500_100, 213_908, 74, 98_761_234, 23_341, 54, 123, 45_123, 734},
			want: 98_761_234,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := maxChunks(test.data)
			assert.Equal(t, test.want, result, "maxChunks() ожидание: %d, реальность: %d", test.want, result)
		})
	}
}

/* func TestGeneratorWhen0(t *testing.T) {
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
*/
