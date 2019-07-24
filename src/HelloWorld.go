package main

import "fmt"

func main() {
	fmt.Println(Hello())
	fmt.Println("Hello World")
}

// 使用func() 创建出一个新的函数
// 定义时增加了关键字：string，
// 意味着这个函数返回一个字符串
func Hello() string {
	return "Hello World"
}
