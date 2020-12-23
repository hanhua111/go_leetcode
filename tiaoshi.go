package main

import "fmt"

//func main()  {
//	s := "abcdABCD"
//	k1 := []rune(s) // rune是int32类型
//	fmt.Println(k1)
//	k2 := []byte(s) // byte是uint8类型
//	fmt.Println(k2)
//	for _, v := range s{ // 注意这里的v是int32类型
//		println(v)
//	}
//}


func main() {

	slice := []int{0,1,2,3}
	//m := make(map[int]*int)

	for key,val := range slice {
		println(&key)
		println(&val, "kkk")
		//m[key] = &val
	}

	//for k,v := range m {
	//	fmt.Println(k,"->",*v)
	//}
	a := [][]int{{1,2,3},{4,5,6}}
	fmt.Println(a)
	for i,j := range a{
		//fmt.Printf(i)
		println(i)
		println(j)
		fmt.Println(j)

	}

	k := make([][]int, 5)
	for i,j := range k{
		fmt.Println(i,j)
	}
}




