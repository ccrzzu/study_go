package string


//大数相加，两个字符串数字相加，返回string结果
func AddStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var result string
	var carry, digit int
	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || carry > 0; {
		digit = carry
		if i >= 0 {
			digit += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			digit += int(num2[j] - '0')
			j--
		}
		if digit >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		result = strconv.Itoa(digit%10) + result
	}
	return result
}