/*
 * @Author: your name
 * @Date: 2020-08-14 13:48:23
 * @LastEditTime: 2020-08-14 14:08:40
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \src\packageMgr\goMod\utils\print.go
 */
package utils
import "fmt"
///大写才能外部调用

///这个不可用。这个文件不能被调用，除非Main.go放到GOPATH
func Print(){
	fmt.Println("hellow")
}