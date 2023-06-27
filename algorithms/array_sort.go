package algorithms

func Swap(elem1 *int, elem2 *int) {
	temp := *elem1
	*elem1 = *elem2
	*elem2 = temp
}

// Insertion sort: O(n^2)
func InsertionSort(arr []int) {

	for i := 1; i < len(arr); i++ {

		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			Swap(&arr[j], &arr[j-1])
		}

	}
}

// Selection sort: O(n^2)
func SelectionSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				Swap(&arr[i], &arr[min])
				min = j
			}
		}
	}

}

// Bubble sort: O(n^2)
func BubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr)-i-1; j++ {

			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}
