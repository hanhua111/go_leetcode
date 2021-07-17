package main

// 自己的代码
func findRepeatNumber(nums []int) int {
	//index := 0
	//k := nums[0]
	temp := 0
	for index:= 0; index<len(nums); index++{
		if index != nums[index]{
			for temp != index{
				// 一旦找到不相等的说明重复了，此时不再直接返回结果，跳出循环，
				if temp != nums[nums[index]]{
					temp = nums[nums[index]]
					nums[nums[index]] = nums[index]
					nums[index] = temp
					//k = temp
				} else {
					return temp
				}
			}
		} else {
			continue
		}
	}
	//for k,v := range nums{
	//
	//}
	return -1
}

func main()  {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(nums)

}

















