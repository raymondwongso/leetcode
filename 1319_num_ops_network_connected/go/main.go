package main

import "fmt"

func main() {
	fmt.Println(makeConnected(4, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	visited := make([]bool, n)
	adj := make([][]int, n)

	for _, c := range connections {
		adj[c[0]] = append(adj[c[0]], c[1])
		adj[c[1]] = append(adj[c[1]], c[0])
	}

	comp := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, adj, visited)
			comp++
		}

	}

	return comp - 1
}

func dfs(index int, adj [][]int, visited []bool) {
	visited[index] = true

	for _, a := range adj[index] {
		if !visited[a] {
			dfs(a, adj, visited)
		}
	}
}
