package algorithm

import (
	"math"
	"strings"
)

func LongestPalindrome(s string) int {
	var res string
	for i := 0; i < len(s); i++ {
		//偶数情况，以i 为中心
		s1 := palindrome(s, i, i)
		//奇数情况，以i和i+1为中心
		s2 := palindrome(s, i, i+1)
		// res == longest(res,s1,s2)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return len(res)
}

//已l和r为中心的字符串s的最长回文串
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		//向两边展开
		l--
		r++
	}
	return s[l+1 : r]
}

func StringIsPalindrome(s string) bool {
	var newS string
	var asc0 rune = '0'
	var asc9 rune = '9'
	var asca rune = 'a'
	var ascz rune = 'z'
	var ascA rune = 'A'
	var ascZ rune = 'Z'
	for _, item := range s {
		if (item >= asc0 && item <= asc9) || (item >= ascA && item <= ascZ) || (item >= asca && item <= ascz) {
			newS += strings.ToLower(string(item))
		}
	}
	/*fmt.Println(newS)
	var i, j int
	if len(newS)%2 == 0 {
		i = len(newS)/2 - 1
		j = i + 1
	} else {
		i = len(newS) / 2
		j = i
	}
	for i>=0 && j<len(newS) {
		if newS[i] == newS[j]{
			i--
			j++
		}else {
			return false
		}
	}*/
	newSReverse := ReverseString(newS)
	return newS == newSReverse
}

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func CanPermutePalindrome(s string) bool {
	r := []rune(s)
	sm := make(map[rune]int)
	for _, item := range r {
		sm[item]++
	}
	var res int
	for _, v := range sm {
		if v%2 != 0 {
			res++
			if res == 2 {
				return false
			}
		}
	}
	return true
}

func CountSubstrings(s string) int {
	//函数嵌套定义
	palindrome := func(s string, l, r int) int {
		var count int
		for l >= 0 && r < len(s) && s[l] == s[r] {
			//向两边展开
			l--
			r++
			count++
		}
		return count
	}
	var ans int
	for i := 0; i < len(s); i++ {
		ans += palindrome(s, i, i)
		ans += palindrome(s, i, i+1)
	}
	return ans
}

func MinInsertions(s string) int {
	slen := len(s)
	dp := make([][]int, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]int, slen)
	}
	for i := slen - 2; i >= 0; i-- {
		for j := i + 1; j < slen; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i+1][j]), float64(dp[i][j-1])) + 1)
			}
		}
	}
	return dp[0][slen-1]
}
