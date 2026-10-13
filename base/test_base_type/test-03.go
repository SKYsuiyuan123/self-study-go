package main

import (
	"fmt"
	"strings"
)

// 字符串

// unsafe.Sizeof 功能：unsafe.Sizeof(x) 返回类型 x 占用的字节数，参数是表达式、编译期求值、不执行。
// 基本类型大小固定，但结构体要考虑内存对齐，字段顺序会影响总大小。
// int/指针大小取决于平台（32/64 位）。
// 常用来看结构体布局、切片头大小，配合 Alignof、Offsetof 理解内存。属于 unsafe，慎用。

// 字符串常见操作

// 1. 多行字符串 ``，字符串拼接可以使用 + 或者 fmt.Sprintf()，也可以使用 strings.Builder 进行拼接。
// 2. 分割字符串使用 strings.Split()，返回切片。
// 3. 把一个切片转换为字符串使用 strings.Join() 返回字符串。
// 4. 判断字符串是否包含子串使用 strings.Contains()，返回 bool。
// 5. 判断字符串是否以某个子串开头使用 strings.HasPrefix()，返回 bool。
// 6. 判断字符串是否以某个子串结尾使用 strings.HasSuffix()，返回 bool。
// 7. 查找子串在字符串中第一次出现的位置使用 strings.Index(s, substr)，返回索引(返回的事字节下标，不是字符下标)。不存在返回 -1。
// 8. strings.LastIndex(s, substr) 查找最后一次出现的位置(返回的事字节下标，不是字符下标)。不存在返回 -1。注意区分 汉字和英文的不同。substr 是空串 "" 时，返回 len(s)。
// 9. 查看字符串长度 len(s) 返回字节数，len([]rune(s)) 返回字符数。
// 10. strings.Repeat(s, count) 重复字符串 count 次，返回新字符串。
// 11. 循环字符串需要使用 for range 遍历，range 会按字符遍历，返回索引和字符 rune。
// 12. 修改字符串需要先转换为 []rune 或 []byte, 然后修改后再转换为字符串。无论哪种转换，都会重新分配内存，原来的字符串不会被修改。

func main() {
	str1 := `"你好" + "go"`

	fmt.Println("str1 ==>", str1)            // str1 ==> "你好" + "go"
	fmt.Printf("str1 len: %v \n", len(str1)) // str1 len: 15

	str2 := "好吧"
	str3 := fmt.Sprintf("%v %v", str1, str2)

	fmt.Println("str3 ==>", str3) // str3 ==> "你好" + "go" 好吧

	arr1 := strings.Split(str1, "+")
	fmt.Printf("arr1: %#v \n", arr1) // arr1: []string{"\"你好\" ", " \"go\""}

	arr2 := []string{"JS", "TS", "GO"}
	str4 := strings.Join(arr2, "_")

	fmt.Println("str4 ==>", str4) // str4 ==> JS_TS_GO

	fmt.Println("----------------分割线---111-----------------")

	bool1 := strings.Contains(str1, "go")
	fmt.Println("bool1 ==>", bool1) // bool1 ==> true

	bool2 := strings.Contains(str1, "abc")
	fmt.Println("bool2 ==>", bool2) // bool2 ==> false

	bool3 := strings.HasPrefix(str1, "abc")
	bool4 := strings.HasSuffix(str1, "go")
	bool5 := strings.HasSuffix(str1, `go"`)

	fmt.Println("bool3 ==>", bool3) // bool3 ==> false
	fmt.Println("bool4 ==>", bool4) // bool4 ==> false
	fmt.Println("bool5 ==>", bool5) // bool5 ==> true

	index1 := strings.Index(str1, "abc")
	index2 := strings.Index(str1, `go"`)

	fmt.Println("index1 ==>", index1) // index1 ==> -1
	fmt.Println("index2 ==>", index2) // index2 ==> 12

	index3 := strings.LastIndex(str2, "好")
	index4 := strings.LastIndex(str2, `吧`)

	fmt.Println("index3 ==>", index3) // index3 ==> 0
	fmt.Println("index4 ==>", index4) // index4 ==> 4

	fmt.Println("----------------分割线---222-----------------")

	str10 := "你好 Go语言"

	byteStr10 := []rune(str10)

	byteStr10[0] = '不'

	str11 := string(byteStr10)

	fmt.Printf("str11: %v, type: %T \n", str11, str11) // str11: 不好 Go语言, type: string
}
