# [Go by Example](../gobyexample.md): Random Numbers

Go 的 math/rand 包提供虛擬亂數(因為理論上計算機無法產生真正的亂數，除非是量子電腦)

``` go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 底下兩行產生2個 [0,100) 的整數值亂數，也就是 0<= n < 100
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // Go 簡單易用的地方在於，不像很多語言只產生整數或浮點數一種亂數
    // 底下產生的是 [0,1) 的亂數。事實上你大概是看不到 0.0 的亂數值啦
    fmt.Println(rand.Float64())

    // 底下方法是一般將亂數放大到特定值，例如: [5.0,10.0)
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // 後面結果裡面顯示，重複執行後，前面的『亂數值』竟然每次都一樣
    // 這牽涉到亂數理論，我們只要想像一件事，「亂數」是算出來的
    // 它會受『初始化』影響，因此，亂數通常用時間作引子來初始化
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    // 這樣下面的亂數值，因為每次執行，前面初始化的值不同
    // 才有辦法讓接下來的亂數值變成『真正』的亂數
    // 當然，開頭說過，目前的計算機產生的都是『虛擬亂數』
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // 當然可以初始化多次，只是我覺得意義不大
    // 所以底下就來以相同的值做初始化，看看是不是會產生相同序列的亂數
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()

    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
```
[執行](http://play.golang.org/p/0ooeler0RfR)
``` shell
$ go run random-numbers.go
81,87
0.6645600532184904
7.123187485356329,8.434115364335547
0,28
5,87
5,87
```

``` shell
再執行一遍看看:
81,87
0.6645600532184904
7.1885709359349015,7.123187485356329
0,28
5,87
5,87
```

進階請閱讀 Go [math/rand](http://golang.org/pkg/math/rand/) 包的說明文件

下一範例: [數值轉換](number-parsing.md)
