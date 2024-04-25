package main

import "fmt"

func main() {
	question := [3][2]int{{1, 2}, {1, 3}, {2, 3}}

	edges := [3][3]int{}
	visit := [3]bool{false, false, false}
	hasCircle := false
	fmt.Println("初始化： ", edges, visit, question)

	for _, value := range question {
		// 无向, 加两次
		edges = addEdges(edges, value[0], value[1])
		edges = addEdges(edges, value[1], value[0])

		for i := 0; i < 3; i++ {
			visit[i] = false
		}

		// 检查是否有环
		hasCircle = find(1, -1, edges, visit)
		if hasCircle {
			fmt.Println("有环了: ", value)
			return
		}
	}

	fmt.Println("加边后: ", edges)
}

func find(now int, father int, edges [3][3]int, visit [3]bool) bool {
	// 第一步: 标记走过
	visit[now] = true

	for _, value := range edges[now] {
		if value == father {
			continue
		}

		if visit[value] {
			return true
		}

		return find(value, now, edges, visit)
	}

	return false
}

func addEdges(question [3][3]int, father int, son int) [3][3]int {
	question[father-1][son-1] = son
	return question
}
