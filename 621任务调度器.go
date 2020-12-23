package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	m := map[byte]int{}
	for _,val := range tasks{
		m[val] += 1
	}
	fmt.Println(m)
	l := 1
	k := 0
	for _, val := range m{
		if k > val{
			// 注意大于的时候什么也不做，
			//k = val
			//l = 1-15
		} else{
			if k == val{
				l += 1
			} else {
				k = val
				l = 1
			}
		}
	}
	w := len(tasks)
	if w < (n+1)*(k-1)+l {
	    w = (n+1)*(k-1)+l
	}
	return w
}

func main()  {
	tasks := []byte{'A','A','A','B','B','B','C','D','E','F','G','H','I','J','K'}
	n := 7
	res := leastInterval(tasks,n)
	fmt.Println(res)
	
}