# [Go by Example](../gobyexample.md): String Functions

標準庫的字串包提供了許多有用的與字串相關的功能。 以下是一些示例，可以使您對字串有所了解。


``` go
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println	// 這邊展示簡化寫法

func main() {

    // `Strings` 的函式很多，可以仔細看範例來學
    p("含有: ", s.Contains("test", "es"))
    p("字數: ", s.Count("test", "t"))
    p("前置: ", s.HasPrefix("test", "te"))
    p("後置: ", s.HasSuffix("test", "st"))
    p("索引: ", s.Index("test", "e"))
    p("合併: ", s.Join([]string{"a", "b"}, "-"))
    p("重複: ", s.Repeat("a", 5))
    p("取代: ", s.Replace("foo", "o", "0", -1))
    p("取代: ", s.Replace("foo", "o", "0", 1))
    p("分割: ", s.Split("a-b-c-d-e", "-"))
    p("小寫: ", s.ToLower("TEST"))
    p("大寫: ", s.ToUpper("test"))
    p()

    p("長度: ", len("hello"))
    p("字元:", "hello"[1])
}
```
[執行](http://play.golang.org/p/fZ_FqN5WlSz)

``` shell
$ go run string-functions.go
含有:  true
字數:  2
前置:  true
後置:  true
索引:  1
合併:  a-b
重複:  aaaaa
取代:  f00
取代:  f0o
分割:  [a b c d e]
小寫:  test
大寫:  TEST

長度:  5
字元: 101
```

完整說明請見[String-Functions](https://golang.org/pkg/strings/).

下一範例: [字串函式](string-functions.md)
