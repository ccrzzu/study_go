package algorithm

import "container/list"

func OpenLock(deadends []string, target string) int {
	//记录已经穷举过的密码
	var visited []string
	//var queue []string
	step := 0
	queue := list.New()
	queue.PushFront("0000")
	visited = append(visited, "0000")
	for queue.Len() > 0 {
		qSize := queue.Len()
		for i := 0; i < qSize; i++ {
			cur := queue.Remove(queue.Back()).(string)

			if InSliceString(cur, deadends) {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				one := plusOne(cur, j)
				if !InSliceString(one, visited) {
					queue.PushFront(one)
					visited = append(visited, one)
				}
				one = minusOne(cur, j)
				if !InSliceString(one, visited) {
					queue.PushFront(one)
					visited = append(visited, one)
				}
			}
			//queue = SliceUnique(queue)
			//visited = SliceUnique(visited)
		}
		step++
	}
	return -1
}

func InSliceString(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func SliceUnique(slice []string) (uniqueSlice []string) {
	for _, v := range slice {
		if !InSliceString(v, uniqueSlice) {
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	return
}

// 将 s[j] 向上拨动一次
func plusOne(s string, j int) string {
	bytes := []byte(s)
	if bytes[j] == '9' {
		bytes[j] = '0'
	} else {
		bytes[j] += 1
	}
	return string(bytes)
}

// 将 s[i] 向下拨动一次
func minusOne(s string, j int) string {
	bytes := []byte(s)
	if bytes[j] == '0' {
		bytes[j] = '9'
	} else {
		bytes[j] -= 1
	}
	return string(bytes)
}
