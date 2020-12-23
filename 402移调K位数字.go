package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		digit := num[i]
		// 比较的前提是k为正，并且单调栈非空，
		// 这里的for循环类似while循环，
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			// golang没有pop函数，要用切片来实现弹出最后一个元素，
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	// 防止只有一个数，且k为1的情况，
	stack = stack[:len(stack)-k]
	// 把前缀中的0去掉，  strings.TrimLeft用法https://www.jianshu.com/p/a7c5e53e0e62
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}

func main()  {
	var num string = "1432219"
	res := removeKdigits(num, 3)
	fmt.Printf(res)
	//fmt.Println(num)
	//for i:= range num{
	//	fmt.Println(i)
	//	_ = i
	//
	//}

	var n = []byte{49,51,53}
	k := string(n)
	println("aaa", k)
	println(string(n))
}
