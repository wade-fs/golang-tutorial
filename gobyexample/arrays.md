# [Go by Example](../gobyexample.md): Arrays

在 go 語言裡，`array` 是明確長度的元素序列  
這樣說很難表達清楚，先強調一下， go 的 `array` 跟 C/C++ 的『陣列』意函完全不同  
C/C++ 的陣列意函比較接近的是另一個詞 [Slice](slice.md)  

``` go
package main

import "fmt"

func main() {

    var a [5]int	// 建立 **5** 個整數的 `array`, 此時的個數與型態都明確指定  
                        // 若未給初值，整數的初始值是 0
    fmt.Println("emp:", a)

    a[4] = 100		// 跟所有語言的 `array` 用法一樣，直接修改某元素值
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])	// 取得值也是跟所有語言一樣

    fmt.Println("len:", len(a))	// 取得 `array` 長度跟 python 一樣用 len()

    b := [5]int{1, 2, 3, 4, 5}	// 宣告時同時指定初始值靠 {...}
    fmt.Println("dcl:", b)

    var twoD [2][3]int		// 二維的宣靠也很直覺
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j	// 二維的用法跟 C/C++ 一樣
        }
    }
    fmt.Println("2d: ", twoD)
}
```
[執行](http://play.golang.org/p/TaahifSGSwU)

``` shell
$ go run arrays.go
emp: [0 0 0 0 0]	// 特別說明，在 fmt.Println 時的 Array 沒有明確的分隔
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]
```

go 的 `array` 類似常數，無法更動長度, 甚至也不是像 C/C+ 一樣的指標意函，  
因此常見的傳進函式後的更動，返回後 `array` 的值不會被更動。  
建議與 [Slice](slice.md)一起看，會有更多說明  

下一範例: [Slices](slices.md)
