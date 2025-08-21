package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

// ==== Simple Test

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d; Expected: %d", result, expected)
	}
}

// ===== Multiple tests by table driven

func TestAddTableDriven(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d; Expected: %d", test.a, test.b, result, test.expected)
		}
	}
}

// ==== Multiple tests by subtests

func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("result: %d; expected: %d", result, test.expected)
			}
		})
	}
}

// ===== Benchmarking

func BenchmarkAddSmallInput(b *testing.B) {
	for range b.N {
		Add(2, 3)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	for range b.N {
		Add(200, 300)
	}
}

func BenchmarkAddLargeInput(b *testing.B) {
	for range b.N {
		Add(20000, 30000)
	}
}

// ===== Profiling

func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)

	if len(slice) != size {
		t.Errorf("Expected slice size %d, received %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer()

	for range b.N {
		SumSlice(slice)
	}
}
