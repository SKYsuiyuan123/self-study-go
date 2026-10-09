package main

import "fmt"

// 结构体 —— 匿名字段

// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
// 匿名字段默认采用类型名作为字段名，匿名字段在使用时可以直接通过类型名来访问。
// 结构体要求字段名称必须唯一，所以匿名字段在同一个结构体中不能重复出现。(一个结构体中，匿名字段只能出现一个)

// 普通声明里边不推荐，只适合使用在嵌套结构体中，或者是继承的场景中。
type person2_1 struct {
	string
	int
	// string // 不能重复出现
}

// 结构体的字段类型可以是：基本数据类型、切片、数组、指针、函数、接口、结构体等。
// 如果结构体的字段类型是：指针、切片、数组、函数、接口、结构体、map等，那么该字段的零值就是对应类型的零值。
// 如果结构体的字段类型是：指针，slice 和 map 的零值都是 nil，即还没有分配空间。
// 如果需要使用 (指针、slice、map) 类型的字段，需要先为其分配空间(make 或者初始化)，否则会出现空指针异常。

type person2_2 struct {
	name   string
	hobby  []string
	hobby2 []string
	hobby3 [3]string
	map1   map[string]string
	map2   map[string]int
}

// 结构体嵌套 匿名结构体

type person2_3 struct {
	username     string
	password     string
	address2_1   address2_1 // 表示 person2_3 结构体嵌套 address2_1 结构体
	address2_1_2            // 匿名结构体
	abc          string
	email2_1
}

type address2_1 struct {
	name1  string
	phone1 string
	city1  string
}

type address2_1_2 struct {
	name2   string
	phone2  string
	city2   string
	abc     string
	addTime string
}

type email2_1 struct {
	addTime string
}

func main() {
	// 1. 匿名字段

	p1 := person2_1{
		"张三",
		18,
	}

	fmt.Printf("p1: %#v\n", p1)          // p1: main.person2_1{string:"张三", int:18}
	fmt.Println("p1.string:", p1.string) // p1.string: 张三

	fmt.Println("----------------分割线---111-----------------")

	// 2. 结构体的字段类型是：指针、slice、map 的零值都是 nil，即还没有分配空间。

	p2 := person2_2{}

	p2.name = "sky"

	// p2.hobby[0] = "学习" // hobby 还没有分配内存，报错 panic: runtime error: index out of range [0] with length 0
	p2.hobby = append(p2.hobby, "学习")

	fmt.Printf("p2: %#v \n", p2)             // p2: main.person2_2{name:"sky", hobby:[]string{"学习"}, hobby2:[]string(nil), hobby3:[3]string{"", "", ""}, map1:map[string]string(nil), map2:map[string]int(nil)}
	fmt.Printf("p2.hobby: %#v \n", p2.hobby) // p2.hobby: []string{"学习"}

	p2.hobby2 = make([]string, 3) // 使用 make 为切片分配空间
	p2.hobby2[0] = "看书"

	fmt.Printf("p2.hobby2: %#v \n", p2.hobby2) // p2.hobby2: []string{"看书", "", ""}

	p2.hobby3[0] = "运动"                        // 数组已经分配了空间，可以直接使用
	fmt.Printf("p2.hobby3: %#v \n", p2.hobby3) // p2.hobby3: [3]string{"运动", "", ""}

	// p2.map1["address"] = "上海" // map 还没有分配内存，报错 panic: assignment to entry in nil map
	p2.map1 = make(map[string]string, 3) // 使用 make 为 map 分配空间
	p2.map1["address"] = "上海"

	fmt.Printf("p2.map1: %#v \n", p2.map1) // p2.map1: map[string]string{"address":"上海"}

	p2.map2 = map[string]int{"age": 18}    // 使用 map 字面量为 map 分配空间
	fmt.Printf("p2.map2: %#v \n", p2.map2) // p2.map2: map[string]int{"age":18}

	fmt.Printf("p2 %#v \n", p2) // p2 main.person2_2{name:"sky", hobby:[]string{"学习"}, hobby2:[]string{"看书", "", ""}, hobby3:[3]string{"运动", "", ""}, map1:map[string]string{"address":"上海"}, map2:map[string]int{"age":18}}

	fmt.Println("----------------分割线---222-----------------")

	// 3. 结构体嵌套结构体

	p3 := person2_3{}

	p3.username = "sky"
	p3.password = "123"
	p3.address2_1 = address2_1{
		phone1: "1234567",
		name1:  "浦东",
		city1:  "上海",
	}

	fmt.Printf("p3: %#v \n", p3) // p3: main.person2_3{username:"sky", password:"123", address2_1:main.address2_1{name1:"浦东", phone1:"1234567", city1:"上海"}, address2_1_2:main.address2_1_2{name2:"", phone2:"", city2:"", abc:"", addTime:""}, abc:"", email2_1:main.email2_1{addTime:""}}

	p3.city2 = "abc"             // 只能在匿名结构体中访问 city2 字段
	p3.address2_1_2.city2 = "北京" // 也可以通过 address2_1_2 来访问 city2 字段

	p3.name2 = "大兴" // 当访问结构体成员时，会先在结构体中查找该字段，找不到再去匿名结构体中查找该字段，所以这里访问的是 p3.address2_1_2.name2 字段。

	p3.abc = "abc"               // 当访问结构体成员时，会先在结构体中查找该字段，找不到再去匿名结构体中查找该字段，所以这里访问的是 p3.abc 字段。
	p3.address2_1_2.abc = "abc2" // 也可以通过 address2_1_2 来访问 abc 字段。

	fmt.Printf("p3: %#v \n", p3) // p3: main.person2_3{username:"sky", password:"123", address2_1:main.address2_1{name1:"浦东", phone1:"1234567", city1:"上海"}, address2_1_2:main.address2_1_2{name2:"大兴", phone2:"", city2:"北京", abc:"abc2", addTime:""}, abc:"abc", email2_1:main.email2_1{addTime:""}}

	// p3.addTime = "2023-01-01" // 匿名结构体 email2_1 和 address2_1_2 里边 都有 addTime 字段，因此不能直接这样简写。
	p3.email2_1.addTime = "2023-01-01"     // 只能通过 email2_1 来访问 addTime 字段。
	p3.address2_1_2.addTime = "2023-01-02" // 也可以通过 address2_1_2 来访问 addTime 字段。

	fmt.Printf("p3: %#v \n", p3)
	// p3: main.person2_3{
	// 		username:"sky",
	// 		password:"123",
	//  	address2_1:main.address2_1{name1:"浦东", phone1:"1234567", city1:"上海"},
	// 		address2_1_2:main.address2_1_2{name2:"大兴", phone2:"", city2:"北京", abc:"abc2", addTime:"2023-01-02"},
	// 	 	abc:"abc",
	// 		email2_1:main.email2_1{addTime:"2023-01-01"}
	// }
}
