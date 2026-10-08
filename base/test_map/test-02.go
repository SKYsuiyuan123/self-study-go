package main

import (
	"fmt"
	"slices"
	"strings"
)

// map 的键值 排序

// map 的 range 顺序不可依赖，需要借助 slice。

// map 排序的 几个坑
// 	1. 别依赖 map 的 range 顺序。它每次运行都可能不同，这不是 bug，是 Go 故意的。
//  2. 排序是 O(n log n)，取 key 是 O(n)。大 map 频繁排序要考虑性能，必要时缓存排序结果。
//  3. key 类型必须可比较（comparable）。slice、map、func 不能做 key，也就无从排序。
//  4. 稳定排序用 slices.SortStableFunc，普通 slices.Sort 不保证相等元素的相对顺序。
//  5. 浮点 key 有 NaN 问题，NaN 参与排序行为特殊，实际很少这么用。

func main() {
	// 1. map 的键值 排序

	map1 := map[int]string{1: "a", 3: "c", 2: "b"}

	fmt.Printf("map1: %#v \n", map1) // map1: map[int]string{1:"a", 2:"b", 3:"c"}

	// map 按照 key 的升序排列 (借助 slice)

	var map1KeySlice1 = make([]int, 0, len(map1))
	// 尽量不要使用 var map1KeySlice1 = []int{}，原因是 在 map 很大的情况下，可能会导致 slice 扩容，浪费内存。

	for key, _ := range map1 {
		map1KeySlice1 = append(map1KeySlice1, key)
	}

	fmt.Printf("map1KeySlice1: %#v \n", map1KeySlice1) // map1KeySlice1: []int{1, 3, 2}

	slices.SortFunc(map1KeySlice1, func(a, b int) int { return a - b })

	fmt.Printf("map1KeySlice1: %#v \n", map1KeySlice1) // map1KeySlice1: []int{1, 2, 3}

	// 排序后的 map

	map2 := map[int]string{}

	for _, item := range map1KeySlice1 {
		map2[item] = map1[item]
	}

	fmt.Printf("map1: %#v \n", map1) // map1: map[int]string{1:"a", 2:"b", 3:"c"}
	fmt.Printf("map2: %#v \n", map2) // map2: map[int]string{1:"a", 2:"b", 3:"c"}

	fmt.Println("----------------分割线---111-----------------")

	const str1 = "how do you do how do you do"

	wordCountMap := countWords(str1)

	fmt.Printf("wordCountMap: %#v \n", wordCountMap) // wordCountMap: map[string]int{"do":4, "how":2, "you":2}
}

// 写一个程序，统计一个字符串中 每个单词出现的次数。比如 "how do you do" 中 how = 1, do = 2, you = 1.

func countWords(s string) map[string]int { // github copilot 生成
	// 1. 将字符串按空格分割成单词切片
	wordsSlice := strings.Fields(s) // Fields 会自动去除多余的空格

	// 2. 创建一个 map 来存储每个单词出现的次数。 (key: 单词, value: 出现次数)
	wordCountMap := make(map[string]int)

	// 3. 遍历单词切片，统计每个单词出现的次数。(利用 map 和 int 的默认值是 0 的特性)
	for _, item := range wordsSlice {
		wordCountMap[item]++
	}

	return wordCountMap
}

func countWords2(s string) map[string]int {
	// 1. 将字符串按空格分割成单词切片
	wordsSlice := strings.Split(s, " ")

	// 2. 创建一个 map 来存储每个单词出现的次数。 (key: 单词, value: 出现次数)

	wordCountMap := make(map[string]int)

	// 3. 遍历单词切片，统计每个单词出现的次数。(利用 map 和 int 的默认值是 0 的特性)

	for _, item := range wordsSlice {
		if item != "" { // 排除空字符串
			wordCountMap[item]++
		}
	}

	return wordCountMap
}
