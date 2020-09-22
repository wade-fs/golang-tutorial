# [Go by Example](../gobyexample.md): Number Parsing

數學上或是計算機領域有一門叫數值分析的課題，這邊講的是非常簡單的數字的轉換. 


``` go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // Go 的有效位元數，一般常用的有 2/8/16/32/64,
    // 但是浮點數只限 32/64 兩種
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // 整數除了有效位元數外，基底也很重要，
    // 底下基底 0 表示常見的限定值， 0b 二進位, 0 或 0o 8進位，0x 16進位
    // 所以底下三例仔細看看就清楚了
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("0x789", 0, 64)
    fmt.Println(u)

    // Atoi() 是類 C 的方法，簡便易用，限整數
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // 所以非整數就錯了
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
```
[執行](https://play.golang.org/p/OG6ZHVUlyNj)
``` shell
$ go run number-parsing.go 
1.234
123
456
1929
135
strconv.ParseInt: parsing "wat": invalid syntax
```

接下來請見另一解析工作 Url

下一範例: [URL 解析](url-parsing.md)
