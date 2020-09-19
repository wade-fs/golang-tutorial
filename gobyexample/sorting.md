# [Go by Example](../gobyexample.md): Sorting

Go的排序程序包可對內置函數和用戶定義的類型進行排序。   
我們將首先考慮對內建函數進行排序。

``` go
package main

import (
    "fmt"
    "sort"
)

func main() {

    strs := []string{"c", "a", "b"}	// 排序方法特定於內置類型。這是字符串的示例。
    sort.Strings(strs)			// 請注意，排序是就地進行的，
    fmt.Println("Strings:", strs) 	// 因此它會更改給定的切片，並且不會返回新的切片。

    ints := []int{7, 2, 4}		// 同上，排序正常的整數串
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
```
[執行](http://play.golang.org/p/_gY0tANzJ4l)

``` shell
$ go run sorting.go
Strings: [a b c]	// 運行我們的程序將打印排序後的字符串和int切片，
Ints:    [2 4 7]
Sorted:  true 		// 並根據AreSorted測試結果返回true。
```

有興趣者可以參閱[自訂排序](custom-sorting.md)

下一範例: [Sorting by Functions](sorting-by-functions.md)
