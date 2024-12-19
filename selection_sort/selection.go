package sort

// SelectionSort performs selection sort on an integer slice
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stable: No
func SelectionSort(arr []int) []int {
	n := len(arr)
	// 원본 배열을 수정하지 않기 위해 복사본 생성
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		// 현재 위치를 최소값 위치로 가정
		minIdx := i

		// i 이후의 모든 요소들과 비교하여 최소값 찾기
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}

		// 최소값을 현재 위치로 이동
		if minIdx != i {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}

	return result
}
