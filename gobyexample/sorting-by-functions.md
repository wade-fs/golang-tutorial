# [Go by Example](../gobyexample.md): Sorting by Functions(https://gobyexample.com/collection-functions)

有時，我們會希望按照自然順序以外的其他方式對集合進行排序。  
例如，假設我們要按字符串的長度而不是按字母順序對它們進行排序。  
這是Go中自定義排序的示例。

``` go
package main

import (
    "fmt"
    "sort"
)

// 為了要自定排序函式，需要相對應的自定Slice型態
// 底下示例是 string slice, 可以是複雜結構
type byLength []string

// 自訂 sort 需要三個函式 Len(), Swap(), Less()
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}
```
[執行](http://play.golang.org/p/h4g4vaLBtkw)

``` shell
$ go run sorting-by-functions.go 
[kiwi peach banana]
```

在[Custom Sorting2](custom-sorting2.md)有說明，這邊就不多說

下一範例: [Panic](panic.md)
