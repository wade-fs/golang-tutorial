# [Go by Example](../gobyexample.md): Testing

單元測試是依Go原則編寫程序的重要組成部分。 測試包提供了編寫單元測試所需的工具，而go test命令則運行測試。  
為了演示起見，此代碼位於main包中，實際運用上測試檔可以是任何包。 測試代碼通常與所測試的代碼位於同一包中。

``` go
package main

import (
    "fmt"
    "testing"
)

func IntMin(a, b int) int {	// 這邊先實作一個小函式，通常這個函式是實作在獨立的完整檔案中
    if a < b {			// 例如這個函式可能存在 intutils.go
        return a		// 那麼底下測試內容就存在 intutils_test.go
    }				// 這樣執行 go test intutils.go 就會自動找 intutils_test.go 
    return b
}

func TestIntMinBasic(t *testing.T) {	// 跟測試檔名相反的，以 Test開頭會被自動調用進行測試
    ans := IntMin(2, -2)		// t 含有很多工具，像這邊的 t.Errorf()
    if ans != -2 {

        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

func TestIntMinTableDriven(t *testing.T) {	// *testing.T 會包含很多有用的工具
    var tests = []struct {			// 而且以 TableDriven 風格是慣用法
        a, b int				// 其中包含 input,
        want int				// 及 output
    }{						// 這種方法有點像 online-judge :)
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {	// t.Run() 更是常用的工具
            ans := IntMin(tt.a, tt.b)		// t.Run() 會啟動子測試程序
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
```
[執行](http://play.golang.org/p/GFuPdlBlyMU)

``` shell
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing    0.023s
```

上面用 go test -v 進行 verbose 模式，會顯示比較多訊息

下一範例: [命令列參數](command-line-arguments.md)
