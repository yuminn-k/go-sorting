package sort

import (
	"reflect"
	"testing"
	"time"

	"math/rand"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "정렬되지 않은 배열",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "이미 정렬된 배열",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "역순으로 정렬된 배열",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "중복된 요소가 있는 배열",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9},
		},
		{
			name:     "음수를 포함한 배열",
			input:    []int{-5, 3, 0, -8, 7, 2, -1},
			expected: []int{-8, -5, -1, 0, 2, 3, 7},
		},
		{
			name:     "빈 배열",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "단일 요소 배열",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertionSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertionSort() = %v, want %v", result, tt.expected)
			}

			// 원본 배열이 변경되지 않았는지 확인
			originalCopy := make([]int, len(tt.input))
			copy(originalCopy, tt.input)
			if !reflect.DeepEqual(tt.input, originalCopy) {
				t.Error("원본 배열이 변경되었습니다")
			}
		})
	}
}

// 벤치마크를 위한 테스트 데이터 생성 헬퍼 함수들
func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func generateSortedSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}

func generateReversedSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = n - i - 1
	}
	return slice
}

func generatePartiallySortedSlice(n int, sortedPortion float64) []int {
	slice := generateSortedSlice(n)
	unsortedStart := int(float64(n) * sortedPortion)
	for i := n - 1; i > unsortedStart; i-- {
		j := rand.Intn(i-unsortedStart) + unsortedStart
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// 벤치마크 테스트
func BenchmarkInsertionSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{10, 100, 1000, 5000}

	for _, size := range sizes {
		// 1. 랜덤 데이터
		b.Run("Random-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateRandomSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				InsertionSort(inputCopy)
			}
		})

		// 2. 이미 정렬된 데이터 (최선의 경우)
		b.Run("Sorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateSortedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				InsertionSort(inputCopy)
			}
		})

		// 3. 역순 정렬된 데이터 (최악의 경우)
		b.Run("Reversed-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateReversedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				InsertionSort(inputCopy)
			}
		})

		// 4. 부분적으로 정렬된 데이터
		b.Run("PartiallySorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generatePartiallySortedSlice(size, 0.75)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				InsertionSort(inputCopy)
			}
		})
	}
}
