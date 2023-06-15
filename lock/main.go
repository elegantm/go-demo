/*
@Time : 2022/7/29 10:53
@Author : elegantm
@File : main
@Software: gtm for ddi
//Copyright (c) 2003-2021 xxx
//All rights reserved.
功能介绍：
修订历史:
*/
package main

import (
	"fmt"
	"sync"
)

type Obj struct {
	mu *sync.Mutex
	// ... 其他字段
}

func (o Obj) Lock() {
	fmt.Printf("lock obj address  %p \n", o.mu)
	o.mu.Lock()

}
func (o Obj) Dosomething() {}
func (o Obj) Unlock() {
	fmt.Printf("unlock obj address  %p \n", o.mu)
	o.mu.Unlock()
}

func main() {
	o := Obj{
		mu: &sync.Mutex{},
	}
	fmt.Printf("obj address  %p \n", o.mu)
	o.Lock()
	o.Dosomething()
	o.Unlock()

	//o.Lock()
	//o.Dosomething()
	//o.Unlock()
}
