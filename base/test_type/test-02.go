package main

import (
	"fmt"
	"strconv"
)

// 基本类型 —— 互相转换

// strconv 包 是字符串类型 与 数值类型 互相转换的标准工具

func main() {
	var f1 float32 = 3.14159265359 // float32 会截取(有效小数 —— 从非零位开始计算), 有效精度为 5 - 7 位。
	fmt.Println("f1 ==>", f1)      // 3.1415927

	var f2 float64 = 3.14159265359 // float64 会截取小数，有效精度 15 - 16 位。
	fmt.Println("f2 ==>", f2)      // 3.14159265359

	fmt.Println("----------------分割线---111-----------------")

	// 1. 浮点数转为 int 类型 ==> int() 会丢失精度 (舍弃小数部分)

	f3 := 3.14
	f4 := int(f3)

	fmt.Println("f4 ==>", f4) // 3

	// 2. float64 转 float32 会丢失精度

	var f5 = 3.0000000032123 // 默认 float64 类型
	var f6 float32 = float32(f5)

	fmt.Println("f5 ==>", f5) // 3.0000000032123
	fmt.Println("f6 ==>", f6) // 3

	// 3. int 类型转为 浮点数类型 ==> float64()

	num1 := 20
	num2 := 32.233
	num3 := float64(num1) + num2 // 不能直接相加，必须转换为同一类型。

	fmt.Println("num3 ==>", num3)                   // 52.233
	fmt.Println("float64(num1) ==>", float64(num1)) // 20

	fmt.Println("----------------分割线---222-----------------")

	// 4. 数字转为字符串

	num4 := 1
	num4Str := strconv.Itoa(num4)                         // 底层使用的是 strconv.FormatInt(int64(n), 10)
	fmt.Printf("num4Str ==> %v, %T \n", num4Str, num4Str) // 1, string

	num1Str := fmt.Sprintf("%d", num1)  // 性能最差，不推荐。只适合用来打印日志
	fmt.Println("num1Str ==>", num1Str) // "20"

	num1_1Str := strconv.FormatInt(int64(num1), 10)
	fmt.Println("num1_1Str ==>", num1_1Str) // "20"

	num2Str := strconv.FormatFloat(float64(num2), 'f', -1, 64) // 1: 保留一位小数，-1 保留全部小数
	fmt.Printf("num2Str ==> %v -- %T \n", num2Str, num2Str)    // "32.233" -- string

	fmt.Println("----------------分割线---333-----------------")

	// 5. 字符串 转为数字 (只有 数字类型字符串(整数) 可以转成 数字，其它转换后为 0, 并且报错。)

	str1 := "ag"
	str1Int, err1 := strconv.Atoi(str1)             // strconv.Atoi 内部就是封装的 strconv.ParseInt
	fmt.Println("str1Int, err1 ==>", str1Int, err1) // 0 strconv.Atoi: parsing "ag": invalid syntax

	str2 := "1"
	str2Int, err2 := strconv.Atoi(str2)
	fmt.Println("str2Int, err2 ==>", str2Int, err2) // 1 <nil>

	str3 := "1.2"
	str3Int, err3 := strconv.Atoi(str3)
	fmt.Println("str3Int, err3 ==>", str3Int, err3) // 0 strconv.Atoi: parsing "1.2": invalid syntax

	str4 := "20"
	str4Int, err4 := strconv.ParseInt(str4, 10, 64)
	fmt.Println("str4Int, err4 ==>", str4Int, err4)         // 20 <nil>
	fmt.Printf("str4Int ==> %v -- %T \n", str4Int, str4Int) // 20 -- int64

	str5 := "10.2"
	str5Int, err5 := strconv.ParseInt(str5, 10, 64)
	fmt.Println("str5Int, err5 ==>", str5Int, err5) // 0 strconv.ParseInt: parsing "10.2": invalid syntax

	str6 := "1.2"
	str6F, err6 := strconv.ParseFloat(str6, 64)
	fmt.Println("str6F, err6 ==>", str6F, err6)       // 1.2 <nil>
	fmt.Printf("str6F ==> %v -- %T \n", str6F, str6F) // 1.2 -- float64

	fmt.Println("----------------分割线---444-----------------")

	// 6.数值类型 和 bool类型 不能 互相进行转换

	// bool 想变数值只能手动写 if，数值想变 bool 只能写比较表达式（如 n != 0）。这是 Go 为了类型安全和代码清晰故意做的限制。

	// var num6 int = 1
	// var bool0 = bool(num6) // cannot convert num6 (variable of type int) to type bool

	fmt.Println("----------------分割线---555-----------------")

	// 7. 字符串转为 bool 类型。strconv.ParseBool 该函数返回布尔值。

	// 该函数接受 "1"、"t"、"T"、"TRUE"、"true"、"True"、"0"、"f"、"F"、"FALSE"、"false"、"False" 作为输入。
	// 任何其他值都会返回转换错误。(值是 false)

	// "1", t, T, true, TRUE, True => true
	// "0", f, F, FALSE, false, False => false

	bool1, err10 := strconv.ParseBool("xxx")
	fmt.Println("bool1, err10 ==>", bool1, err10) // false, strconv.ParseBool: parsing "xxx": invalid syntax

	bool2, err11 := strconv.ParseBool("false")
	fmt.Println("bool2, err11 ==>", bool2, err11)           // false <nil> (转成功了，转为 bool 后值是 false。)
	fmt.Printf("bool2 %v %v -- %T \n", bool2, err11, bool2) // false <nil> -- bool

	bool3, err12 := strconv.ParseBool("False")
	fmt.Println("bool3, err12 ==>", bool3, err12) // false <nil>

	bool4, err13 := strconv.ParseBool("FALSE")
	fmt.Println("bool4, err13 ==>", bool4, err13) // false <nil>

	bool5, err14 := strconv.ParseBool("FAlse")
	fmt.Println("bool5, err14 ==>", bool5, err14) // false, strconv.ParseBool: parsing "FAlse": invalid syntax

	bool6, err15 := strconv.ParseBool("0")
	fmt.Printf("bool6 %v -- %T -- %v \n", bool6, bool6, err15) // false -- bool -- <nil>

	bool7, err16 := strconv.ParseBool("f")
	fmt.Printf("bool7 %v -- %T -- %v \n", bool7, bool7, err16) // false -- bool -- <nil>

	bool8, err17 := strconv.ParseBool("F")
	fmt.Printf("bool8 %v -- %T -- %v \n", bool8, bool8, err17) // false -- bool -- <nil>

	bool9, err18 := strconv.ParseBool("true")
	fmt.Printf("bool9 %v -- %T -- %v \n", bool9, bool9, err18) // true -- bool -- <nil>

	bool10, err19 := strconv.ParseBool("True")
	fmt.Printf("bool10 %v -- %T -- %v \n", bool10, bool10, err19) // true -- bool -- <nil>

	bool11, err20 := strconv.ParseBool("TRUE")
	fmt.Printf("bool11 %v -- %T -- %v \n", bool11, bool11, err20) // true -- bool -- <nil>

	bool12, err21 := strconv.ParseBool("1")
	fmt.Printf("bool12 %v -- %T -- %v \n", bool12, bool12, err21) // true -- bool -- <nil>

	bool13, err22 := strconv.ParseBool("t")
	fmt.Printf("bool13 %v -- %T -- %v \n", bool13, bool13, err22) // true -- bool -- <nil>

	bool14, err23 := strconv.ParseBool("T")
	fmt.Printf("bool14 %v -- %T -- %v \n", bool14, bool14, err23) // true -- bool -- <nil>

	// 8. 布尔类型转为 string 类型

	bool1Str := strconv.FormatBool(true)
	fmt.Printf("bool1Str %v -- %T \n", bool1Str, bool1Str) // true -- string

	bool2Str := strconv.FormatBool(false)
	fmt.Printf("bool2Str %v -- %T \n", bool2Str, bool2Str) // false -- string
}
