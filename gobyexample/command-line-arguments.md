# [Go by Example](../gobyexample.md): Command Line Arguments

命令行參數是參數化程序執行的常用方法。  
例如，go run hello.go使用go程序的run和hello.go參數。

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    // os.Args提供對原始命令行參數的訪問。  
    // 請注意，此片中的第一個值是程序的路徑，
    // 而os.Args [1：]保存了程序的參數。
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    // 您可以使用標準索引獲得單個arg。
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
```
[執行](http://play.golang.org/p/oSxtj7v_v1K)

``` shell
// 要嘗試使用命令行參數，最好先使用go build來構建二進製文件。
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
```

接下來，我們將介紹帶有標誌的更高級的命令行處理。

下一範例: [命令列旗標](command-line-flags.md)
