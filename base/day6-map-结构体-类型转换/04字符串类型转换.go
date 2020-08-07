package main

import (
	"fmt"
	"strconv"
)

//将字符串转成字符切片  强制类型转换
func main0401() {
	str := "hello world"
	slice := []byte(str)
	fmt.Println(slice) //[104 101 108 108 111 32 119 111 114 108 100]
	slice[4] = 'a'
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%c", slice[i])
	}
}

//字符切片转字符串
func main0402() {
	slice := []byte{'h', 'e', 'l', 'l', 'o', 97}
	fmt.Println(slice)
	fmt.Println(string(slice))
}

//将其他类型转成字符串  Format
func main0403() {
	b := false
	str := strconv.FormatBool(b)
	fmt.Printf("%T:%s\n", str, str)

	str = strconv.FormatInt(123, 10)
	fmt.Printf("%T:%s\n", str, str)

	str = strconv.FormatFloat(3.14159, 'f', 4, 64)
	fmt.Printf("%T:%s\n", str, str)

	str = strconv.Itoa(123)
	fmt.Printf("%T:%s\n", str, str)
}

//将字符串转成其他类型
func main0404() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("类型转化出错")
	} else {
		fmt.Printf("%T\n", b)
	}

	value, _ := strconv.ParseInt("abc", 16, 64)
	fmt.Printf("%T:%d\n", value, value)

	value2, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("%T:%.5f\n", value2, value2)

	value3, _ := strconv.Atoi("123")
	fmt.Printf("%T:%d\n", value3, value3)
}

//将其他类型转成字符串添加到字符切片里面
func main0405() {

	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, false)
	slice = strconv.AppendInt(slice, 123, 2)

	slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
	slice = strconv.AppendQuote(slice, "hello")

	fmt.Println(string(slice))
}
