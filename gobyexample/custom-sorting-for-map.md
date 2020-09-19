# [Go by Example](../gobyexample.md): Custom Sorting by map

標準示範及前例都在講結構 Slice 的排序，實際上可以應用在 map 上  
如果把 map 看成無序的 Struct Slice 就很容易理解

``` go
package main

import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

    // 先取得 keys slice of map
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    // 針對 keys 排序
    sort.Strings(keys)

    for _, k := range keys {
        fmt.Println(k, m[k])
    }
}
```
[執行](https://play.golang.org/p/CccxydFpWzj)

``` shell
$ go run custom-sorting2.go
[{Eve 2} {Alice 23} {Bob 25}]
```

下一範例: [Sorting by Functions](sorting-by-functions.md)
