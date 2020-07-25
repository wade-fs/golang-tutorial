# [Go by Example](../gobyexample.md): Hello World

我們的第一個程序將打印經典的“ hello world”消息。 這是完整的源代碼。

``` go
package main

	

import "fmt"

	

func main() {
    fmt.Println("hello world")
}
```
[執行](http://play.golang.org/p/NeviD0awXjt)

要運行該程序，請將代碼放在 `hello-world.go` 中並使用 `go run` 。

``` shell
$ go run hello-world.go
hello world
```

有時我們想將程序構建為二進製文件。 我們可以使用 `go build` 來做到這一點。

``` shell
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

之後就可以執行它:  

``` shell
$ ./hello-world
hello world
```

下一範例: [值](values.md)
