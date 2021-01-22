package double_pointer


//反转字符串
func reverseString(s []byte)  {
	left,right := 0, len(s)-1
	for left < right{
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}
}