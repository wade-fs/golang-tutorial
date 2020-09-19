# [Go by Example](../gobyexample.md): Collection Functions

我們經常需要我們的程序對數據集合執行操作，例如選擇滿足給定謂詞的所有項或使用自定義函數將所有項映射到新集合。

在某些語言中，慣用的是通用的數據結構和算法。  
Go不支持泛型； 在Go中，慣用上是提供收集函式在你的程序及資料上  

此處是一些用於字符串切片([]string)的示例收集函數。  
您可以使用這些示例來構建自己的功能。   

請注意，在某些情況下，最簡單的方法是直接內聯收集操作代碼，而不是創建和調用輔助程序函數。

``` go
package main

import (
    "fmt"
    "strings"
)

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    fmt.Println(Map(strs, strings.ToUpper))

}
```
[執行](http://play.golang.org/p/uKnePZM91WV)

``` shell
$ go run collection-functions.go 
2
false
true
false
[peach apple pear]
[PEACH APPLE PEAR PLUM]
```


下一範例: [String Functions](string-functions.md)
