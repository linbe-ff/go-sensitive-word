package sensitive_word

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

// 定义结构体来映射 JSON 数据
type Data struct {
	IllegalKeywords []string `json:"illegalKeywords"`
}

var (
	dfa       *DFA
	benchData Data
	text      = strings.Repeat("今夜总会想起你夜总", 10) + "最淫官员"
)

func init() {
	dfa = NewDFA()
	// 读取文件
	file, err := os.ReadFile("illegalWords.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 将 JSON 数据解析到结构体中
	err = json.Unmarshal(file, &benchData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// 添加敏感词
	for _, keyword := range benchData.IllegalKeywords {
		dfa.AddWord(keyword)
	}
}

func BenchmarkDFAFilterAll(b *testing.B) {
	// 输入文本
	err := dfa.Check(text, true)
	if err == nil {
		b.Error("期望得到错误，但实际上没有收到")
	} else {
		fmt.Println("捕获到错误:", err)
	}
	//cmd: go test -bench=BenchmarkDFAFilterAll
	//BenchmarkDFAFilterAll-12        捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//1000000000               0.0000965 ns/op
	//PASS
	//ok      sensitive_word  0.186s
}

func BenchmarkDFAFilterForr(b *testing.B) {
	// 添加敏感词
	for _, keyword := range benchData.IllegalKeywords {
		if strings.Contains(text, keyword) {
			fmt.Println("存在敏感词", keyword)
			return
		}
	}
	// cmd: go test -bench=BenchmarkDFAFilterForr
	//BenchmarkDFAFilterAll-12        捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//捕获到错误: 包含敏感词: 最淫官员
	//1000000000               0.0000965 ns/op
	//PASS
	//ok      sensitive_word  0.186s
}

func BenchmarkDFAFilterReplace(b *testing.B) {
	//dfa.Filter(text)
	fmt.Println(dfa.Filter(text, true))
}
