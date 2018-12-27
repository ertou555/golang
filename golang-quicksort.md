	package main
	import "fmt"
	
	var count int = 0
	
	func quickSort(arr *[]int,low ,high int){
		count ++
		i ,j := low,high
		k:=low
		temp:=(*arr)[low]
		
	for i <= j {
		for j >= k && (*arr)[j] >= temp{ //temp <= (*arr)[j] && k <= j{  注释条件temp在前面则会导致切片越界panic，必须将j >= k条件放在前面
			j--
		}
		if j >= k{
			(*arr)[k] = (*arr)[j]
			k = j
		}
	
		for i <= k && (*arr)[i] <= temp {
			i++
		}
		if i <= k {
			(*arr)[k] = (*arr)[i]
			k = i
		}
	}
		(*arr)[k] = temp
		fmt.Println("k:", k, "i:", i,"j:",j, "low:", low, "high:", high, "arr=", arr)
	
		if k > low + 1{ //这里low + 1 可以减少递归次数，这里也可以改成k > low(这样会增加执行次数) 
			quickSort(arr, low, k-1)
		}
		if k < high - 1{ //这里也可以直接写成k < high (这样会增加执行次数)，通过count计数计算递归次数
			quickSort(arr, k+1, high)
		}
	}
	
	func main() {
		var num=[]int{4,2,3,1,5,8,6,7}
		fmt.Println("before sort :",num,"len:",len(num))
	
		quickSort(&num,0,len(num)-1)
		fmt.Println("after sort :",num)
		fmt.Println("count:",count)
	}
	
	重点：
	1. for j >= k && (*arr)[j] >= temp{//temp <= (*arr)[j] && k <= j{  注释条件temp在前面则会导致切片越界panic，必须将j >= k条件放在前面；
	2. if k < high - 1{ //这里也可以直接写成k < high (这样会增加执行次数)，通过count计数计算递归次数。