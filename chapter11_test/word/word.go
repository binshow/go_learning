package word

import "unicode"

//IsPalindrome 判断一个字符串是不是回文字符串
func IsPalindrome(s string) bool {
	//优化2: 在开始为每个字符预先分配一个足够大的数组，这样就可以避免在append调用时可能会导致内存的多次重新分配
	letters := make([]rune , 0 , len(s))
	for _ , r := range s{
		if unicode.IsLetter(r) {
			letters = append(letters , unicode.ToUpper(r))
		}
	}
	// 优化1：只遍历1半数组
	n := len(letters) / 2
	for i := 0 ; i < n ; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
