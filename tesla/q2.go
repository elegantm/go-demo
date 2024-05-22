package main

import "fmt"

func SolutionA(A []int, B []int, N int) int {
	// 创建邻接表
	graph := make(map[int][]int)
	for i := 0; i < len(A); i++ {
		a, b := A[i], B[i]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	var dfs func(node int, visited map[int]bool) int
	dfs = func(node int, visited map[int]bool) int {
		visited[node] = true
		roads := len(graph[node])
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				roads += dfs(neighbor, visited)
			}
		}
		return roads
	}

	maxNetwork := 0
	visited := make(map[int]bool)
	//visited := make(map[int]bool)
	for i := 1; i <= N; i++ {
		size := dfs(i, visited) / 2
		if size > maxNetwork {
			maxNetwork = size
		}
		// reset visited
		visited = make(map[int]bool)
	}

	return maxNetwork
}

func mainC() {
	A := []int{1, 1, 2, 2, 3, 3}
	B := []int{2, 4, 4, 3, 4, 5}
	N := 5

	//A := []int{1, 2, 4, 5}
	//B := []int{2, 3, 5, 6}
	//N := 6
	fmt.Println(SolutionA(A, B, N)) // 输出 3
}

//select  name AS UserName,'Creadted'AS type ,CreatedBy AS by , COUNT(*) AS NoOfCreatedRoles FROM UserRole group by Name,CreatedBy
//
//-- UNION ALL
//
//-- select  name AS UserName,'Creadted'AS type ,CreatedBy AS by , COUNT(*) AS NoOfCreatedRoles FROM UserRole  where isEnable =1  group by Name,CreatedBy
//-- UNION ALL
//-- select  name AS UserName,'Updated'AS type ,UpdatedBy AS by , COUNT(*) AS NoOfUpdatedRoles FROM UserRole group by Name,UpdatedBy
