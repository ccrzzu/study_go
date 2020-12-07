package sort

import "fmt"

func QuickSort(a []int) {
	recurSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func recurSort(a []int, left, right int) {
	if left < right {
		partIndex := partition(a, left, right)
		recurSort(a, left, partIndex-1)
		recurSort(a, partIndex+1, right)
	}
}
func partition(a []int, left, right int) int {
	for left < right {
		for left < right && a[left] <= a[right] {
			right--
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
			left++
		}
		for left < right && a[left] <= a[right] {
			left++
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
			right--
		}
	}
	return left
}
