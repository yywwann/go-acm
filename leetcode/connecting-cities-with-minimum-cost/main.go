package connecting_cities_with_minimum_cost

import "sort"

type Node struct {
	x, y, c int
}

func minimumCost(n int, connections [][]int) int {
	id := make([]int, n)
	g := make([]Node, 0, n)
	// 初始化并查集
	for i := 0; i < n; i++ {
		id[i] = i
	}

	for _, connection := range connections {
		g = append(g, Node{
			x: connection[0],
			y: connection[1],
			c: connection[2],
		})
	}

	// 将所有的边按照权重从小到大排序
	sort.Slice(g, func(i, j int) bool {
		return g[i].c < g[j].c
	})

	count, cost := 0, 0
	for _, node := range g {
		if count == n-1 { // 如果已经有 n - 1 条边，说明说有点的点都已经联通
			break
		}

		if find(node.x-1, id) == find(node.y-1, id) { // 会形成环，不需要加入
			continue
		}

		union(node.x-1, node.y-1, id) // 关联两个点，并加入到最小生成树中
		count++
		cost += node.c
	}

	if count != n-1 {
		return -1
	}

	return cost
}

func union(i, j int, id []int) {
	x := find(i, id)
	y := find(j, id)
	if x == y {
		return
	}
	id[x] = y
}

func find(x int, id []int) int {
	if x != id[x] {
		id[x] = find(id[x], id)
	}
	return id[x]
}
