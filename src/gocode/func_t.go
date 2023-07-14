package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func (receiver Person) AddAge(num int) {
	fmt.Println(receiver.Age + num)
}

func reflectDemo(i interface{}) {

	t := reflect.TypeOf(i)
	fmt.Println(t)
	fmt.Printf("%T\n", t) //*reflect.rtype

	//t := t.Kind() // 需要实现reflect.Type

	v := reflect.ValueOf(i)
	fmt.Println(v)
	fmt.Printf("%T\n", v) //reflect.Value
	k := v.Kind()
	fmt.Println(k)
	fmt.Printf("%T\n", k) //reflect.Value

	//反射属性和方法
	//修改属性
	fmt.Println(v.NumField(), v.NumMethod())
	//v.Field(0).SetString("李四")
	//v.Elem().NumField()

	//调用方法
	v.Method(0).Call([]reflect.Value{reflect.ValueOf(1)})
	v.MethodByName("AddAge").Call([]reflect.Value{reflect.ValueOf(1)})

}

func main() {
	reflectDemo(Person{Name: "张三", Age: 18, Gender: "男"})
}

//结构体反射只能反射公有方法
//结构体方法是按照ASCII顺序排序的
//反射中所有的操作都是基于reflect.Value进行操作的
