package main

func restoreIpAddresses(s string) []string {
	// 长度小于4或大于12的字符串不可能形成有效 IP 地址
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	var result []string
	// 尝试第1段的长度（1到3位）
	for i := 1; i <= 3 && i <= len(s); i++ {
		// 第1段
		part1 := s[:i]
		if !isValidSegment(part1) {
			continue
		}
		// 剩余部分长度不能超过9（3段，每段最多3位）
		if len(s)-i > 9 {
			continue
		}
		// 尝试第2段的长度
		for j := i + 1; j <= i+3 && j <= len(s); j++ {
			part2 := s[i:j]
			if !isValidSegment(part2) {
				continue
			}
			// 剩余部分长度不能超过6
			if len(s)-j > 6 {
				continue
			}
			// 尝试第3段的长度
			for k := j + 1; k <= j+3 && k <= len(s); k++ {
				part3 := s[j:k]
				if !isValidSegment(part3) {
					continue
				}
				// 剩余部分长度不能超过3
				if len(s)-k > 3 {
					continue
				}
				// 第4段必须刚好用完剩余部分
				part4 := s[k:]
				if len(part4) == 0 || !isValidSegment(part4) {
					continue
				}
				// 拼接成一个有效 IP 地址
				ip := part1 + "." + part2 + "." + part3 + "." + part4
				result = append(result, ip)
			}
		}
	}
	return result
}

// isValidSegment 检查一个字符串是否是有效的 IP 地址段
func isValidSegment(segment string) bool {
	// 长度为0或大于3，无效
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}
	// 如果长度大于1且有前导0，无效
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}
	// 将字符串转换为数字
	num := 0
	for i := 0; i < len(segment); i++ {
		num = num*10 + int(segment[i]-'0')
	}
	// 数字需在0到255之间
	return num >= 0 && num <= 255
}

//果字符串长度大于 12（因为 4 个段，每段最多 3 位，所以最大长度是 12），直接返回空结果。
//逐步分割：
//取第 1 段（1 到 3 位），剩余部分长度不能超过 9（因为后面还有 3 段，每段最多 3 位）。
//取第 2 段（1 到 3 位），剩余部分长度不能超过 6。
//取第 3 段（1 到 3 位），剩余部分长度不能超过 3。
//第 4 段必须刚好用完剩余部分。

func main() {

}
