package sort

import (
	"reflect"
	"testing"
	"time"

	"math/rand"
)

func TestSelectionSort(t *testing.T) {
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
			// 정렬 결과 확인
			result := SelectionSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SelectionSort() = %v, want %v", result, tt.expected)
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
		slice[i] = rand.Intn(1000) // 0-999 사이의 랜덤 값
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

func generateNearlySortedSlice(n int, swaps int) []int {
	slice := generateSortedSlice(n)
	for i := 0; i < swaps; i++ {
		idx1 := rand.Intn(n)
		idx2 := rand.Intn(n)
		slice[idx1], slice[idx2] = slice[idx2], slice[idx1]
	}
	return slice
}

// 벤치마크 테스트
func BenchmarkSelectionSort(b *testing.B) {
	// 난수 생성기 초기화
	rand.Seed(time.Now().UnixNano())

	// 테스트할 배열 크기들
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
				SelectionSort(inputCopy)
			}
		})

		// 2. 이미 정렬된 데이터
		b.Run("Sorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateSortedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				SelectionSort(inputCopy)
			}
		})

		// 3. 역순 정렬된 데이터
		b.Run("Reversed-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateReversedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				SelectionSort(inputCopy)
			}
		})

		// 4. 부분적으로 정렬된 데이터 (75% 정렬)
		b.Run("PartiallySorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generatePartiallySortedSlice(size, 0.75)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				SelectionSort(inputCopy)
			}
		})

		// 5. 거의 정렬된 데이터 (몇 개의 요소만 잘못된 위치)
		b.Run("NearlySorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateNearlySortedSlice(size, size/10) // 10%의 요소만 섞음
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				inputCopy := make([]int, len(input))
				copy(inputCopy, input)
				SelectionSort(inputCopy)
			}
		})
	}
}

// 메모리 사용량 벤치마크
func BenchmarkSelectionSortMemory(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run("Memory-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateRandomSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				SelectionSort(input)
			}
		})
	}
}

// 다양한 데이터 분포에 대한 벤치마크
func BenchmarkSelectionSortDistribution(b *testing.B) {
	size := 1000

	// 1. 모든 요소가 같은 경우
	b.Run("AllSame", func(b *testing.B) {
		input := make([]int, size)
		for i := 0; i < b.N; i++ {
			SelectionSort(input)
		}
	})

	// 2. 두 가지 값만 있는 경우
	b.Run("TwoValues", func(b *testing.B) {
		input := make([]int, size)
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				input[i] = 1
			}
		}

		for i := 0; i < b.N; i++ {
			inputCopy := make([]int, len(input))
			copy(inputCopy, input)
			SelectionSort(inputCopy)
		}
	})
}
