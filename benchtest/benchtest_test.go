package benchtest

import (
	"fmt"
	"testing"
)

// func GenerateRandomSlice(size int) []int {
// 	slice := make([]int, size)

// 	for i := range slice {
// 		slice[i] = rand.Intn(100)
// 	}
// 	return slice
// }

// func SumSlice(slice []int) int {
// 	var sum int
// 	for _, v := range slice {
// 		sum += v
// 	}
// 	return sum
// }

// func TestGenerateRandomSlice(t *testing.T) {

// 	size := 100

// 	slice := GenerateRandomSlice(size)

// 	if len(slice) != size {
// 		t.Errorf("Expected slice size: %d\nRecieved: %d", size, len(slice))
// 	}
// }

// func BenchmarkGenerateRandomSlice(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		GenerateRandomSlice(1000)
// 	}
// }

// func BenchmarkSumSlice(b *testing.B) {
// 	slice := GenerateRandomSlice(1000)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		SumSlice(slice)
// 	}
// }

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// ================== TESTING
func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 2, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d,%d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Result = %d; Wanted = %d", result, test.expected)
			}
		})

	}

	// func TestAddTableDriven(t *testing.T) {
	// 	tests := []struct{ a, b, expected int }{
	// 		{2, 3, 4},
	// 		{0, 0, 0},
	// 		{-1, 1, 0},
	// 	}

	// 	for _, test := range tests {
	// 		result := Add(test.a, test.b)
	// 		if result != test.expected {
	// 			t.Errorf("Add(%d,%d) = %d; Wanted %d", test.a, test.b, result, test.expected)
	// 		}
	// 	}

	// func TestAdd(t *testing.T) {
	// 	result := Add(2, 4)
	// 	expected := 5

	// 	if result != expected {
	// 		t.Errorf("Add(2,3) = %d; Wanted %d", result, expected)
	// 	}
}
