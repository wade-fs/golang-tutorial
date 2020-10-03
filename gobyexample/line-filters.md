# [Go by Example](../gobyexample.md): Line Filters

行過濾器是一種常見的程序類型，它讀取stdin上的輸入，對其進行處理，然後將一些派生結果打印到stdout。  
例如常見的命令 grep 和 sed 是常見的行過濾器。

古老的編譯器 ed 就是行編輯器，現在就讓我們來看看一次讀進一行，變成大寫後輸出。  


``` go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    // stdin 是非緩存的，先將它過濾成緩存式的 Reader
    // 主要的好處是一次可以讀一行，這是 bufio 提供的方法
    scanner := bufio.NewScanner(os.Stdin)

    // bufio.Scan() 每次可以讀進一行(可以更換 token)
    // 每次調用 Scan() 就會傳回至下一個 token 前的資料
    // 預設的 token 是換行，所以等於一次讀一行
    for scanner.Scan() { 
        // 轉成大寫
        ucl := strings.ToUpper(scanner.Text())
        // 印出來
        fmt.Println(ucl)
    }
    // 處理完資料再檢查錯誤，通常只要有錯誤就不會讀到資料
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
```
[執行](http://play.golang.org/p/kNcupWRsYPP)
``` shell
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
```

下一範例：[檔案路徑](file-paths.md)
