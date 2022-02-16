package two_sum

import (
	"sort"
)

type Num struct {
	idx int
	val int
}

func twoSum(nums []int, target int) []int {
	Nums := make([]Num, 0, len(nums))
	for i, num := range nums {
		Nums = append(Nums, Num{
			idx: i,
			val: num,
		})
	}

	sort.Slice(Nums, func(i, j int) bool {
		return Nums[i].val < Nums[j].val
	})

	// return twoPoints(Nums, target)
	// return binarySearch(Nums, target)
	return Hash(Nums, target)
}

// 双指针O(n)
func twoPoints(Nums []Num, target int) []int {
	ans := make([]int, 0, 2)
	i, j := 0, len(Nums)-1
	for i < j {
		x := Nums[i].val
		y := Nums[j].val
		if x+y < target {
			i++
		} else if x+y > target {
			j--
		} else {
			ans = append(ans, Nums[i].idx, Nums[j].idx)
			return ans
		}
	}

	return nil
}

// 二分O(nlogn)
func binarySearch(Nums []Num, target int) []int {
	ans := make([]int, 0, 2)
	for i, num := range Nums {
		x := target - num.val
		l := i + 1
		r := len(Nums) - 1
		t := -1
		for l <= r {
			mid := (l + r) / 2
			if Nums[mid].val <= x {
				t = mid
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		if t != -1 && Nums[t].val == x {
			ans = append(ans, Nums[i].idx, Nums[t].idx)
			return ans
		}
	}

	return nil
}

// 哈希 O(n)
func Hash(Nums []Num, target int) []int {
	ans := make([]int, 0, 2)
	hashMap := make(map[int]int, len(Nums))
	for _, num := range Nums {
		idx := num.idx
		val := num.val
		if t, ok := hashMap[target-val]; ok {
			ans = append(ans, t, idx)
			return ans
		} else {
			hashMap[val] = idx
		}
	}

	return nil
}
