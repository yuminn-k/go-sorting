# go-sorting

A collection of sorting algorithms implemented in Go.

## Implemented Algorithms

### Basic Sorts
- Bubble Sort
- Selection Sort
- Insertion Sort

### Advanced Sorts
- Merge Sort
- Quick Sort
- Heap Sort

### Special Sorts
- Counting Sort
- Radix Sort
- Bucket Sort

## Usage

```go
import "github.com/yuminn-k/go-sorting"

// Example usage
numbers := []int{64, 34, 25, 12, 22, 11, 90}
sorted := sorting.BubbleSort(numbers)
```

## Time Complexity

| Algorithm      | Best       | Average    | Worst      | Space    | Stability |
| -------------- | ---------- | ---------- | ---------- | -------- | --------- |
| Bubble Sort    | O(n)       | O(n²)      | O(n²)      | O(1)     | Stable    |
| Selection Sort | O(n²)      | O(n²)      | O(n²)      | O(1)     | Unstable  |
| Insertion Sort | O(n)       | O(n²)      | O(n²)      | O(1)     | Stable    |
| Merge Sort     | O(n log n) | O(n log n) | O(n log n) | O(n)     | Stable    |
| Quick Sort     | O(n log n) | O(n log n) | O(n²)      | O(log n) | Unstable  |
| Heap Sort      | O(n log n) | O(n log n) | O(n log n) | O(1)     | Unstable  |

## Algorithm Details

### Basic Sorts
- **Bubble Sort**: Repeatedly steps through the list, compares adjacent elements and swaps them if they are in the wrong order.
- **Selection Sort**: Divides the input into a sorted and unsorted region, and repeatedly selects the smallest element from the unsorted region.
- **Insertion Sort**: Builds the final sorted array one item at a time, by repeatedly inserting a new element into the sorted portion of the array.

### Advanced Sorts
- **Merge Sort**: Divides the array into two halves, recursively sorts them, and then merges the sorted halves.
- **Quick Sort**: Uses a pivot element to partition the array into two sub-arrays and recursively sorts them.
- **Heap Sort**: Uses a binary heap data structure to sort elements.

### Special Sorts
- **Counting Sort**: Works by counting the number of objects having distinct key values.
- **Radix Sort**: Sorts integers by processing each digit position, starting from the least significant digit.
- **Bucket Sort**: Distributes elements into buckets and then sorts these buckets individually.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by classic sorting algorithms
- Built for educational purposes and practical use