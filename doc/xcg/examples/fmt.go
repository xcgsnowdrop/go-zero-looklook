package main

import "fmt"

func main() {
	a := 42
	b := "Hello, Gophers!"
	fmt.Print("a is ", a, ", b is ", b)    // 输出：a is 42, b is Hello, Gophers! (没有换行)
	fmt.Println("a is", a)                 // 输出：a is 42（自动换行）
	fmt.Printf("a is %d, b is %s\n", a, b) // 输出：a is 42, b is Hello, Gophers!

	result := fmt.Sprint("a is ", a, ", b is ", b) // 返回字符串
	fmt.Println(result)
}
