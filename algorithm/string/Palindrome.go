package string

import (
	"math"
	"strconv"
	"strings"
)

/**
palindrome 回文
*/

/**
 给你一个字符串 s，找到 s 中最长的回文子串。
 */
func LongestPalindrome(s string) int {
	var res string
	for i := 0; i < len(s); i++ {
		//偶数情况，以i 为中心
		s1 := palindrome(s, i, i)
		//奇数情况，以i和i+1为中心
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return len(res)
}

//以l和r为中心的字符串s的最长回文串
func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		//向两边展开
		l--
		r++
	}
	return s[l+1 : r]
}

/**
409
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 “Aa” 不能当做一个回文字符串。
*/
func LongestPalindromeAfterBuild(s string) int {
	var sb ['z' - 'A' + 1]int
	for i := range s {
		sb[s[i]-'A']++
	}
	sum := 0
	for _, v := range sb {
		sum += v / 2 * 2
	}

	//sum等于s的长度，说明字符串本身就可以组成一个回文串
	if sum == len(s) {
		return sum
	} else {
		//一定比s小，且是偶数，那偶数成双成对，加1个奇数的就可以了
		return sum + 1
	}
}

/**
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。
*/
func IsPalindromeString(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isChar(s byte) bool {
	if ('0' <= s && s <= '9') || ('A' <= s && s <= 'Z') || ('a' <= s && s <= 'z') {
		return true
	}
	return false
}

func IsPalindromeString2(s string) bool {
	var newS string
	var asc0 rune = '0'
	var asc9 rune = '9'
	var ascA rune = 'A'
	var ascZ rune = 'Z'
	var asca rune = 'a'
	var ascz rune = 'z'

	for _, item := range s {
		if (item >= asc0 && item <= asc9) ||
			(item >= ascA && item <= ascZ) ||
			(item >= asca && item <= ascz) {
			newS += strings.ToLower(string(item))
		}
	}
	newSReverse := ReverseStringWithString(newS)
	return newS == newSReverse
}

/**
 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/
func IsPalindromeInt(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}
	arr := []int{}
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	for i, j := 0, len(arr)-1; i <= j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//转字符串判断
func IsPalindromeInt2(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	strX := strconv.Itoa(x)
	for i := 0; i <= len(strX)/2; i++ {
		if strX[i] != strX[len(strX)-1-i] {
			return false
		}
	}
	return true
}

/**
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。
回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。
回文串不一定是字典当中的单词。
*/
func CanPermutePalindrome(s string) bool {
	//r := []rune(s)
	sm := make(map[rune]int)
	for _, item := range s {
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

/**
 给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
*/
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
		//奇数
		ans += palindrome(s, i, i)
		//偶数
		ans += palindrome(s, i, i+1)
	}
	return ans
}

/**
 给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。
 请你返回让 s 成为回文串的 最少操作次数 。
「回文串」是正读和反读都相同的字符串。
*/
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
