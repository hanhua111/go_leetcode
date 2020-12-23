package main

import (
	"fmt"
	"sort"
)

var bit = [1e4 + 1]int{}

func init() {
	for i := 1; i <= 1e4; i++ {
		bit[i] = bit[i>>1] + i&1
	}
}

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := bit[x], bit[y]
		return cx < cy || cx == cy && x < y
	})
	return a
}

func main()  {
	var a = []int{1024,512,256,128,64,32,16,8,4,2,1}
	res := sortByBits(a)
	fmt.Println(res)
	fmt.Printf("%v", res)


}
