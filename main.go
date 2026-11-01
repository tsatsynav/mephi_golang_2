package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует случайные элементы.
func generateRandomElements(size int) []int {
	// Если размер 0 или отрицательный, возвращаем пустой слайс
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) // случайное положительное число
	}
	return data
}

// maximum возвращает максимальное значение элемента.
func maximum(data []int) int {
	// Обрабатываем крайний случай: пустой слайс
	if len(data) == 0 {
		return 0
	}

	maxVal := data[0]
	for _, v := range data[1:] {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// maxChunks возвращает максимальное значение элемента в частях.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	// Если элементов меньше, чем частей, ищем в один поток
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		// Для последнего чанка берем остаток элементов, если длина не делится нацело
		if i == CHUNKS-1 {
			end = len(data)
		}

		// Запускаем горутину для поиска максимума в куске
		go func(chunk []int, idx int) {
			defer wg.Done()
			maxes[idx] = maximum(chunk)
		}(data[start:end], i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Находим максимум среди 8 локальных максимумов
	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
