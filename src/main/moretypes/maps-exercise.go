package main

import (
	"fmt"
	"strings"
)

//实现 `WordCount`。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。
//你会发现 strings.Fields 很有帮助。
//import (
//	"code.google.com/p/go-tour/wc"
//)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	resultMap := make(map[string]int, len(fields))
	for _, word := range fields {
		i, ok := resultMap[word]
		if !ok {
			resultMap[word] = 1
		} else {
			resultMap[word] = i + 1
		}
	}
	return resultMap
}

func main() {
	countMap := WordCount("Provide you with the most comprehensive and professional translation of professional vocabulary, literature translation, English-Chinese translation, English-Chinese dictionary, Chinese-English dictionary dictionary, English-Chinese examples, professional vocabulary, technical terminology translation and other free translation services")
	fmt.Println(countMap)
}
