package main

import "fmt"

func xorOperation(n int, start int) int {
	r := start
    for k := 0; k < n-1; k++{
		r += 2
		start ^= r
		//fmt.Println(start, r,)
		//fmt.Println(r)
	}
	return start
}

func main()  {
	res := xorOperation(4,3)
	fmt.Print(res)
	//fmt.Println("-------------------------------------")
	//fmt.Println("*************************************")
	//print(3^5)
	//print(3^5^7)
	//print(3^5^7^9)
}












































































