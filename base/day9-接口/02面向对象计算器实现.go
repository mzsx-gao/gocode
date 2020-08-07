package main

import "fmt"

//定义接口
type Opter interface {
	//方法声明
	result() int
}

//父类
type Operate struct {
	num1 int
	num2 int
}

//加法子类
type Add struct {
	Operate
}

//减法子类
type Sub struct {
	Operate
}

//乘法子类
type Mlt struct {
	Operate
}

//工厂类
type Factory struct {
}

//加法子类的方法
func (a *Add) result() int {
	return a.num1 + a.num2
}

//减法子类的方法
func (s *Sub) result() int {
	return s.num1 - s.num2
}

//乘法子类的方法
func (m *Mlt) result() int {
	return m.num1 * m.num2
}

//多态实现
func result(o Opter) {
	value := o.result()
	fmt.Println(value)
}

func (f *Factory) result(num1 int, num2 int, ch string) {
	switch ch {
	case "+":
		var a Add
		a.num1 = num1
		a.num2 = num2
		result(&a)
	case "-":
		var s Sub
		s.num1 = num1
		s.num2 = num2
		result(&s)
	case "*":
		var m Mlt
		m.num1 = num1
		m.num2 = num2
		result(&m)
	}
}

//通过设计模式调用
func main() {
	//创建工厂对象
	var f Factory
	f.result(10, 20, "*")
}

//多态调用
func main0203() {
	//var a Add = Add{Operate{10, 20}}
	var s Sub = Sub{Operate{10, 20}}
	result(&s)
}

//接口调用
func main0202() {
	var o Opter
	var a Add = Add{Operate{10, 20}}

	o = &a
	value := o.result()
	fmt.Println(value)

}

//方法调用
func main0201() {
	//创建加法对象
	//var a Add
	//a.num1 = 10
	//a.num2 = 20
	//v := a.Result()
	//fmt.Println(v)

	//创建减法对象
	var s Sub
	s.num1 = 10
	s.num2 = 20
	v := s.result()
	fmt.Println(v)
}
