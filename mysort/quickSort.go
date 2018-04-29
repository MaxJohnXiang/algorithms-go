package mysort
func __partition(arr []int, l int ,r int) int {
	v := arr[l]
	j := l
	for i := l +1; i <= r; i++ {
		if (arr[i] < v) {
			arr[j + 1] , arr[i] = arr[i], arr[j +1 ]			
			j ++
		}
	}
	arr[l],arr[j] = arr[j] ,arr[l]
	return j
}

func __qucikSort(arr []int,l int , r int ) {
	if (l > r) {
		return
	}
    p := __partition(arr, l, r)
	__qucikSort(arr, l, p-1)
	__qucikSort(arr, p+1, r)
} 

func QuickSort(arr []int) []int {
	__qucikSort(arr,0,len(arr)-1) 
	return arr
}
