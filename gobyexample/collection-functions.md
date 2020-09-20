# [Go by Example](../gobyexample.md): Collection Functions

我們經常需要我們的程序對數據集合執行操作，例如選擇滿足給定謂詞的所有項或使用自定義函數將所有項映射到新集合。這種需要泛稱為『收集功能』或『收集函式』

在某些語言中，慣用的是通用的數據結構和算法(物件導向的關係)。  
Go不支持泛型； 在Go中，慣用上是提供收集函式在你的程序及資料上  

此處是一些用於字符串切片([]string)的示例收集函數。  
您可以使用這些示例來構建自己的功能。   

請注意，在某些情況下，最簡單的方法是直接內聯收集操作代碼，而不是創建和調用輔助程序函數。  
(參考 main() 最後一行)

``` go
package main

import (
    "fmt"
    "strings"
)

// 在字串切片 vs中搜尋目標字串 t, 傳回其索引，若找不到則傳回 -1
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// 類上，但是傳回真假值(有/無)
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// 這邊採用的不是搜尋目標字串本身，而是
// 把字串切片 vs 代入函式中，有任意為真即真，否則為假
// 通常這種作法在別的程式語言中意謂著 map() 或 filter() 這類函式
//   請見下方 Filter()
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// 與上面互補，要所有字串切片成員代入 f() 都為真方為真，否則為假
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// 跟 Any() 異曲同工之妙，將所有滿足 f() 的切片成員傳回
// 這種功能也很適合應用在 映射(map)中
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// 跟 Any() 異曲同工之妙，將所有串字切片應用 f()並傳回
// 這邊的函式名採取的是其他程式語言慣用的 map(), 
// 千萬不要跟 Go 的 map 資料結構搞混
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {
    // 底下會將上面的所有收集功能用一遍
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

    // 上面的收集功能採用的是無名函式的方法，
    // 底下則使用內建的函式在收集功能上
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
