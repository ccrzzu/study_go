package array

//合并两个排序的数组，合并B到A，形成一个新的有序数组
func MergeB2A(A []int, m int, B []int, n int) {
	var res []int
	var i, j int
	for i < m && j < n {
		if A[i] <= B[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, B[j])
			j++
		}
		//fmt.Println(res)
	}
	if i < m {
		res = append(res, A[i:m]...)
	}
	if j < n {
		res = append(res, B[j:n]...)
	}
	for i, item := range res {
		A[i] = item
	}
	// copy方式如下
	// copy(A, res)
}

/*
 *看到是排好序的数组，则考虑用双指针策略
 *合并B数组到A数组，原地合并
 *这个方法不需要申请新的空间，相较与上个方法省了空间，空间复杂度为O(1)
 */
func MergeB2AWithNoNewMem(A []int, m int, B []int, n int) {
	first := m - 1
	second := n - 1

	for i := m + n - 1; i >= 0; i-- {
		// B has already been merged.
		if second < 0 {
			break
		}
		if first >= 0 && A[first] > B[second] {
			A[i] = A[first]
			first--
		} else {
			A[i] = B[second]
			second--
		}
	}
}

//归并排序逻辑备忘
func MergeSort(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, tail := m-1, n-1, m+n-1
	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[tail] = nums1[i1]
			i1--
		} else {
			nums1[tail] = nums2[i2]
			i2--
		}
		tail--
	}
	for tail >= 0 && i1 >= 0 {
		nums1[tail] = nums1[i1]
		i1--
		tail--
	}
	for tail >= 0 && i2 >= 0 {
		nums1[tail] = nums2[i2]
		i2--
		tail--
	}
}
