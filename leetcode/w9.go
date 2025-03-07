package main

import "strings"

// 给你一个字符串数组 words ，数组中的每个字符串都可以看作是一个单词。请你按 任意 顺序返回 words 中是其他单词的 子字符串 的所有单词。
func stringMatching(words []string) []string {
	resultmap := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[i], words[j]) {
				{
					resultmap[words[i]] = true
				}
			}
		}
	}
	result := []string{}
	for word := range resultmap {
		result = append(result, word)
	}
	return result
}

// 最长的一定不可能是子字符串，则只需要每次将当前字符串当中的最长的和其他的匹配，只要存在子串，则输出
// 当一轮结束之后将当前这个最大的字符串排出 当前字符串数组
func main() {

}
