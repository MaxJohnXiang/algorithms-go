package TestHelper
import "math/rand"
import "time"
import "fmt"

type fn func([]int) []int
func AenerateRandomArray(n int, rangeL int, rangeR int) []int {
	a := make([]int, n)
	for i := 0; i< n; i++ {
	   a[i] = rand.Intn(rangeR - rangeL) + rangeL
	}
	return a;
}


//test function 
func TestSort(sortName string,f fn,arr []int) {
	start := time.Now()
	arr2 := f(arr)
    elapsed := time.Since(start)
	isSuccess := isSort(arr2)
	fmt.Printf("%s took %v, sort success:%t \n",sortName, elapsed, isSuccess)
}

func isSort(arr []int) bool {
	for i:=1;i< len(arr); i++ {
		//如果当前索引小于上一个索引，说明排序失败
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}



