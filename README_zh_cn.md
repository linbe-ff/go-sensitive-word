# go-sensitive-word
[English](https://github.com/linbe-ff/go_study/blob/main/README.md) | 简体中文
```
具体使用：
    在启动项目的时候，先将数据库中或者文件中的敏感词加载到内存中，然后就可以使用敏感词检测了
    数据库修改或者删除敏感词，需要调用对应的 UpdateOldWord 方法、DeleteWord 方法
```

## sensitive-word
    敏感词检测
### 性能
```
1、当字符串为：今夜总会想起你夜总最淫官员，并且用strings.Repeat重复10000次的性能如下
    text      = strings.Repeat("今夜总会想起你夜总", 10000) + "最淫官员"
    
    DFA 算法性能 BenchmarkDFAFilterAll
       BenchmarkDFAFilterAll           1000000000               0.002005 ns/op
       BenchmarkDFAFilterAll-2         1000000000               0.002516 ns/op
    
    普通遍历性能 BenchmarkDFAFilterForr
       BenchmarkDFAFilterForr         1        1249947300 ns/op
       BenchmarkDFAFilterForr-2       1        1291353500 ns/op    
   
2、当字符串为：今夜总会想起你夜总最淫官员，并且用strings.Repeat重复10次的性能如下
        text      = strings.Repeat("今夜总会想起你夜总", 10) + "最淫官员"
        
    DFA 算法性能 BenchmarkDFAFilterAll（op单次时间忽略不计） 
        BenchmarkDFAFilterAll-2         1000000000
    
    普通遍历性能 BenchmarkDFAFilterForr
        BenchmarkDFAFilterForr-2        1000000000               0.001047 ns/op
    
    
结论 
    使用DFA算法性能更好，特别是在成文本匹配中速度遥遥领先。
```
### 包含功能
```
新增敏感词：func (d *DFA) AddWord(word string)
更新敏感词：func (d *DFA) UpdateOldWord(oldWord, newWord string)
删除敏感词：func (d *DFA) DeleteWord(word string) bool

是否预处理输入文本：isPreprocessText bool
过滤敏感词：Filter(text string, isPreprocessText bool) string
检测敏感词：Check(text string, isPreprocessText bool) error

```


