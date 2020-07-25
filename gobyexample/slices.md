# [Go by Example](../gobyexample.md): Slices

`Slices` 在 go 語言中是更關鍵性的型態，比前面的 `array` 更彈性的用法。

``` go
package main

import "fmt"

func main() {

    s := make([]string, 3)	// slice 與 array 最大不同，用 make() 製造 slice 空間
    fmt.Println("emp:", s)

    s[0] = "a"	// 用法上跟 array 很像，常讓人分別不出是 array 或 slice
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    fmt.Println("len:", len(s))	// 一樣用 len() 取得長度

    s = append(s, "d")		// 透過 append() 來增加元素， `array` 就無法這麼做
    s = append(s, "e", "f")	// append() 一次可以增加多個元素
    fmt.Println("apd:", s)

    c := make([]string, len(s))
    copy(c, s)			// 透過 copy() 來複製 s 到新的 `slice` c
    fmt.Println("cpy:", c)	// 複製後的新切片 c 就是獨立的記憶體空間

    l := s[2:5]			// 這邊的 l 就如英文的 `slice` 一詞，它就是切片
    fmt.Println("sl1:", l)

    l = s[:5]			// 指定切片忽略前面索引就是從頭(0)開始
    fmt.Println("sl2:", l)

    l = s[2:]			// 指定切片忽略後面索引就是直到尾部
    fmt.Println("sl3:", l)

    t := []string{"g", "h", "i"}	// 這語法跟 `array` 很像，沒有長度的宣告就是切片
    fmt.Println("dcl:", t)

    twoD := make([][]int, 3)	// 二維切片的宣告，此處 3 是指第一維長度 3
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        // 此處示例二維切片與二維 Array 最大的不同
        twoD[i] = make([]int, innerLen)	// 這邊每一維又是一維切片，長度可以不同
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```
[執行](http://play.golang.org/p/iLnoIEIxeQ1)

``` shell
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]
```

查看Go團隊撰寫的精彩[博客文章](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)，以獲取有關Go中slice的設計和實現的更多詳細信息。   
記住 array 與 slice 最大的不同:  
array 明確指定長度，而 slice 宣告時只有型態，所以其個數可變  
事實上，slice 內部也是透過 array 來實作，細節部份就當做進階課題吧


下一範例: [Maps](maps.md)
