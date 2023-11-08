package main

func main() {

}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := arr[:middle]
	right := arr[middle:]
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)

}

func merge(left []int, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else {
			if left[0] <= right[0] {
				merged = append(merged, left[0])
				left = left[1:]
			} else {
				merged = append(merged, right[0])
				right = right[1:]
			}
		}
	}
	return merged
}
