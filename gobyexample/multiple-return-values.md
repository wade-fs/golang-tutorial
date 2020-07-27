# [Go by Example](../gobyexample.md): Multiple Return Values

Go `內建函式` 很常遇到多傳回值的情況，算是 Go 的慣用法。
通常多傳回值的後面是 error, 後面會看到:


``` go
package main

import (
    "errors"
    "fmt"
)

func devide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Devide by zero")
    }
    return a/b, nil
}

func main() {
    a, b := 10, 0
    ratio, err := devide(a, b)
    if err != nil {
	fmt.Println(err)
	return
    }
    fmt.Printf("Devide %d by %d got %+v\n", a, b, ratio)
}
```
[執行](https://play.golang.org/p/HzCQOBpULUn)

``` shell
$ go run prog.go 
Devide by zero
```

如果把 b 設為非 0 值，傳回 error = nil 做為成功依據，這樣就可以少掉 try...catch

``` go
package main

import (
    "errors"
    "fmt"
) 

func devide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Devide by zero")
    }
    return a/b, nil
}

func main() {
    a, b := 10, 2
    ratio, err := devide(a, b)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Devide %d by %d got %+v\n", a, b, ratio)
}

```

[執行](https://play.golang.org/p/Xsb_w15b9SH)

``` shell
$ go run prog.go
Devide 10 by 2 got 5
```


下一範例: [Variadic Functions](variadic-functions.md)
