# [Go by Example](../gobyexample.md): Strings

`Range` 範圍迭代各種數據結構中的元素。 讓我們看看如何將範圍用於我們已經學習的一些數據結構。  
比較常見的就是用在 `array`, `slice`, `map` 或 `struct` 中


``` go
package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}		// 這是切片 `slice`, 並已賦值
    sum := 0
    for _, num := range nums {		// range 前項是索引，後項是值
        sum += num			// 通常不需要索引時就用 _ 忽略它，有點像 /dev/null
    }					// 如果跟 python 比較可以發現十分相像
    fmt.Println("sum:", sum)

    for i, num := range nums {		// 見面前的說明
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {		// slice, array, map 都能用 range
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {		// 這邊是忽略後者，相當於 k,_ := ...
        fmt.Println("key:", k)
    }

    for i, r := range "go" {		// 可以發現 字串 是 Unicode 字元陣列的事實
        fmt.Println(i, r)		// r 正確說明叫 rune, 如後述
    }
}
```
[執行](http://play.golang.org/p/pdZOtv4g-7J)

``` shell
$ go run range.go
sum: 9
index: 1
a -> apple
b -> banana
key: a
key: b
0 103
1 111
```

在字串中沒有特別聲明字串的元素並非 C/C++ 中的字元，而是叫做 `rune`, 請繼續看下一範例 [Strings](strings.md)

下一範例: [字串函式](string-functions.md)
