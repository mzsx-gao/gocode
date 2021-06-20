package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func main() {
	encode()
	decode()
}

func encode() {
	p := Person{"luhan", "男"}
	//1.生成JSON文本
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("对象转json字符串...")
	fmt.Println(string(b))
	//2.生成格式化json，没啥用
	fmt.Println("格式化json字符串...")
	b, err = json.MarshalIndent(p, "", "    ")
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func decode() {
	//1.准备一段json
	b := []byte(`{"Name":"luhan","Hobby":"男"}`)
	//2.声明结构体
	var p Person
	//3.解析
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println("json字符串转对象...")
	fmt.Println(p)
}
