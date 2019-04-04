package main

import (
	"fmt"
	"sort"
)

var score map[byte]int64 = map[byte]int64{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 1,
}

func cal_other_cow(other []byte) int64{
	str := string(other)

	remain := (score[str[0]]+score[other[1]]) % 10
	fmt.Println(remain)
	if remain == 0{
		return 10
	}else {
		return remain
	}
}

//求出5个数种可以组合成10的倍数的3个数，若找到，则返回剩下2个数的切片
//返回值：int64代表是否存在牛，[] byte剩下2个数。
func cal_cow(a string) (int64,[] byte) {  
	len := len(a)
	var slice []byte
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			for k := 0; k < len; k++ {
				if i != j && i != k && j != k {
					if (score[a[i]]+score[a[j]]+score[a[k]])%10 == 0 {
						fmt.Printf("%d,%d,%d\n", score[a[i]], score[a[j]], score[a[k]])
						for m := 0; m < len; m++{
							if m != i && m != j && m != k {
								slice = append(slice,a[m])
							}
						}
						return 1,slice
					}
				}
			}
		}
	}

	return 0,nil
}
func main() {
	a := "4579K"
	b := "AAAA2"

	reta, slicea := cal_cow(a)
	retb, sliceb := cal_cow(b)

	if reta > retb{
		fmt.Println(1)
		return
	}else if reta < retb {
		fmt.Println(-1)
		return
	}
	//这里都存在牛，判断剩下2张的大小
	if len(slicea) == len(sliceb) && len(sliceb) == 2{
		ret1 := cal_other_cow(slicea)
		ret2 := cal_other_cow(sliceb)
		if ret1 > ret2{
			fmt.Println(1)
			return
		}else if ret1 < ret2 {
			fmt.Println(-1)
			return
		}else {
			fmt.Println(0)
			return
		}
	}else{ //都不存在牛，先对各自拥有的牌排序，然后比较。
		slicea := mapToSlice(a)
		sliceb := mapToSlice(b)
		ret := compare(slicea,sliceb)
		fmt.Println(ret)
	}
}

func compare(a []int64, b []int64) int{
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		}else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func mapToSlice(str string) []int64 {
	s := make([]int64, 0, len(str))

	for i := 0; i < len(str); i++{
		if str[i] == 'A'{
			s = append(s,14)
		}else {
			s = append(s, score[str[i]])
		}
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
	return s
}

