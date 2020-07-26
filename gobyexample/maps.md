# [Go by Example](../gobyexample.md): Maps

`Maps` 映射是Go內置的關聯數據類型（有時用其他語言稱為哈希或字典）。這型態跟其他程式語言差異不大。  


``` go
package main

import "fmt"

func main() {

    m := make(map[string]int)	// go 建立變數空間是透過 make(), 使用 := 賦值省工易讀

    m["k1"] = 7			// 這邊 get/set 跟其它語言是一樣的
    m["k2"] = 13

    fmt.Println("map:", m)

    v1 := m["k1"]		// 這邊 get/set 跟其它語言是一樣的
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))	// 取長度透過 len() 具通用性，很好記

    delete(m, "k2")		// set 直接增加，刪除則透過 delete()
    fmt.Println("map:", m)

    _, prs := m["k2"]		// 判斷 key/value 是否存在的正確用法
				// 若直接取值判斷有時會因為值本身是 0, nil, false 而判斷失誤
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2} // go 建立變數空間賦值總是用 {...} 
    fmt.Println("map:", n)
}
```
[執行](http://play.golang.org/p/agK2Ro2i-Lu)

``` shell
$ go run maps.go 
map: map[k1:7 k2:13]		// 顯示映射內容與 Arrays 很像，都缺 '' 或 , 等明確分隔元素
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
```


下一範例: [Range](range.md)
