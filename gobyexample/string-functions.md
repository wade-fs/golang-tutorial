# [Go by Example](../gobyexample.md): String Functions

標準庫的字串包提供了許多有用的與字串相關的功能。 以下是一些示例，可以使您對字串有所了解。  
以我來看，如果能將字串當成物件，而不是透過 strings 來操作就會很方便。  

``` go
package main

import (
    "fmt"
    s "strings"		// 這邊示範將長的相依套件名 strings 更名為短的 s
)

var p = fmt.Println	// 這邊展示簡化寫法，將 fmt.Println() 更名為 p()
			// 這種寫法看起來簡便，但是不易閱讀

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

    // 底下兩個不是字符串的一部分，
    // 值得一提的是用於獲取字符串長度（以字節為單位）並按索引獲取字節的機制。
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

請注意，上面的len和indexing在字節級別工作。  
Go使用UTF-8編碼的字符串，因此通常是非常有用的。  
如果您使用的可能是多字節字符，則需要使用可識別編碼的操作。 
有關更多信息，請參見Go中的字符串，字節，符文和字符。

完整說明請見[Strings](https://golang.org/pkg/strings/).

下一範例: [字串格式](string-formatting.md)
