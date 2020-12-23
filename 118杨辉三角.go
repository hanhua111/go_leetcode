package main

import "fmt"

//// 自己的代码
//func generate(numRows int) [][]int {
//	res := [][]int{}
//	c := 0
//	if numRows == 0{
//		return res
//	}else {
//		// 当numRous为1时，直接返回
//		res = append(res, []int{1})
//		for ; numRows > 1; numRows-- {
//			// 为了取出二维数组的最后一个元素
//			l := len(res[c])
//			temp := []int{1}
//			for i,k := range res[l-1]{
//				// 因为第一个元素没有前一个元素，无法相加，所以跳过，
//				if i != 0{
//					temp = append(temp, k+res[l-1][i-1])
//				}
//			}
//			// 补上最后一个元素1，
//			temp = append(temp, 1)
//			res = append(res, temp)
//			c++
//		}
//		return res
//	}
//}


// 官方代码
// 写法很巧妙，
func generate(numRows int) [][]int {
	// 直接new了一个满足结果shape的二维数组，
	ans := make([][]int, numRows)
	// 只写一个i的时候就是索引，
	for i := range ans {
		// 初始化长度，
		ans[i] = make([]int, i+1)
		// 先把首末的填上，
		ans[i][0] = 1
		ans[i][i] = 1
		// i-1表示上一行，列j必须大于等于1才有前一个数，
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}


func main()  {
	i := 5
	res := generate(i)
	fmt.Println(res)

}