package main

import "fmt"

func intersection(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}
func main()  {
	var nums1 = []int{1,2,2,3}
	var nums2 = []int{2,3,3,4}
	res := intersection(nums1, nums2)
	fmt.Println(res)

}