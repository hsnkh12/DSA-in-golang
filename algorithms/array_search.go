package algorithms

// O(n)
func LinearSearch(arr []interface{}, item interface{}) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}

	return false
}

// O(logn)
func BinarySearch(arr []int, low int, high int, value int) bool {

	if low > high {
		return false
	}

	middle := (low + high) / 2

	if arr[middle] == value {
		return true
	} else if arr[middle] > value {
		return BinarySearch(arr, low, middle-1, value)
	} else {
		return BinarySearch(arr, middle+1, high, value)
	}

}
