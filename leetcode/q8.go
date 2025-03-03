package main

func patternMatching(pattern string, value string) bool {
	// 边界情况
	if len(pattern) == 0 {
		return len(value) == 0
	}
	if len(value) == 0 {
		return len(pattern) <= 1
	}

	// 统计a和b数量
	countA := 0
	countB := 0
	for _, char := range pattern {
		if char == 'a' {
			countA++
		} else {
			countB++
		}
	}

	n := len(value)

	// 处理只有a或只有b的情况
	if countB == 0 {
		if n%countA != 0 {
			return false
		}
		lenA := n / countA
		segment := value[:lenA]
		for i := 0; i < n; i += lenA {
			if value[i:i+lenA] != segment {
				return false
			}
		}
		return true
	}
	if countA == 0 {
		if n%countB != 0 {
			return false
		}
		lenB := n / countB
		segment := value[:lenB]
		for i := 0; i < n; i += lenB {
			if value[i:i+lenB] != segment {
				return false
			}
		}
		return true
	}

	// 计算最大长度
	maxLenA := n / countA
	maxLenB := n / countB

	// 枚举a的长度
	for lenA := 0; lenA <= maxLenA; lenA++ {
		remaining := n - countA*lenA
		if remaining%countB != 0 {
			continue
		}
		lenB := remaining / countB
		if lenB > maxLenB {
			continue
		}

		// 验证
		pos := 0
		var strA, strB string
		matched := true
		for _, char := range pattern {
			if char == 'a' {
				if pos+lenA > n {
					matched = false
					break
				}
				curr := value[pos : pos+lenA]
				if strA == "" {
					strA = curr
				} else if strA != curr {
					matched = false
					break
				}
				pos += lenA
			} else {
				if pos+lenB > n {
					matched = false
					break
				}
				curr := value[pos : pos+lenB]
				if strB == "" {
					strB = curr
				} else if strB != curr {
					matched = false
					break
				}
				pos += lenB
			}
		}
		if matched && pos == n && strA != strB {
			return true
		}
	}

	return false
}

// 在统计出来ab数量之后，直接尝试匹配数量能不能对应上，
// 当a或者b在一个长度可以匹配上时测试另一个
// ，如果某个字符的长度超过value总长度除以对应字符数量
// （即lenA > len(value)/countA），直接返回false。
func main() {

}
