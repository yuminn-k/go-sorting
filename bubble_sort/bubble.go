package sort

// BubbleSort performs bubble sort on an integer slice.
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stable: Yes
func BubbleSort(arr []int) []int {
	n := len(arr)
	// 원본 배열을 수정하지 않기 위해 복사본 생성
	result := make([]int, n)
	copy(result, arr)

	// 최적화를 위한 스왑 발생 여부 체크
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			// 인접한 요소를 비교하여 순서가 잘못되었으면 교환
			if result[i-1] > result[i] {
				result[i-1], result[i] = result[i], result[i-1]
				swapped = true
			}
		}
		// 최적화: 정렬이 완료되면 더 이상 반복하지 않음
		n--
	}

	return result
}
