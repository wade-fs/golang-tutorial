# [Go by Example](../gobyexample.md): Switch

Switch 敘述允許一個條件式有多重判斷值作為分支。  

``` go
package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("Write ", i, " as ")

    switch i {			// 基本的分支用法跟 C/C++ 很像
    case 1:
        fmt.Println("one")	// 最大的不同是『不必』break，不會繼續往下跑
    case 2:
        fmt.Println("two")
    case 3:			// 後面也不必用 C/C++ 慣用的 default:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {	// 這邊示範的是用函式傳回值
    case time.Saturday, time.Sunday:	// 而 case 也可以有多種值，這就像 C 裡沒有 break 的用法
        fmt.Println("It's the weekend")
    default:				// go 當然也支持 default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {			// 很特別的，go 裡的 switch 另一大特色是 完全不必條件式
    case t.Hour() < 12:		// 這種用法更像是 if ... else 的組合
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {	// 這邊是無名函式，或稱為函式指標
        // 判斷型別用例， 此時的 i 依呼叫時決定內容，[interface](interface.md)是 go 特色
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```
[執行](http://play.golang.org/p/QlMkcwHvmns)

``` shell
$ go run switch.go 
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```

整理一下:  
- switch 可以像 C/C++ 一樣的標準用法
- default: 可有可無
- switch 可以不必條件式
- case: 不必 break 就不會往下走，這用法才是更直覺式的語言定義
- case 允許多重值判斷，這種用法像 C/C++ 中沒有 break 的多個 case

下一範例: [Arrays](arrays.md)
