package main

import (
	"fmt"
	"unsafe"
)

type MyBase struct {
	age       int    //只有大写开头的才能被导出
	No        int32  //只有大写开头的才能被导出
	birthTime int64  //只有大写开头的才能被导出
	isAdmin   bool   //只有大写开头的才能被导出
	Name      string //只有大写开头的才能被导出
}

func main() {
	var my *MyBase
	fmt.Println("var my *MyBase is Nil", my)
	my = new(MyBase)
	fmt.Println("var my *MyBase [new(MyBase)] ,default value is--->", my)
	fmt.Println("my.Name-->", my.Name)
	fmt.Println("my.Name ==\"\"-->", my.Name == "")

	fmt.Println("struct sizes--->", unsafe.Sizeof(my))
	/////////////////
	fmt.Println("==========create default mybase myBase=============")
	myBase := MyBase{
		age:       1,
		No:        333,
		birthTime: 159,
		isAdmin:   false,
		//Name:      "name",
	}
	fmt.Println("create default mybase-->", myBase)
	fmt.Println("myBase.age-->", myBase.age)
	fmt.Println("myBase.Name-->", myBase.Name)
	fmt.Println("struct sizes--->", unsafe.Sizeof(myBase))

}
