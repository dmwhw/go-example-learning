/*
 * @Author: your name
 * @Date: 2020-08-13 18:42:57
 * @LastEditTime: 2020-08-14 11:58:07
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \src\packageMgr\gopathInSide\src\myapp\myapp.go
 */
package main

/*
一定要有main包，否则不是可执行程序
**/
import (
	"fmt"
	"time"
	"github.com/garyburd/redigo/redis"
)

func main() {
	fmt.Println("hello")
	main2()
}

func main2() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
