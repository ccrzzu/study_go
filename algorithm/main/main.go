package main

import (
	"fmt"
)

func main() {
	//fmt.Println("CoinChangeByDG:", algorithm.CoinChange([]int{186, 419, 83, 408}, 6249))
	//fmt.Println("CoinChangeByIter:", algorithm.CoinChangeByIter([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(dynamic_programming.Fibonacci(92))
	//fmt.Println(algorithm.FindNumberOfLIS([]int{1,2,4,3,5,4,7,2}))
	//fmt.Println(algorithm.LengthOfLIS([]int{1,2,4,3,5,4,7,2}))
	//fmt.Println(algorithm.SplitToString([]int{1,2,3,4,5,6,7},","))
	//fmt.Println(algorithm.FindSubsequences([]int{1,2,3,4,5,6,7}))
	//fmt.Println(algorithm.FindLHS([]int{1,3,3,4,5,6,7}))
	//fmt.Println(algorithm.MinDistance("horse", "ros"))
	//fmt.Println(algorithm.LongestCommonSubsequence("abcde","ace"))
	//fmt.Println(algorithm.LongestPalindrome("abccccdezzz"))
	//fmt.Println(algorithm.LongestPalindromeAfterBuild("abccccdezzz"))
	//fmt.Println(algorithm.StringIsPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(algorithm.CanPermutePalindrome("aaabbbbcccccaa"))
	//fmt.Println(algorithm.CountSubstrings("aaa"))
	//fmt.Println(algorithm.MinInsertions("mbadm"))
	//fmt.Println(algorithm.Permutations([]int{5,4,6,2}))
	//fmt.Println(algorithm.SolveNQueens(1))
	//fmt.Println(algorithm.BuildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7}))
	//fmt.Println(algorithm.OpenLock([]string{"1131","1303","3113","0132","1301","1303","2200","0232","0020","2223"},"3312"))
	//fmt.Println(tree.OpenLockByBidirectional([]string{"1131","1303","3113","0132","1301","1303","2200","0232","0020","2223"},"3312"))
	//fmt.Println(algorithm.MaxProfit_k_2([]int{3,3,5,0,0,3,1,4}))
	//fmt.Println(dynamic_programming.MaxProfit_k_s(2,[]int{2,4,1}))
	//fmt.Println(double_pointer.TwoSum([]int{3,2,4},6))
	//fmt.Println(dynamic_programming.SuperEggDrop(2,6))
	//fmt.Println(array.ArrSearch([]int{3, 2, 1}, 1))
	//sort.QuickSort([]int{1,2,2,6,5,4,3,3})
	//sort.BubbleSort([]int{1, 2, 2, 6, 5, 4, 3, 3})
	//fmt.Println(stack.SumSubarrayMins([]int{3, 1, 2, 4}))
	//fmt.Println(dynamic_programming.ClimbStairsByDynamicProgram(90))
	//a := number{1}
	//a.print()
	//fmt.Println(a)
	/*gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}*/
	readChannel := make(chan<- int)
	close(readChannel)
}

type number struct {
	Age int
}

func (n number) print() {
	n.Age = 2
	fmt.Println(n)
}
func (n *number) pprint() {
	n.Age = 3
	fmt.Println(*n)
}

func f() int {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
