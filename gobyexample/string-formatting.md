# [Go by Example](../gobyexample.md): String Formatting

Go為 C 程式語言中的 printf 字符串格式提供了出色的支持。  
以下是一些常見的字符串格式化任務示例。

``` go
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}

    // 對物件(結構)提供各種簡易格式
    fmt.Printf("%v\n", p)	// {1 2}
    fmt.Printf("%+v\n", p)	// {x:1 y:2} // 我覺得最好用的方式
    fmt.Printf("%#v\n", p)	// main.point{x:1, y:2} // 除錯方便
    fmt.Printf("%T\n", p)	// main.point 想成印出其 Type
    fmt.Printf("%t\n", true)	// 印出真假值

    // 數字格式也相當方便
    fmt.Printf("%d\n", 123)		// 123 最常用的用法
    fmt.Printf("%b\n", 14)		// 1110 印二進制
    fmt.Printf("%c\n", 33)		// ! 印 ASCII 碼字元
    fmt.Printf("%x\n", 456)		// 1c8 印16進制，%X 為 1C8, 大小寫很自由
    fmt.Printf("%f\n", 78.9)		// 浮點數
    fmt.Printf("%e\n", 123400000.0)	// 科學符號表示法，跟 %x %X 一樣有大小寫
    fmt.Printf("%E\n", 123400000.0)

    // 對字串內容，也有不同格式
    fmt.Printf("%s\n", "\"string\"")	// "string" 這是最常用的
    fmt.Printf("%q\n", "\"string\"")	// "\"string\"" q 想成 quote, 用在 SQL, HTTP 很有用
    fmt.Printf("%x\n", "hex this")	// 6865782074686973 用來編碼傳輸很有用

    fmt.Printf("%p\n", &p)		// 指標

    // 在指定格式時，同時指定長度
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // 將訊息丟到另一個字串中
    s := fmt.Sprintf("a %s", "string")	// 使用 fmt.Sprintf() 也很常見
    fmt.Println(s)			// 可以將不同資料變成字串

    // 將訊息丟到 stderr 中，或者使用 io.Writer 功能
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
```
[執行](http://play.golang.org/p/L6BkGeaN_p4)
``` shell
$ go run string-formatting.go
{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0x42135100
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error
```

下一範例: [正規表示法](regular-expressions.md)
