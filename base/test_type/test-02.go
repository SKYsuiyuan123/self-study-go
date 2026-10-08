package main

import "fmt"

// 类型别名 与 类型定义 区别

// 1. 类型别名 (type alias) 是给一个已有的类型起一个新的名字，它们是同一个类型，互相可以赋值。
// 2. 类型定义 (type definition) 是创建一个新的类型，它们是不同的类型，不能直接赋值。
// 3. 类型别名通常用于简化类型的使用，或者在大型项目中重命名类型以提高可读性。
// 4. 类型定义 可以绑定方法，类型别名不可以绑定方法。（因为别名就是原类型，你不能给 原始类型 加方法
// 5. 类型定义 和类型别名在类型检查时是不同的类型，不能互相赋值。
// 6. 类型定义 和类型别名在编译时的表现不同，类型定义 会生成新的类型，而类型别名只是给已有类型起一个新的名字。
// 7. 类型定义 和类型别名在使用时的语法不同，类型定义 需要使用类型转换，而类型别名可以直接使用。
// 8. 类型别名 打印类型时 还是原始类型，类型定义  打印类型时 是类型定义 。
// 9. 和原始类型比较，类型别名不需要转换，但是 类型定义 需要转换。

// 类型别名主要是为了渐进式重构和迁移而设计的
// byte / rune 就是类型别名，byte 是 uint8 的别名，rune 是 int32 的别名。

// 类型定义 (type definition) (new type)
type myCode int

// 类型别名 (type alias)
type myAliasCode = int

// 自定义类型绑定方法
func (m myCode) printMyCode() {
	fmt.Println("mycode ==>", m)
}

// 类型别名不可以绑定方法
// func (m myAliasCode) printMyAliasCode() {
// 	fmt.Println("myAliasCode ==>", m)
// }

func main() {
	const myCode1 myCode = 1
	const myAliasCode1 myAliasCode = 2

	fmt.Printf("myCode1: %v, type: %T \n", myCode1, myCode1)                // myCode1: 1, type: main.myCode
	fmt.Printf("myAliasCode1: %v, type: %T \n", myAliasCode1, myAliasCode1) // myAliasCode1: 2, type: int

	if myCode1 == 1 {
		fmt.Println("myCode1 == 1") // myCode1 == 1
	}

	// 1. 类型比较

	var age = 1

	// 1.1 自定义类型 不能和基本类型进行比较

	// if myCode1 == age {
	// 	fmt.Println("myCode1 == age") // myCode1 == age
	// }

	// 1.2 自定义类型 必须使用类型转换 才能和原始类型进行比较。

	if int(myCode1) == age {
		fmt.Printf("myCode1 == age ==> %v \n", true) // myCode1 == age ==> true
	}

	// 1.3 或者把基本类型转换成 自定义类型

	if myCode1 == myCode(age) {
		fmt.Printf("myCode1 == myCode(age) ==> %v \n", true) // myCode1 == myCode(age) ==> true
	}

	// 1.4 类型别名 可以和原始类型进行比较

	if myAliasCode1 != age {
		fmt.Printf("myAliasCode1 == age ==> %v \n", false) // myAliasCode1 == age ==> false
	}

}
