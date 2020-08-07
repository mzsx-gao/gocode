package main

import (
	"fmt"
	"strings"
)

func main() {

	//Contains(被查找字符串，查找字符串) ，一般用作于模糊查找
	str1 := "hello world"
	str2 := "llo"
	b := strings.Contains(str1, str2)
	if b {
		fmt.Println("找到了")
	} else {
		fmt.Println("未找到")
	}

	//查找一个字符串在另外一个字符串中第一次出现的位置 返回值为int类型是下标
	/*str1 := "hello worllod"
	str2 := "llo"
	i := strings.Index(str1, str2)
	fmt.Println(i)*/

	//将一个字符串重复n次
	/*str := "性感法师在线讲课"
	str1 := strings.Repeat(str, 3)
	fmt.Println(str1)*/

	//字符串替换  用作于屏蔽敏感词汇,最后一个参数小于0 表示全部替换
	/*str := "性感法师在线性感讲课"
	str1 := strings.Replace(str, "性感", "**", -1)
	fmt.Println(str1)*/

	//将字符串按照标志位进行切割变成切片,切片通过strings.Join方法变成字符串
	/*str1 := "123456@qq.com"
	slice := strings.Split(str1, "@")
	fmt.Println(slice)
	fmt.Println(strings.Join(slice,""))*/

	//去掉字符串头尾内容
	/*str := "=====a=u=ok========"
	str1 := strings.Trim(str, "=")
	fmt.Printf("%s", str1)*/

	//去掉字符串中的空格转成切片  一般用于做统计单词个数
	/*str := "    are you ok     "
	slice := strings.Fields(str)
	//fmt.Println(slice)
	//for _,v:=range slice{
	//	fmt.Printf("=%s=\n",v)
	//}
	fmt.Println(len(slice))*/

}
