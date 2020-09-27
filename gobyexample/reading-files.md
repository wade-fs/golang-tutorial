# [Go by Example](../gobyexample.md): Reading Files

讀取和寫入文件是許多Go程序所需的基本任務。  
首先，我們來看一些讀取文件的示例。

``` go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// 使用 panic() 方便檢查錯誤，錯誤發生時停止程式
// Go 的 if err != nil { ... } 也讓我感覺到厭煩
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // 將文本內容一次讀入記憶體中是最省事的
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // 底下則手動控制要讀進來的部份，
    f, err := os.Open("/tmp/dat")
    check(err)
    // 例如只讀 5 bytes，有時這代表特定檔頭
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
    // 或是跳到某位置之後開始讀檔，例如該位置被記錄在檔頭
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: ", n2, o2)
    fmt.Printf("%v\n", string(b2[:n2]))

    // 承上，更明確指定要讀『至少幾個』bytes
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // 再回到檔案的開頭
    _, err = f.Seek(0, 0)
    check(err)

    // 使用 bufio 基本上對於每次讀一小部份或是並發式讀檔會更有效率
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
}
```
[執行](http://play.golang.org/p/kF0cDC0drsX)
``` shell
$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
```


下一範例: [Writing Files](writing-files.md)
