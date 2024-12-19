package sort

// InsertionSort performs insertion sort on an integer slice
// Time Complexity: O(n^2)
// Space Complexity: O(1)
// Stable: Yes
func InsertionSort(arr []int) []int {
	n := len(arr)
	// 원본 배열을 수정하지 않기 위해 복사본 생성
	result := make([]int, n)
	copy(result, arr)

	// 두 번째 요소부터 시작하여 적절한 위치에 삽입
	for i := 1; i < n; i++ {
		// 현재 검사하는 요소를 저장
		key := result[i]
		// 정렬된 부분에서 삽입할 위치를 찾음
		j := i - 1

		// key보다 큰 요소들을 뒤로 이동
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}

		// 찾은 위치에 key 삽입
		result[j+1] = key
	}

	return result
}
