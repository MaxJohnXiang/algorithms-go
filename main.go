package main
import "./TestHelper"
import "./mysort"
func main() {
	//a := []int{1,3,4,5,6,7,8,9,2,10}
	//fmt.Println(mysort.SelectSort(a))
	arr1 := TestHelper.AenerateRandomArray(100, 1,300)
	arr2 := make([]int, len(arr1))
	arr3 := make([]int, len(arr1))
	arr4 := make([]int, len(arr1))
	arr5 := make([]int, len(arr1))
	copy(arr2, arr1)
	copy(arr3, arr1)
	copy(arr4, arr1)
	copy(arr5, arr1)
 	TestHelper.TestSort("SelectSort", mysort.SelectSort,arr1)
	TestHelper.TestSort("InsertSort", mysort.InsertSort,arr2)
	TestHelper.TestSort("MergeSort", mysort.MergeSort,arr3)
	TestHelper.TestSort("MergeSort2", mysort.MergeSort2,arr4)
	TestHelper.TestSort("QuickSort2", mysort.QuickSort,arr5)
}


