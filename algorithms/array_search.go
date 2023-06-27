package algorithms

// Binary search: O(nlogn(n)) with pre-sorting the array
func BinarySearch(arr []int, p int, r int, v int) int {

	if p == r {
		if arr[p] == v {
			return 1
		}
		return 0
	}

	q := int((p + r) / 2)

	if v <= arr[q] {
		return BinarySearch(arr, p, q, v)

	}
	return BinarySearch(arr, q+1, r, v)
}
