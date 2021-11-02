package main

import "fmt"

func main() {
	// 示例1。
	//value3 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	// switch语句不允许case表达式中的子表达式结果值存在相等的情况，不论这些结果值相等的子表达式，是否存在于不同的case表达式中
	//switch value3[4] { // 这条语句无法编译通过。
	//case 0, 1, 2:
	//	fmt.Println("0 or 1 or 2")
	//case 2, 3, 4:
	//	fmt.Println("2 or 3 or 4")
	//case 4, 5, 6:
	//	fmt.Println("4 or 5 or 6")
	//}

	// 示例2。
	value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or26")
	}

	// 示例3。
	//value6 := interface{}(byte(127))
	// 示例2这种绕过方式对用于类型判断的switch语句（以下简称为类型switch语句）就无效了。
	// 因为类型switch语句中的case表达式的子表达式，都必须直接由类型字面量表示，而无法通过间接的方式表示
	// byte类型是uint8类型的别名类型，因此，它们两个本质上是同一个类型，只是类型名称不同罢了。
	//switch t := value6.(type) { // 这条语句无法编译通过。
	//case uint8, uint16:
	//	fmt.Println("uint8 or uint16")
	//case byte:
	//	fmt.Printf("byte")
	//default:
	//	fmt.Printf("unsupported type: %T", t)
	//}
}
