package main

import (
	"fmt"
)

func mainB() {
	Solution(12345)
}

func Solution(N int) {
	var enable_print int
	//enable_print = N % 10
	for N > 0 {
		tailNum := N % 10
		if enable_print == 0 && tailNum != 0 {
			enable_print = 1
		} else if enable_print == 1 {
			fmt.Print(tailNum)
		}
		N = N / 10
	}
}

func Solution1(N int) {
	var enable_print int

	for N > 0 {
		num := N % 10
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
		}
		if enable_print == 1 {
			fmt.Print(num)
		}
		N = N / 10
	}
}
