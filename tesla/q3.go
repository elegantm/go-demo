package main

import "fmt"

//func maximalNetworkRank(n int, starts []int, ends []int) int {
//	// Create a map to store the count of roads connected to each city
//	connectedRoads := make(map[int]int)
//
//	// Iterate through all roads and count the connections for each city
//	for i := 0; i < len(starts); i++ {
//		connectedRoads[starts[i]]++
//		connectedRoads[ends[i]]++
//	}
//
//	// Calculate the maximal network rank
//	maxRank := 0
//	for i := 0; i < len(starts); i++ {
//		rank := connectedRoads[starts[i]] + connectedRoads[ends[i]] - 1
//		if rank > maxRank {
//			maxRank = rank
//		}
//	}
//
//	return maxRank
//}

func maximalNetworkRank(n int, starts []int, ends []int) int {
	connectRoads := make(map[int]int)
	for i := 0; i < len(starts); i++ {
		connectRoads[starts[i]]++
		connectRoads[ends[i]]++
	}

	maxRank := 0
	for i := 0; i < len(starts); i++ {
		rank := connectRoads[starts[i]] + connectRoads[ends[i]] - 1
		if rank > maxRank {
			maxRank = rank
		}
	}

	return maxRank
}

func main() {
	//starts := []int{1, 2, 3, 3}
	//ends := []int{2, 3, 1, 4}
	//n := 4

	//starts := []int{1, 1, 2}
	//ends := []int{2, 3, 4}
	//n := 4

	//starts := []int{1, 2, 4, 5}
	//ends := []int{2, 3, 5, 6}
	//n := 6

	starts := []int{1, 1, 2, 2, 3, 3}
	ends := []int{2, 4, 4, 3, 4, 5}
	n := 5

	fmt.Println(maximalNetworkRank(n, starts, ends)) // Output: 4

	A := []int{1, 1, 2, 2, 3, 3}
	B := []int{2, 4, 4, 3, 4, 5}
	N := 5

	//A := []int{1, 2, 4, 5}
	//B := []int{2, 3, 5, 6}
	//N := 6
	fmt.Println(SolutionA(A, B, N)) // 输出 3
}
