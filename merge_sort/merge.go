package sort

// MergeSort performs merge sort on an integer slice
// Time Complexity: O(n log n)
// Space Complexity: O(n)
// Stable: Yes
func MergeSort(arr []int) []int {
	n := len(arr)
	// 원본 배열을 수정하지 않기 위해 복사본 생성
	result := make([]int, n)
	copy(result, arr)

	// 배열의 길이가 1 이하면 이미 정렬된 상태
	if n <= 1 {
		return result
	}

	// 배열을 반으로 나눔
	mid := n / 2
	left := MergeSort(result[:mid])
	right := MergeSort(result[mid:])

	// 정렬된 두 배열을 반환
	return merge(left, right)
}

// merge 함수는 두 개의 정렬된 배열을 하나의 정렬된 배열로 병합
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	// 두 배열을 비교하면서 작은 값을 결과 배열에 삽입
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// 남은 요소들을 결과 배열에 복사
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
