# [Go by Example](../gobyexample.md): Command Line Arguments


``` go
package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
```
[執行](http://play.golang.org/p/oSxtj7v_v1K)

``` shell
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
```


下一範例: [命令列旗標](command-line-flags.md)
