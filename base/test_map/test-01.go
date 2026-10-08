package main

import "fmt"

// map 类型

// map 是引用数据类型
// map 必须初始化后才能使用，使用 make 或者 map[string]string{} 字面量创建。

// map 结构，底层是一个 hash 表，无序的。

// map 不能并发读写，要么加锁(sync.Mutex), 要么用 sync.Map, 要么每个 goroutine 自己一份最后合并。

func main() {
	scores := map[string]string{
		"语文": "90",
	}

	// 1. 新增键

	scores["物理"] = "80"

	fmt.Printf("scores: %#v \n", scores) // scores: map[string]string{"物理":"80", "语文":"90"}

	// 2. 删除键

	delete(scores, "a") // 可以删除不存在的键，但是这没有意义。
	delete(scores, "物理")

	fmt.Println("scores ==>", scores) // map[语文:90]

	// 3. 读取 键值

	val1, ok1 := scores["物理"]
	fmt.Printf("val1: %#v, ok1: %#v \n", val1, ok1) // val1: "", ok1: false. false: 不存在该键。

	val2, ok2 := scores["语文"]
	fmt.Printf("val2: %#v, ok2: %#v \n", val2, ok2) // val2: "90", ok2: true. true: 存在该键。

	// 4. 修改

	_, ok3 := scores["语文"]

	// 判断是否存在该 键
	if ok3 {
		scores["语文"] = "70"
	}

	fmt.Printf("scores: %#v \n", scores) // scores: map[string]string{"语文":"70"}

	fmt.Println("----------------分割线---111-----------------")

	// 5. map 扩容

	var hobby map[string]string

	if hobby == nil {
		fmt.Println("hobby map 的默认值是 nil") // hobby map 的默认值是 nil
	}

	// hobby["学习1"] = "good1" // 直接给 nil 的 map 新增键会报错

	// 需要先给 map 扩容，不然直接 新增键会报错。
	// 可以使用 make 或者 在初始化的时候直接定义好 map 的值 (var hobby map[string]string = map[string]string{})。

	hobby = make(map[string]string, 5)

	// len(map): 返回 map 键的个数

	fmt.Printf("hobby: %#v, len: %v \n", hobby, len(hobby)) // hobby: map[string]string{}, len: 0

	// 新增键会自动扩容
	hobby["学习1"] = "good1"
	hobby["学习2"] = "good2"
	hobby["学习3"] = "good3"
	hobby["学习4"] = "good4"
	hobby["学习5"] = "good5"
	hobby["学习6"] = "good6"

	fmt.Printf("hobby: %#v, len: %v \n", hobby, len(hobby)) // hobby: map[string]string{"学习1":"good1", "学习2":"good2", "学习3":"good3", "学习4":"good4", "学习5":"good5", "学习6":"good6"}, len: 6

	fmt.Println("----------------分割线---222-----------------")

	// 6. map 键对应的 value 是一个函数

	var fnMap = map[string]func(){
		"login":    login1,
		"register": register1,
	}

	fun, ok4 := fnMap["login"]

	if ok4 { // 存在该函数，就执行该函数。
		fun() // login1 func...
	}

	fmt.Println("----------------分割线---333-----------------")

	// 7. map 是引用类型

	userInfo3 := make(map[string]string)

	userInfo3["name"] = "sky"
	userInfo3["age"] = "13"

	// 直接赋值给 userInfo4，userInfo4 和 userInfo3 指向同一个 map 的内存地址。

	userInfo4 := userInfo3

	// 修改 userInfo4 的值，userInfo3 的值也会被修改。
	userInfo4["age"] = "21"

	fmt.Printf("userInfo3: %#v \n", userInfo3) // userInfo3: map[string]string{"age":"21", "name":"sky"}
	fmt.Printf("userInfo4: %#v \n", userInfo4) // userInfo4: map[string]string{"age":"21", "name":"sky"}

	// 新增 userInfo4 的键值对，userInfo3 也会被修改。
	userInfo4["age2"] = "30"

	fmt.Printf("userInfo3: %#v \n", userInfo3) // userInfo3: map[string]string{"age":"21", "age2":"30", "name":"sky"}
	fmt.Printf("userInfo4: %#v \n", userInfo4) // userInfo4: map[string]string{"age":"21", "age2":"30", "name":"sky"}

	fmt.Println("----------------分割线---444-----------------")

	// 8. map 类型的切片 (切片中每一个元素的值都是 map 类型)

	var userInfos = make([]map[string]string, 3) // slice 的长度为 3，容量为 3

	fmt.Printf("userInfos: %#v, len: %v, cap: %v \n", userInfos, len(userInfos), cap(userInfos)) // userInfos: []map[string]string{map[string]string(nil), map[string]string(nil), map[string]string(nil)}, len: 3, cap: 3

	// map 的默认值是 nil
	if userInfos[0] == nil {
		fmt.Println("userInfos[0] map 的默认值是 nil") // userInfos[0] map 的默认值是 nil
	}

	// 初始化 userInfos[0]
	userInfos[0] = map[string]string{
		"name": "sky",
		"age":  "13",
	}

	fmt.Printf("userInfos[0]: %#v \n", userInfos[0]) // userInfos[0]: map[string]string{"age":"13", "name":"sky"}

	// 初始化 userInfos[1]
	userInfos[1] = make(map[string]string)
	userInfos[1]["name"] = "blue"

	fmt.Printf("userInfos: %v \n", userInfos) // userInfos: [map[age:13 name:sky] map[name:blue] map[]]

	age3, ok5 := userInfos[1]["age"]
	fmt.Printf("age3: %#v, ok5: %#v \n", age3, ok5) // age3: "", ok5: false

	for index, itemMap := range userInfos {
		fmt.Printf("userInfos index: %v, itemMap: %#v \n", index, itemMap)
		// userInfos index: 0, itemMap: map[string]string{"age":"13", "name":"sky"}
		// userInfos index: 1, itemMap: map[string]string{"name":"blue"}
		// userInfos index: 2, itemMap: map[string]string(nil)

		for key, mapVal := range itemMap {
			fmt.Printf("mapVal key: %v, mapVal: %#v \n", key, mapVal)
			// mapVal key: name, mapVal: "sky"
			// mapVal key: age, mapVal: "13"
			// mapVal key: name, mapVal: "blue"
		}
	}

	fmt.Println("----------------分割线---555-----------------")

	// 9. 值为 切片类型的 map

	// 如果我们想在 map 对象中存放一系列的属性的时候，我们就可以把 map 类型的值定义成切片。

	userInfos2 := make(map[string][]string, 3)

	fmt.Printf("userInfos2: %#v, len: %v \n", userInfos2, len(userInfos2)) // userInfos2: map[string][]string{}, len: 0

	userInfos2["name"] = []string{"sky", "blue", "green"}
	userInfos2["age"] = []string{"13", "14", "15"}

	fmt.Printf("userInfos2: %#v \n", userInfos2) // userInfos2: map[string][]string{"age":[]string{"13", "14", "15"}, "name":[]string{"sky", "blue", "green"}}

	for key, mapVal := range userInfos2 {
		fmt.Printf("userInfos2 key: %v, mapVal: %#v \n", key, mapVal)
		// userInfos2 key: age, mapVal: []string{"13", "14", "15"}
		// userInfos2 key: name, mapVal: []string{"sky", "blue", "green"}

		for index, item := range mapVal {
			fmt.Printf("mapVal index: %v, item: %#v \n", index, item)
			// mapVal index: 0, item: "sky"
			// mapVal index: 1, item: "blue"
			// mapVal index: 2, item: "green"

			// mapVal index: 0, item: "13"
			// mapVal index: 1, item: "14"
			// mapVal index: 2, item: "15"
		}
	}

	fmt.Println("----------------分割线---666-----------------")

}

func login1() {
	fmt.Println("login1 func...")
}

func register1() {
	fmt.Println("register1 func...")
}
