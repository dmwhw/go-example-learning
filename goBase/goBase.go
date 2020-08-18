package main

import (
	"fmt"
	"reflect"
)

var (
	a int
	b int
)

func maps() {
	fmt.Println(" ****************map******************")

	var myMap map[string]interface{}
	fmt.Println("myMap is null?", myMap)
	myMap = make(map[string]interface{})
	fmt.Println("myMap is null?", myMap)
	fmt.Println("myMap type？？？?", reflect.TypeOf(myMap))
	myMap["key"] = 1
	myMap["key2"] = "3333"

	fmt.Println("myMap is null?", myMap)

	myMap["key3"] = nil
	fmt.Println("myMap放nil?", myMap)

	for key, value := range myMap {
		fmt.Println("myMap删除，key-->%s value-->%s", key, value)
		delete(myMap, key)
	}
	fmt.Println("myMap删除，支持遍历删除", myMap)
	fmt.Println("判断有没有key", myMap)
	value, hasKey := myMap["key"]
	if hasKey {
		fmt.Println("判断有没有key hasKey-->?", hasKey, value)
	} else {
		fmt.Println("判断有没有key hasKey-->?", hasKey, value)
	}
}

func arrays() {
	fmt.Println(" *******************数组测试*********************")
	fmt.Println("  需要指定大小，定长 ，值传递")

	var a1 [3]int
	a2 := [...]int{1, 2, 3}
	a3 := [...]int{2: 2, 4: 4} ///按照索引赋值
	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)
	fmt.Println("  object数组")

	d := []interface{}{1, "2", "2", 3.2}
	fmt.Println("dArrays interface{} is ", d, d[2])
	a3[0] = 3333
	fmt.Printf(" a3 address is %p,Type is %v,%v\n", &a3, reflect.TypeOf(a3), a3)
	_setArray(a3) ///实测是值传递

	fmt.Printf(" now a3 address is %p,Type is %v,%v\n", &a3, reflect.TypeOf(a3), a3)

	////////////////////////////////////////////////////
	var a4 []int
	var a5 []int

	fmt.Printf("  a4 address is %p,Type is %v,%v\n", &a4, reflect.TypeOf(a4), a4)
	fmt.Printf("  a5 address is %p,Type is %v,%v\n", &a5, reflect.TypeOf(a5), a5)

	a4 = a5
	fmt.Printf(" now a4 address is %p,Type is %v,%v\n", &a4, reflect.TypeOf(a4), a4)

}

func slicess() {
	fmt.Println(" *******************slice测试*********************")
	fmt.Println("  需要初始化大小，有最大的容量，可以不停增长，每次增长是上次容量的1倍，引用传递")

	var b1 []int
	fmt.Println("b1=", b1)

	b2 := make([]int, 2)
	fmt.Println("len=%d cap=%d", len(b2), cap(b2))

	b3 := make([]int, 2, 4)
	fmt.Printf("len=%d cap=%d\n", len(b3), cap(b3))
	b3[0] = 3333
	fmt.Printf(" b3 address is %p,Type is %v,%v\n", &b3, reflect.TypeOf(b3), b3)

	_setSlice(b3)
	fmt.Printf(" now b3 address is %p,Type is %v,%v\n", &b3, reflect.TypeOf(b3), b3)
}

func _setSlice(s []int) { ///引用传递
	s[0] = 1111

}
func _setArray(s [5]int) { ///实测是值传递
	s[0] = 1111

}

/*
任意类型的不定参数，用interface{}表示
*/
func MyPrintf(args ...interface{}) {
	fmt.Println(" *******************不定参数*********************")

	for _, arg := range args { //迭代不定参数
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is int")
		case string:
			fmt.Println(arg, "is string")
		case float64:
			fmt.Println(arg, "is float64")
		case bool:
			fmt.Println(arg, " is bool")
		default:
			fmt.Println("未知的类型")
		}
	}
}

func returnAll(i int, str string, d interface{} /*未知类型 Object*/) (int, string, interface{}) {
	fmt.Println(" *******************返回多个参数*********************")

	return i, str, d
}

func newFunc() {

}

func makeFunc() {

}

func golbalA() {
	fmt.Println("global a is ", a)
}

func main() {
	a := 1
	b := 2
	var f interface{} = nil ///无论如何都要有一个类型，即使是神秘类型
	fmt.Println(b)

	_, _, f = returnAll(a, "", "ddddd")
	fmt.Println("f is ", f)
	_, _, f = returnAll(a, "", 11111)
	fmt.Println("f is ", f)
	golbalA() //全局变量
	maps()    //map
	arrays()
	slicess()
	MyPrintf(2, "2", 2.5, false)
}
