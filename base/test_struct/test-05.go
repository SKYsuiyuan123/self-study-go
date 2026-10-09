package main

import "fmt"

// map 的 value 是结构体

// map 存的是 "值" 而不是 "引用"，所以在 map 中修改结构体的字段时，需要先取出结构体，修改后再放回 map 中。
// map 里的值没法定位到某个字段去原地修改，要么把整个值取出来修改完放回去，要么让 map 直接存指针。

type person5_1 struct {
	name string
	age  int
}

func main() {
	// 1. map 的值是 普通值

	aMap := map[string]string{"a": "aaa"}
	aMap["a"] = "bbb" // 修改 map 中的值

	fmt.Printf("aMap: %#v \n", aMap) // aMap: map[string]string{"a":"bbb"}

	fmt.Println("----------------分割线---111-----------------")

	// 2. map 的值是 结构体

	m2 := map[string]person5_1{
		"alice": {name: "Alice", age: 30},
	}

	// map 底层是哈希表，元素散落在不同的桶 (bucket) 里，而且随着扩容、rehash, 元素可能会移动到不同的桶里，所以 map 里的值没法定位到某个字段去原地修改。
	// 因此 go 规定：map 的值是不可寻址的，不能直接修改 map 中的结构体字段。

	// m2["alice"].age = 31 // ❌ 编译报错：cannot assign to struct field m2["alice"].age in map

	// 2.1 整体替换 修改

	m2["alice"] = person5_1{name: "Alice", age: 31}

	fmt.Printf("m2: %#v \n", m2) // m2: map[string]main.person5_1{"alice":main.person5_1{name:"Alice", age:31}}

	// 2.2 取出修改后再放回 map 中 (推荐)

	p := m2["alice"] // 取出结构体
	p.age = 25       // 修改结构体字段
	m2["alice"] = p  // 放回 map 中

	fmt.Printf("m2: %#v \n", m2) // m2: map[string]main.person5_1{"alice":main.person5_1{name:"Alice", age:25}}

	// 3. map 的值是 结构体指针

	m3 := map[string]*person5_1{
		"alice": &person5_1{name: "Alice", age: 30},
	}

	m4 := map[string]*person5_1{
		"alice": {name: "Alice2", age: 20},
	}

	m3["alice"].age = 32 // ✅ 可以直接修改结构体字段
	m4["alice"].age = 22 // ✅ 可以直接修改结构体字段

	fmt.Printf("m3: %#v \n", m3)                            // m3: map[string]*main.person5_1{"alice":(*main.person5_1)(0x2ca6ff3ba030)}
	fmt.Printf("m3[\"alice\"] age: %d \n", m3["alice"].age) // m3["alice"] age: 32

	fmt.Printf("m4: %#v \n", m4)                            // m4: map[string]*main.person5_1{"alice":(*main.person5_1)(0x2ca6ff3ba048)}
	fmt.Printf("m4[\"alice\"] age: %d \n", m4["alice"].age) // m4["alice"] age: 22
}
