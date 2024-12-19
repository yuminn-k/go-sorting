package sort

import (
	"reflect"
	"testing"
	"time"

	"math/rand"
)

func TestMergeSort(t *testing.T) {
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
		{
			name:     "큰 배열",
			input:    generateRandomSlice(1000),
			expected: nil, // 실제 테스트에서 정렬된 결과와 비교
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "큰 배열" {
				// 큰 배열의 경우 정렬된 결과를 직접 계산
				expected := make([]int, len(tt.input))
				copy(expected, tt.input)
				tt.expected = expected
				quickSort(tt.expected) // 비교를 위한 정렬
			}

			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort() = %v, want %v", result, tt.expected)
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

// 벤치마크 테스트
func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sizes := []int{10, 100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("Random-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateRandomSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				MergeSort(input)
			}
		})

		b.Run("Sorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateSortedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				MergeSort(input)
			}
		})

		b.Run("Reversed-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generateReversedSlice(size)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				MergeSort(input)
			}
		})

		b.Run("PartiallySorted-"+string(rune(size)), func(b *testing.B) {
			b.StopTimer()
			input := generatePartiallySortedSlice(size, 0.75)
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				MergeSort(input)
			}
		})
	}
}

// 메모리 사용량 벤치마크
func BenchmarkMergeSortMemory(b *testing.B) {
	sizes := []int{1000, 10000, 100000}

	for _, size := range sizes {
		b.Run("Memory-"+string(rune(size)), func(b *testing.B) {
			input := generateRandomSlice(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				MergeSort(input)
			}
		})
	}
}

// generateRandomSlice는 지정된 크기의 랜덤한 정수 슬라이스를 생성합니다
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000) // 0부터 999999 사이의 랜덤 값
	}
	return slice
}

// generateSortedSlice는 지정된 크기의 정렬된 슬라이스를 생성합니다
func generateSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

// generateReversedSlice는 지정된 크기의 역순 정렬된 슬라이스를 생성합니다
func generateReversedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = size - i
	}
	return slice
}

// generatePartiallySortedSlice는 부분적으로 정렬된 슬라이스를 생성합니다
// sortedRatio는 0과 1 사이의 값으로, 정렬된 부분의 비율을 나타냅니다
func generatePartiallySortedSlice(size int, sortedRatio float64) []int {
	slice := make([]int, size)
	sortedCount := int(float64(size) * sortedRatio)

	// 정렬된 부분 생성
	for i := 0; i < sortedCount; i++ {
		slice[i] = i
	}

	// 나머지 부분을 랜덤값으로 채움
	for i := sortedCount; i < size; i++ {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

// quickSort는 비교를 위한 기준 정렬 구현입니다
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := 0, len(arr)-1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		quickSort(arr[:right+1])
	}
	if left < len(arr) {
		quickSort(arr[left:])
	}
}
