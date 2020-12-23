package main

import (
	"fmt"
)

//func maxLengthBetweenEqualCharacters(s string) int {
//	var res int = -1
//    var data = make(map[byte]int)
//	for ind, val := range s{
//        //注意val的类型是int，所以要用byte转为字符类型
//		if i,ok := data[byte(val)]; ok{
//			k := ind-i-1
//			if k > res{
//				res = k
//			}
//		}else{
//			data[byte(val)] = ind
//		}
//	}
//	return res
//}

func maxLengthBetweenEqualCharacters(s string) int {
	var res int = -1
	var data = make(map[int32]int)
	for ind, val := range s{
		//注意val的类型是int，所以要用byte转为字符类型
		if i,ok := data[val]; ok{
			k := ind-i-1
			if k > res{
				res = k
			}
		}else{
			data[val] = ind
		}
	}
	return res
}

func main(){
	s := "aa"
	maxLengthBetweenEqualCharacters(s)
	fmt.Println(s)
}

