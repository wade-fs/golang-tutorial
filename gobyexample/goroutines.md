# [Go by Example](../gobyexample.md): Errors

`錯誤` 是 Go 語言另一重要特色  
在Go中，習慣性地通過顯式的，單獨的返回值傳達錯誤。  
這個語言特色一方面讓人覺得煩索，卻讓人不必使用套疊的 try catch, 也讓人不易掉進錯誤裡面。  
這與Java和Ruby等語言中使用的異常以及C中有時使用的重載單個結果/錯誤值形成對比。  
也就是說，像 C 語言，一般函式傳回值只能單個，錯誤的處理就相當麻煩，而 Go的方法可以輕鬆查看哪些函數返回錯誤，並使用與其他任何函數相同的語言構造來處理錯誤及非錯誤任務。

``` go
package main

import (
    "errors"	// errors 是內建界面, 後面會看到傳回值有 error 型別
    "fmt"
)

func f1(arg int) (int, error) { // 多傳回值中函有 error 很方便, 後面 if 會用到
    if arg == 42 {
				// 底下 errors.New() 建立新的 error 值
        return -1, errors.New("can't work with 42")

    }

    return arg + 3, nil		// 如果沒錯誤，就讓它的 error 為 nil
}

// 下面的結構與其方法的結合，可以自訂錯誤型別
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string { // 只要結構具備 Error() 就算是 error 型別
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
				// 可以用 argError 當錯誤型別返回
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {	// 這是 Go 的慣用法
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {	// 再用慣用法
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {	// 改用 ok(=!e) 會不會更舒服點？
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
```
[執行](http://play.golang.org/p/NiJOpCPO3L0)

``` shell
$ go run errors.go
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it
```

因為錯誤及早處理過，上面第一個 for 程式片段，可以寫成如下方式，既減少縮排又更易讀:    

``` go  
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {	// 這是 Go 的慣用法
            fmt.Println("f1 failed:", e)
            continue
        }
        fmt.Println("f1 worked:", r)
        .....
    }
```

有興趣的讀者可以參考[Go 語言的錯誤處理](#https://blog.golang.org/error-handling-and-go)一文  

下一範例: [Goroutines](goroutines.md)
