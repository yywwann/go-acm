package number_of_ways_to_reconstruct_a_tree

func checkWays(pairs [][]int) int {
	vec := make(map[int]map[int]bool, 0)
	for _, pair := range pairs {
		x := pair[0]
		y := pair[1]
		if vec[x] == nil {
			vec[x] = make(map[int]bool, 0)
		}
		if vec[y] == nil {
			vec[y] = make(map[int]bool, 0)
		}
		vec[x][y] = true
		vec[y][x] = true
	}

	root := -1
	for x, ys := range vec {
		if len(ys) == len(vec)-1 {
			root = x
			break
		}
	}

	if root == -1 {
		return 0
	}

	ans := 1
	for x, ys := range vec {
		if x == root {
			continue
		}

		// 找到x的父节点
		parentNode := -1
		parentDeg := 500
		for y := range ys {
			if len(vec[y]) >= len(vec[x]) && len(vec[y]) < parentDeg {
				parentDeg = len(vec[y])
				parentNode = y
			}
		}

		if parentNode == -1 {
			return 0
		}

		// 检验是否合法
		// ys 属于 vec[parentNode]
		for y := range ys {
			if y != parentNode && vec[parentNode][y] != true {
				return 0
			}
		}

		// 检验树是否唯一
		if len(ys) == len(vec[parentNode]) {
			ans = 2
		}

	}

	return ans
}
