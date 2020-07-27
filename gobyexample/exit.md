# [Go by Example](../gobyexample.md): Exit

不知道有沒有人注意到，好像一般的 C/C++ 的 main() 宣告都是 int, 也都會有返回值 return 0;  
但是 golang 的範例怎麼好像都沒在用返回值呢？  
原因很簡單，C/C++ 一開始是用來開發 command-line 工具，當初的設想跟 golang 不同，golang 做為現代語言比較常用在 平行/Web Service 上頭，而不再是命令列工具，因此它不注重最後的返回值。  
當然要做到也很容易，就是用 os.Exit(Value), 如下

``` go
package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")	// 關於 defer 遇到 Exit() 與 return() 完全不同待遇
				// 請看後面的說明
    os.Exit(3)
}
```
[執行](http://play.golang.org/p/b9aYzlENkb__R)

``` shell
$ go run exit.go
exit status 3
```

有發現執行結果沒有出現 '!' 嗎？  
是的，defer 原本是延後處理，它只對 return 的方式有效，遇到 Exit() 時，它會被拋棄。  
如果是編譯成執行檔，要取得返回值，可以透過 shell 的 $? 來取得:  

``` shell
$ go build exit.go
$ ./exit
$ echo $?
3
```
