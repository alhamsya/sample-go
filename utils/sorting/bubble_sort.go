package utils

func BubbleSorting(arr []int) []int {
	//looping unordered until index
	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				//swap data
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}

	return arr
}
