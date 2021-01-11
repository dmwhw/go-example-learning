package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type DTO struct {
	id        int
	name      string
	No        string //只有大写开头的才能被导出
	class     string
	NullArray []string
}

type DTOAlias struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	No        string
	class     string `json:"class"` //只有大写开头的才能被导出
	Password  string `json:"password"`
	NullArray []string
}

func main() {

	fmt.Println(os.Args)
	jsonstr := `{"id":1,"name":"ddd","No":"SN-1023","class":"class-2","password":"1234"}`
	dto := stringToDTO(jsonstr)
	fmt.Printf("%+v\r\n", dto)
	fmt.Printf("============Alias===================\r\n")

	dtoa := stringToDTOAlias(jsonstr)
	fmt.Printf(" alias...%+v\r\n", dtoa)
	fmt.Println("dtoa.NullArray-->", dtoa.NullArray)

	fmt.Printf("============json2str===================\r\n")
	databytes, _ := json.Marshal(dtoa)
	fmt.Println("json2 databytes", databytes)
	fmt.Println("json2 str", string(databytes))
	fmt.Printf("============unknownJson===================\r\n")
	f := unknownJson(jsonstr)
	fmt.Println("----unknownJson=-->", f)
}

func unknownJson(str string) (f interface{}) {
	json.Unmarshal([]byte(str), &f)
	fmt.Println("unknow json data--->", f)
	//使用interface数据 .()转换     interface数据"点"(类型)
	m := f.(map[string]interface{})
	fmt.Println("unknow json data--->", m)
	return f
}

func stringToDTO(str string) DTO {
	var dto DTO
	err := json.Unmarshal([]byte(str), &dto)
	fmt.Println("error is ", err)
	fmt.Printf("parse...%+v\r\n", dto)

	return dto
}
func stringToDTOAlias(str string) DTOAlias {
	var dto DTOAlias
	err := json.Unmarshal([]byte(str), &dto)
	fmt.Println("error is ", err)
	fmt.Printf("stringToDTOAlias parse...%+v\r\n", dto)

	return dto
}
