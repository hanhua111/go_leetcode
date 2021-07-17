package main

import "fmt"

func fib(n int) int {
    a := 0
    b := 1
    if (n==0) || (n==1){
    	return n
    } else {
    	res := 0
    	for n > 1{
    		// 这里加的时候就要取余，否则会加的超出范围，
    		res = (a+b) % (1000000007)
    		a = b
    		b = res
    		n -= 1
		}
		return (res % 1000000007)
	}
}
func main()  {
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))

}

// 优秀解法，
//func fib(N int) int {
//	f0, f1 := 0, 1
//	for i := 0; i < N; i++ {
//		f0,f1 = f1,(f0 + f1) % 1000000007;
//	}
//	return f0
//}
//
//
//作者：deepwjp
//链接：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/solution/jian-zhi-offer-10-i-fei-bo-na-qi-shu-lie-by-deepwj/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

















