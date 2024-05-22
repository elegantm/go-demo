package main

import (
	"fmt"
	"log"
)

func main() {
	c := testA()
	log.Println("main", c)
	log.Println("main", c)
	log.Println("a ptr", &c)

	//fmt.Println("b return:", b()) // 打印结果为 b return: 2
}
func test() int {
	var a int
	a = 1
	log.Println(a)
	log.Println("a ptr", &a)
	//卵用没有，返回值并不会被接收
	defer func() int {
		a = 4
		log.Println("defer", a)
		log.Println("a ptr", &a)
		return a
	}()

	defer func() int {
		a = 3
		log.Println("defer", a)
		log.Println("a ptr", &a)
		return a
	}()

	panic("test")

	return a
}

// 具名返回值
func testA() (a int) {
	a = 1
	log.Println(a)
	log.Println("a ptr", &a)
	//卵用没有，返回值并不会被接收

	defer func() {
		a = 3
		log.Println("defer", a)
		log.Println("a ptr", &a)

	}()
	defer func() {
		a = 4
		log.Println("defer", a)

	}()

	defer func() {
		a = 5
		log.Println("defer", a)

	}()
	defer func() {
		a = 2
		log.Println("defer", a)
	}()

	log.Println("a ptr", &a)
	return a
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}
