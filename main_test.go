package main

import (
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	if len(generateRandomElements(0)) != 0 {
		t.Error("для размера 0 ожидался пустой слайс")
	}
	if len(generateRandomElements(-5)) != 0 {
		t.Error("для отрицательного размера ожидался пустой слайс")
	}
	if len(generateRandomElements(50)) != 50 {
		t.Error("ожидался слайс длиной 50")
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"обычный слайс", []int{1, 5, 3, 9, 2}, 9},
		{"отрицательные числа", []int{-10, -5, -20}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"пустой слайс", []int{}, 0},
		{"меньше 8 элементов", []int{1, 5, 3}, 5},
		{"больше 8 элементов", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 99, 14, 15, 100}, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}