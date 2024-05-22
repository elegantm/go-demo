package main

import "sync"

func syncMap() {
	var m sync.Map

	m.Store("111", 1)

	m.Store("222", 2)
}
