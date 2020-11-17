package algorithm

func OpenLock(deadends []string, target string) int {
	//死亡数字从数组转成map，判断时加快速度
	deadendsMap := map[string]bool{}
	for _, item := range deadends {
		deadendsMap[item] = true
	}
	//记录已经穷举过的密码
	visited := map[string]bool{}
	//var queue []string
	step := 0
	//queue := list.New()
	//queue.PushFront("0000")
	visited["0000"] = true
	queue := []string{"0000"}
	for len(queue) > 0 {
		//qSize := queue.Len()
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			//cur := queue.Remove(queue.Back()).(string)
			cur := queue[0]
			queue = queue[1:]

			if _, ok := deadendsMap[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}

			for j := 0; j < 4; j++ {
				one := plusOne(cur, j)
				if _, ok := visited[one]; !ok {
					//queue.PushFront(one)
					queue = append(queue, one)
					visited[one] = true
				}
				one = minusOne(cur, j)
				if _, ok := visited[one]; !ok {
					//queue.PushFront(one)
					queue = append(queue, one)
					visited[one] = true
				}
			}
		}
		step++
	}
	return -1
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

func OpenLockByBidirectional(deadends []string, target string) int {
	//死亡数字从数组转成map，判断时加快速度
	deadendsMap := map[string]bool{}
	for _, item := range deadends {
		deadendsMap[item] = true
	}
	//记录已经穷举过的密码
	visited := map[string]bool{}
	visited["0000"] = true
	step := 0
	queueFrom := map[string]bool{"0000":true}
	queueTarget := map[string]bool{target:true}

	for len(queueFrom) > 0 && len(queueTarget) > 0 {
		tmp := map[string]bool{}
		for cur := range queueFrom {
			if _, ok := deadendsMap[cur]; ok {
				continue
			}
			if _, ok := queueTarget[cur]; ok {
				return step
			}
			visited[cur] = true

			for j := 0; j < 4; j++ {
				one := plusOne(cur, j)
				if _, ok := visited[one]; !ok {
					tmp[one] = true
				}
				one = minusOne(cur, j)
				if _, ok := visited[one]; !ok {
					tmp[one] = true
				}
			}
		}
		step++
		queueFrom = queueTarget
		queueTarget = tmp
	}
	return -1
}
