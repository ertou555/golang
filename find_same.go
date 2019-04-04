package main

import (
	"fmt"
)

/*现有两个从小到大排好序的 int 数组（每个数组自身没有重复元素）。请找出所有在
这两个数组中都出现过的数。请写一个函数，输入为两个数组。*/

func find_same(a []int64, b []int64){
	i := 0
	j := 0

	for i < len(a) && j < len(b){  //因为是有序数组，所以从最小的元素开始比较，类似快排算法
		if a[i] > b[j] {
			j++
		}else if a[i] < b[j] {
				i++
		}else{
					fmt.Println("same:",a[i])
					i++
					j++
		}
	}
}

func main(){
	var a []int64 = []int64{1,3,5,6,7}
	var b []int64 = []int64{3,7,9,12,33,84}

	find_same(a,b)
}

/*
//下面是对无序数组找出相同元素的方法
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func main(){
	var a []int64 = []int64{1,3,5,6,7}
	var b []int64 = []int64{3,7,9,12,33,84}

	var list map[string]int64 = make(map[string]int64)

	for i := 0; i < len(a);i++ {
		h := sha256.New()
		buf := Int64ToBytes(a[i])
		h.Write(buf)
		data := h.Sum(nil)
		list[string(data)] = a[i]
	}

	for j := 0; j < len(b); j++{
		h := sha256.New()
		buf := Int64ToBytes(b[j])
		h.Write(buf)
		data := h.Sum(nil)

		if value, ok := list[string(data)]; ok{
			fmt.Println("same :",value)
		}
	}
}
*/
