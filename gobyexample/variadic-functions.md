# [Go by Example](../gobyexample.md): 參數個數變化的函式

Go 函式參數個數是可以變化的，呼叫時參數可依需要給予，必須是相同型態，否則函式內必須先予以判斷型態  
最常見的內建函式就是 [fmt](https://golang.org/pkg/fmt/).Printf(), [fmt](https://golang.org/pkg/fmt/).Println() 等。


``` go
package main

import "fmt"

func sum(nums ...int) {		// 此處利用 ... 指示參數個數不定, nums 就是個 []int
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {	// 用 range 來瀏覽參數 nums
        total += num
    }
    fmt.Println(total)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)		// 呼叫時，個數隨需要調整

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
```
[執行](http://play.golang.org/p/Ua6kZOMabBp)

``` shell
$ go run variadic-functions.go 
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

函數還有一個議題，叫做 [閉包](closures.md)

下一範例: [閉包](closures.md)
