package dynamic_programming

/**
一头母羊的寿命是5年,它会在第2年底和第4年底各生下一头母羊,第5年底死去,
问一开始农场有1头母羊,N年后,农场会有多少只母羊？
*/

//解法一 递归解法 时间复杂度o(logn)
var sum = 1

func TotalSheepAfterYear(n int) int {
	for i := 1; i <= n; i++ {
		if i == 2 || i == 4 {
			sum++
			TotalSheepAfterYear(n - i)
		} else if i == 5 {
			sum--
			break
		}
	}
	return sum
}

// 解法二 迭代解法
type Sheep struct {
	Age int
}

func TotalSheepAfterYear2(n int) int {
	//总羊数
	sheepList := []*Sheep{new(Sheep)}

	for i := 1; i <= n; i++ {
		//待添加的新羊
		addSheepList := make([]*Sheep, 0)
		//待删除的羊
		delSheepList := make([]*Sheep, 0)
		
		for _, item := range sheepList {
			item.Age++
			if item.Age == 2 || item.Age == 4 {
				addSheepList = append(addSheepList, new(Sheep))
			} else if item.Age == 5 {
				delSheepList = append(delSheepList, item)
			}
		}

		sheepList = append(sheepList, addSheepList...)
		for _, delItem := range delSheepList {
			for i, item := range sheepList {
				if item == delItem {
					sheepList = append(sheepList[:i], sheepList[i+1:]...)
					break
				}
			}
		}
	}

	return len(sheepList)
}
