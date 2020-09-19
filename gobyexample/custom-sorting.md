# [Go by Example](../gobyexample.md): Custom Sorting

除了前面的基本排序外，針對結構的某個欄位排序，後面還有另一個方法，先看 lambda 函式範例，或 callback

``` go
package main

import (
    "fmt"
    "sort"
)

func main() {
	family := []struct {
	    Name string
	    Age  int
	}{
	    {"Alice", 23},
	    {"David", 2},
	    {"Eve", 2},
	    {"Bob", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
	    return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
}
```
[執行](https://play.golang.org/p/CccxydFpWzj)

``` shell
$ go run custom-sorting.go
[{David 2} {Eve 2} {Alice 23} {Bob 25}]
```

下一範例: [Custom Sorting2](custom-sorting2.md)
下一範例: [Sorting by Functions](sorting-by-functions.md)
