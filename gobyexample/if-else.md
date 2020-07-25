# [Go by Example](../gobyexample.md): If ... Else

程式執行過程的分支就要靠 if 及 else 的搭配，  
在 go 裡面就只有直接式的用法，並沒有其他語法，像 elif, elsif, elseif

``` go
package main

import "fmt"

func main() {

    if 7%2 == 0 {			// 跟 for 一樣，條件式『不必也不能』有 (...)
        fmt.Println("7 is even")
    } else {				// if ... else 是標準用法
        fmt.Println("7 is odd")
    }

    if 8%4 == 0 {			// 也可以不必搭配 else
        fmt.Println("8 is divisible by 4")
    }

    // 底下的 if 條件式中透過 `:=` 另宣告局域變數 num
    if num := 9; num < 0 {		// 這邊的用法類似 for, 這樣的用法才是 go 的特色
        fmt.Println(num, "is negative")
    } else if num < 10 {		// 多重 if else 就是像這個用例
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
```
[執行](http://play.golang.org/p/QlMkcwHvmns)

``` shell
$ go run if-else.go 
7 is odd
8 is divisible by 4
9 has 1 digit
```

上面用例中的 if 條件式有例用 `;` 的多重敘述，if 靠最後一條當條件式判斷依據。  
跟 for 類似，因為 go 有多重傳回值的特性，if 中的特殊用法在後面還會持續出現. 

下一範例: [Switch](switch.md)
