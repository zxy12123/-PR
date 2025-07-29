package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// WordCount 结构体用于存储单词及其出现次数
type WordCount struct {
	Word  string
	Count int
}

// WordCountList 用于排序的切片类型
type WordCountList []WordCount

// 实现 sort.Interface 接口
func (w WordCountList) Len() int           { return len(w) }
func (w WordCountList) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w WordCountList) Less(i, j int) bool { return w[i].Count > w[j].Count }

// 统计词频的主函数
func countWords(text string) map[string]int {
	// 将文本转换为小写
	text = strings.ToLower(text)

	// 使用正则表达式分割单词，只保留字母和数字
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	words := re.FindAllString(text, -1)

	// 统计每个单词的出现次数
	wordCount := make(map[string]int)
	for _, word := range words {
		if len(word) > 0 {
			wordCount[word]++
		}
	}

	return wordCount
}

// 将词频统计结果转换为可排序的切片
func convertToSlice(wordCount map[string]int) WordCountList {
	var result WordCountList
	for word, count := range wordCount {
		result = append(result, WordCount{Word: word, Count: count})
	}
	return result
}

// 打印词频统计结果
func printWordCount(wordCountList WordCountList) {
	fmt.Println("\n词频统计结果:")
	fmt.Println("================")

	if len(wordCountList) == 0 {
		fmt.Println("没有找到任何单词")
		return
	}

	// 按出现次数降序排列
	sort.Sort(wordCountList)

	// 打印结果
	for i, item := range wordCountList {
		fmt.Printf("%d. %s: %d次\n", i+1, item.Word, item.Count)
	}
}

func main() {
	fmt.Println("CLI 词频统计器")
	fmt.Println("================")
	fmt.Println("请输入要统计的文本 (按 Ctrl+C 退出):")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		// 检查是否为空输入
		if strings.TrimSpace(input) == "" {
			fmt.Println("请输入一些文本进行统计")
			continue
		}

		// 统计词频
		wordCount := countWords(input)
		wordCountList := convertToSlice(wordCount)

		// 打印结果
		printWordCount(wordCountList)
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入时发生错误: %v\n", err)
	}
}
