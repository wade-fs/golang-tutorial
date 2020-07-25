# [Go by Example](../gobyexample.md): For

`for` 跟 C/C++ 很像, 用於迴圈結構，但是跟 C/C++ 的差異性又很明顯，一個字：簡  
很重要的點就是，go 僅僅只有 for 迴圈，沒有 while do, do while 這類的迴圈  

``` go
package main

import "fmt"

func main() {

    i := 1			// 宣告變數的最快方法，自動型態判定
    for i <= 3 {		// 條件式『不必也不能』有 (...), 其餘跟 C/C++ 類似
        fmt.Println(i)
        i = i + 1
    }

    for j := 7; j <= 9; j++ {	// 跟 C/C++ 類似的用法
        fmt.Println(j)
    }

    for {			// go 沒有 while 迴圈，這邊算是 do while 的變型
        fmt.Println("loop")
        break
    }

    for n := 0; n <= 5; n++ {	// 跟 C/C++ 類似的用法
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```
[執行](http://play.golang.org/p/2-4H-ArwHHS)

``` shell
$ go run for.go
1
2
3
7
8
9
loop
1
3
5
```

就 go 而言，for 的用法相當多變，上面與 C/C++ 類似用法外，我們在其餘章節也會提到  
for 的慣用法反而是搭配 range() 及多返回值  

下一範例: [If/Else](if-else.md)
