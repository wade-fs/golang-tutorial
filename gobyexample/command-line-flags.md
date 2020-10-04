# [Go by Example](../gobyexample.md): Command Line Flags

命令行旗標是為命令行程序指定選項的常用方法。  
例如，在wc -l中，-l是命令行旗標。

``` go
package main

import (
    "flag" // Go提供了一個標誌包，支持基本的命令行標誌解析。
    "fmt"
)

var (
    wordPtr *string
    numbPtr *int
    boolPtr *bool
    svar    string
)

// 習慣上我們會把命令行旗標放在 init() 中初始化
// 一方面用以簡明 main(), 一方面保證程式在一開始就集中初始化
func init() {
    // 兩類方法，一類傳回指標
    wordPtr = flag.String("word", "foo", "a string")
    numbPtr = flag.Int("numb", 42, "an int")
    boolPtr = flag.Bool("fork", false, "a bool")

    // 一類輸入變數
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()
}

func main() {
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
```
[執行](https://play.golang.org/p/QuFMmABF-Jx)

``` shell
# 為了實驗各種命令行旗標，建議編譯成執行檔
$ go build command-line-flags.go

# 旗標用的是 -FLAG=VALUE 這樣的語法
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 後面 a1 a2 a3 這種非旗標的，算是參數，見前一示例
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# 旗標受限於非旗標會中斷解釋
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# 底下是內建的說明 -h，也可以在程式中明確指定
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# 使用到非指定的旗標的話會停止
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
```

下一範例: [Command-Line Subcommands](command-line-subcommands.md)
