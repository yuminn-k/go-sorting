package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
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
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort() = %v, want %v", result, tt.expected)
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

// 벤치마크 테스트 추가
func BenchmarkBubbleSort(b *testing.B) {
	// 다양한 크기의 입력에 대한 벤치마크
	benchCases := []struct {
		name string
		n    int
	}{
		{"Small", 10},
		{"Medium", 100},
		{"Large", 1000},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			// 각 벤치마크 케이스마다 새로운 입력 생성
			input := make([]int, bc.n)
			for i := 0; i < bc.n; i++ {
				input[i] = bc.n - i // 최악의 경우(역순 정렬)로 테스트
			}

			// 벤치마크 실행
			b.ResetTimer() // 타이머 리셋
			for i := 0; i < b.N; i++ {
				BubbleSort(input)
			}
		})
	}

	// 이미 정렬된 배열에 대한 벤치마크 (최선의 경우)
	b.Run("BestCase", func(b *testing.B) {
		input := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			input[i] = i
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			BubbleSort(input)
		}
	})

	// 랜덤 배열에 대한 벤치마크 (평균적인 경우)
	b.Run("RandomCase", func(b *testing.B) {
		input := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			input[i] = i
		}

		// Fisher-Yates 셔플 알고리즘
		for i := len(input) - 1; i > 0; i-- {
			j := i / 2 // 의사 난수 생성
			input[i], input[j] = input[j], input[i]
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			BubbleSort(input)
		}
	})
}
