package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	// создаем массив для записи данных
	data := make([]int, 0, size)
	for range size {
		data = append(data, rand.Int())
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) <= 0 {
		return 0
	}

	return slices.Max(data)
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) <= 0 {
		return 0
	}

	var wg sync.WaitGroup

	// слайс для сохранение результатов по чанку
	chunkResult := make([]int, CHUNKS)

	for i := range CHUNKS {
		step := len(data) / CHUNKS
		startIndex := i * step
		endIndex := startIndex + step
		if i == CHUNKS-1 {
			endIndex = len(data)
		}

		wg.Add(1)
		go func(index, start, end int) {
			defer wg.Done()
			chunkResult[index] = maximum(data[start:end])
		}(i, startIndex, endIndex)
	}
	wg.Wait()
	return maximum(chunkResult)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
