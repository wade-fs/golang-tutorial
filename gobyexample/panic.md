# [Go by Example](../gobyexample.md): Panic

一個 Panic 通常意味著發生了意外錯誤(導致系統停止執行)。  
我們會使用它來快速解決正常操作期間不應該發生的錯誤，  
或者我們不准備妥善處理的錯誤(通常)。

``` go
package main

import "os"

func main() {
    // 這邊故意示範，程式在此就停了，後面沒機會執行到
    panic("a problem")

    // 若用底下當示例(將前一行註解掉), 會是慣用法(之一)
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
```
[執行](http://play.golang.org/p/h4g4vaLBtkw)

``` shell
$ go run panic.go
panic: a problem		// 出現 Panic 訊息，
goroutine 1 [running]:		// 後面指示進一步資訊以供除錯
main.main()
    /.../panic.go:12 +0x47
...
exit status 2
```

請注意，與某些使用例外處理許多錯誤的語言不同，在Go中，盡可能使用錯誤指示返回值才是慣用的。  
意思就是說，Go 語言不使用 try ... catch 用法，而是像此例，用 err != nil 來處理錯誤。  

下一範例: [Defer](defer.md)
