package main

import "fmt"

// 将一个字符串中的空格替换成 "%20"。
// Input:
//"A B"
//
//Output:
//"A%20B"
func replaceSpace(s string) string {
	str1 := ""
	for _, value := range s {
		if value == ' ' {
			str1 += "%20"
		} else {
			str1 += string(value)
		}
	}
	return str1
}

func main() {
	str := "A B CD E F G"
	fmt.Println(replaceSpace(str))
}
