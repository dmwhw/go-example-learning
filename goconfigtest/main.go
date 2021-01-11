package main

import (
	"fmt"
	"reflect"

	"github.com/unknwon/goconfig"
)

func main() {
	c, err := goconfig.LoadConfigFile("./test.conf")
	fmt.Println("error1", err)
	if err == nil {
		value, err := c.Int("", "port")
		fmt.Println(value)
		fmt.Println("error2", err)
	}
	fmt.Println("end...")

	value := c.MustValueArray("", "list", ",")
	fmt.Println(value, "len is", len(value), reflect.TypeOf(value))
	fmt.Println(reflect.TypeOf([5]string{}))
}
