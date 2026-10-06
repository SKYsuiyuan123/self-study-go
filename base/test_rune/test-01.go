package main

import "fmt"

// 字符

// rune: int32 类型的别名。用来代表一个 Unicode 字符。一个 rune 就存一个汉字 /emoji/ 任意 unicode 字符。
// byte: uint8 类型的别名。单个字节，适合 ASCII 英文数字；中文一个汉字占多个 byte。

func main() {
	str1 := 'a'

	var str2 byte = 'a' // 只能有一个字符
	str3 := '中'         // rune 类型，也是只能有一个字符。

	fmt.Println("str1, str2, str3 ==>", str1, str2, str3) // 97 97 20013

	fmt.Println("----------------分割线---111-----------------")

	s1 := "峰峰123"

	// ❌ 用普通for，取的是byte，中文会拆成字节
	for i := 0; i < len(s1); i++ {
		fmt.Println("s1 for ==>", s1[i]) // s1[i] 类型是 byte
		// s1 for ==> 229
		// s1 for ==> 179
		// s1 for ==> 176
		// s1 for ==> 229
		// s1 for ==> 179
		// s1 for ==> 176
		// s1 for ==> 49
		// s1 for ==> 50
		// s1 for ==> 51
	}

	fmt.Println("----------------分割线---222-----------------")

	// ✅ for range 遍历string，拿到的是 rune。
	for index, val := range s1 {
		fmt.Println("s1 for range ==>", index, val, string(val)) // val 类型是 rune
		// s1 for range ==> 0 23792 峰
		// s1 for range ==> 3 23792 峰
		// s1 for range ==> 6 49 1
		// s1 for range ==> 7 50 2
		// s1 for range ==> 8 51 3
	}

	fmt.Println("----------------分割线---333-----------------")

	s1Count := len([]rune(s1))

	fmt.Println("s1Count ==>", s1Count) // 5, s1 字符串中共有 5 个字符(rune)。
	fmt.Println("s1 len ==>", len(s1))  // 9, s1 字符串中共有 9 个字节(byte)，其中 一个汉字占三个字节。
}
