package mysort

func merge(arr []int, l int, mid int ,r int) {
	aux := make([]int,r-l +1)	
	for i :=l; i<=r; i++ {
		aux[i-l] = arr[i]
	} 
	i := l 
	j := mid +1
	for k := l ; k <= r; k++ {
		if (i >mid ) {
			arr[k] = aux[j-l]
			j++
		}else if  j > r  {
			arr[k] = aux[i-l]
			i++
		}else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		}else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

func mergeSort(arr []int, l int ,r int) {
	if(l >= r) {
		return
	}
	mid := (l+r)/2
	mergeSort(arr,l,mid)
	mergeSort(arr,mid+1,r)
	merge(arr,l,mid,r)
}

func MergeSort(arr []int) []int {
	mergeSort(arr, 0 ,len(arr)-1)
	return arr
}
