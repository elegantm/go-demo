package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("enter sync map")

	var syncMap sync.Map

	syncMap.Store("abc", 11)
	syncMap.Store("xyz", 33)
	v, ok := syncMap.Load("a")
	fmt.Println(ok, v)

	fmt.Println("----")
	syncMap.Range(func(key, value any) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})

	syncMap.Delete("abc")

	syncMap.LoadOrStore("xyz", 99)
	v, ok = syncMap.Load("xyz")
	fmt.Println(ok, v)

}

func testMap() {
	m := make(map[int]string)
	m[1] = "100"
}
