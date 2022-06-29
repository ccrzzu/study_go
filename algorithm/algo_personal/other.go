package main

import "fmt"

func main() {
	s1 := []int{1, 2, 4, 5, 7, 9}
	s2 := []int{1, 2, 5, 6, 9}
	fmt.Println(diff(s1, s2))
	fmt.Println(binarySearch(s1,9))
}

func intersec(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}

func diff(nums1, nums2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	insec := intersec(nums1, nums2)
	for _, v := range insec {
		m[v]++
	}
	for _, value := range nums1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	for _, value := range nums2 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}

	return nn
}

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
