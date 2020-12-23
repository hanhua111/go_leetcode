package main

// 自己的代码，写的较繁琐，不需要这么臃肿，
//func count (s string) map[int32]int {
//	res := map[int32]int{}
//	for _,v := range s{
//		_, exist := res[v]
//		if exist{
//			res[v] += 1
//		} else {
//			res[v] = 1
//		}
//	}
//	return res
//}
//func isAnagram(s string, t string) bool {
//	s_map := count(s)
//	t_map := count(t)
//	for k := range s_map{
//		v, exist := t_map[k]
//		if exist && v == s_map[k]{
//			delete(t_map, k)
//		} else {
//			return false
//		}
//	}
//	if len(t_map) == 0{
//		return true
//	} else {
//		return false
//	}
//}

//// 处理字符串时，学会这个索引技巧，
//func isAnagram(s, t string) bool {
//	var c1, c2 [26]int
//	for _, ch := range s {
//		c1[ch-'a']++
//	}
//	for _, ch := range t {
//		c2[ch-'a']++
//	}
//	fmt.Println(c1)
//	//println(c2)
//	return c1 == c2
//}


func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	// 这里写的很巧妙，无需再弄一个数组来记录比较，
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}

func main()  {
	s := "anagram"
	t := "nagaram"
	res := isAnagram(s, t)
	println(res)
	m := "79879"
	for i,v := range m{
		println(i, v)
	}

}


