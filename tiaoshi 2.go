package main

import "fmt"

func main() {
	type Student struct {
		name string
		age int
	}
	s := new(Student)  //使用new创建一个 *Student 类型对象
	fmt.Println("s == nil", s == nil) //false
	fmt.Printf("%T\n", s)
	//fmt.Println(*s == nil) //编译错误：cannot convert nil to type Student
	fmt.Printf("%T\n", s)  //*test.Student
	fmt.Printf("%T\n", *s) //test.Student<pre name="code" class="plain">
	fmt.Printf("%T\n", &s) //test.Student<pre name="code" class="plain">

}
