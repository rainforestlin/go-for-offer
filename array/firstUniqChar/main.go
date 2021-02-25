package main

import "fmt"

// 在一个字符串中找到第一个只出现一次的字符，并返回它的位置。
// 字符串只包含 ASCII 码字符。
// Input: abacc
//Output: b
func firstUniqChar(s string) byte {
	var list [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		list[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	input := "abacc"
	output := firstUniqChar(input)
	fmt.Println(string(output))
}
