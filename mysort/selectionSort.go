package mysort

func SelectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		//find (i,n) the min number
		minIndex := i
		for j := i+1; j <len(arr);j++ {
			if ( arr[j] < arr[minIndex]) {
				// after end the loop, minIndex will be the min number
				minIndex = j
			}
		}
				arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}	
		return arr
}
