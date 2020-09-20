# [Go by Example](../gobyexample.md): Regular Expressions

Regular Expressions 不好翻譯，用 google 翻譯是『常用表達』，但是它並不常用  
有的書翻『正則式』有的書翻『正規式』，我採用後者  
正規表示法的用途在於要在較長文本中搜索的「字符串」或「樣式的符號」和「字符序列」。

``` go
package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {
    // 比對字串 "peach" 是否符合 p開頭 ch結尾，中間是 [a-z] 一個以上的組合
    // 這邊將中間 [a-z]+ 用 () 括住的話，在別的函式可以用來單獨取得結果
    // 但是對 MatchString() 這類函式來說，() 只是為了方便閱讀，可以省略
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)	// true

    // 上例 regexp.MatchString() 可以分成兩段, 也就是下面兩行與上例同義
    r, _ := regexp.Compile("p([a-z]+)ch")
    fmt.Println(r.MatchString("peach"))		// true

    fmt.Println(r.FindString("peach punch"))	// peach

    fmt.Println(r.FindStringIndex("peach punch"))// [0 5] 比對字串為 [0:5], 即 peach

    fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea] ea 就是因為() 括住，可單獨取出

    fmt.Println(r.FindStringSubmatchIndex("peach punch"))// [0 5 1 3] 與上行比較下

    fmt.Println(r.FindAllString("peach punch pinch", -1))// [peach punch pinch]

    // [[0 5 1 3] [6 11 7 9] [12 17 13 15]] 請與r.FindAllStringSubmatch() 比較
    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    // [[peach ea] [punch un] [pinch in]]
    fmt.Println(r.FindAllStringSubmatch."peach punch pinch", -1))
    // [peach punch] 因為後面指定只要2筆
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach"))) // true

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)			  // p([a-z]+)ch

    // a <fruit>
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // 底下可以類似『收集功能』那樣，透過函式來做正規運算
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper) // a PEACH
    fmt.Println(string(out))
}
```
[執行](http://play.golang.org/p/LEKGY_d3Nu_P)
``` shell
$ go run regular-expressions.go 
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
```

完整的說明請見[Go 正規表示法](http://golang.org/pkg/regexp/)套件說明  

下一範例: [JSON](json.md)
