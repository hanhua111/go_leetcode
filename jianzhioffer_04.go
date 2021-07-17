package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
    n := len(matrix)
    m := len(matrix[0])
    for i := 0; i < n; i++{
    	if matrix[i][0] <= target && matrix[i][m-1] >= target{
    		left := 0
    		right := m
    		for left < right{
    			mid := (left+right) >> 1
    			if matrix[i][mid]  ==  target{
    				return true
				} else if matrix[i][mid] > target{
					mid = right
				} else {
					mid = left
				}
			}
		} else if matrix[i][0] > target || matrix[n-1][m-1] > target{
			continue
		}
	}
	return false
}

func main()  {
	matrix := [][]int{{1,  4,  7, 11, 15},
		              {2,  5,  8, 12, 19},
					  {3,  6,  9, 16, 22},
					  {10,13, 14, 17, 24},
					  {8, 21, 23, 26, 30}}
	target := 5
	res := findNumberIn2DArray(matrix, target)
	fmt.Println(res)

}














