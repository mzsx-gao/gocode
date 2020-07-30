package main

import "fmt"

func demo(i int) {

	//错误拦截到出现在错误之前
	defer func() {
		//错误拦截 panic异常错误
		err := recover()
		//判断是否出现错误
		if err != nil {
			fmt.Print("拦截到错误信息\n")
			fmt.Println(err)
		}
	}()

	var p *int //0x0 nil
	p = new(int)
	*p = 123

	fmt.Print("制造数组下标越界\n")
	//数组下标越界
	var arr [10]int
	arr[i] = 123
	fmt.Println(arr)

}

func main() {
	demo(10)

	fmt.Println("hello world")
}
