# [Go by Example](../gobyexample.md): Defer

Defer用於確保在程序執行的後期執行函數調用，通常是為了進行清理。  
Defer 通常用於像其他程式語言使用的 ensure 及 finally 用法.  

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 假設我們要建立新檔案，並寫入資料
    // 由於檔案前後操作可能很複雜，
    // 也就是在很後面還在進行各種不同的檔案操作
    // 此時我們很可能忘了『關閉』它，這會造成意外狀況
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)	// 使用 defer 將確保程式在最後會主動幫忙關閉
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

// 比較執行結果會發現此函式會在最後被呼叫，而不是 defer() 的位置
func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close() 
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
```
[執行](http://play.golang.org/p/5SDVfc_jxbg)

``` shell
$ go run defer.go
creating
writing
closing
...
exit status 2
```

檔案的各項操作有不同的錯誤可能會發生，進行必要的檢查是相當重要的，  
前例中有 os.Create() 及 f.Close() 都可能發生錯誤


下一範例: [collection functions](collection-functions.md)
