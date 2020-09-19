# [Go by Example](../gobyexample.md): Custom Sorting 2

除了前面的基本排序外，針對結構的某個欄位排序，後面還有另一個方法，先看 lambda 函式範例，或 callback

``` go
package main

import (
    "fmt"
    "sort"
)

// 這邊的結構跟前例相似，
// 但是以這邊的自訂排序示例來說，可以做到很複雜
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
// 自訂 Slice 型態，做物件化處理
type ByAge []Person

// 只要定義底下三個物件方法(Len(), Less(), Swap())，後面就可以排序
// 按 Go 語言特性，也可以 (a *ByAge) 指標方式
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    family := []Person{
        {"Alice", 23},
        {"Eve", 2},
        {"Bob", 25},
    }
    sort.Sort(ByAge(family))
    fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
```
[執行](https://play.golang.org/p/CccxydFpWzj)

``` shell
$ go run custom-sorting2.go
[{Eve 2} {Alice 23} {Bob 25}]
```

下一範例: [Custom Sorting for MAP](custom-sorting-for-map.md)
下一範例: [Sorting by Functions](sorting-by-functions.md)
