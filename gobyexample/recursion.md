# [Go by Example](../gobyexample.md): 遞迴

這篇會介紹幾個實用的閉包用法。

# 無名函式

``` go

package main

import "fmt"

var DoStuff func() = func() {		// DoStuff 是以函式變數宣告
  fmt.Println(".........START.......")
}

func main() {
  DoStuff()				// 它可以直接調用

  DoStuff = func() {			// 也可以再次賦與另一個無名函式
    fmt.Println("Doing stuff!")
  }
  DoStuff()				// 再次調用時就會變更其函式內容

  DoStuff = func() {			// 同理，可以再次賦與另一個無名函式
    fmt.Println("Doing other stuff.")
  }
  DoStuff()				// 再次調用時就會變更其函式內容
}
```
[執行](https://play.golang.org/p/wi1TFPIoKzE)

``` shell
$ go run prog.go
.........START.......
Doing stuff!
Doing other stuff.
```

# 閉包

閉包說白了，就是無名函式加上使用無名函式外面的變數  

``` go

package main

import "fmt"

func main() {
  n := 0			// 1) 無名函式外面已定義的變數
  counter := func() int {
    n += 1			// 2) 在無名函式內部使用外部的變數
    return n			//    此時無名函式內部的外部變數值，為調用前一刻的值
  }
  fmt.Println(counter())	// 所以第一次調用時，n <- 0
  fmt.Println(counter())	// 第二次調用時, n <- 1
}
```
[執行](https://play.golang.org/p/KU0UTne4y1P)

``` shell
$ go run prog.go
1
2
```

至此，似乎只是利用外部變數，感覺沒有什麼意外(驚喜). 

## 隔離數據

``` go
package main

import "fmt"

func main() {
  fib := makeFibGen()		// 調用 makeFibGen() 利用其傳回值『產生』新的函式
  for i := 0; i < 10; i++ {
    fmt.Println(fib())
  }
}

func makeFibGen() func() int {	// 因為費式數列需要兩個變數，又不想污染外面...
  f1 := 0			// 在 makeFibGen() 內宣告兩個局域變數
  f2 := 1
  return func() int {		// 相對傳回的無名函式來說，f1, f2 是外部變數
    f2, f1 = (f1 + f2), f2	// 因為被 makeFibGen() 包住，所以隔離了 f1, f2 
    return f1			// 使得外面乾乾淨淨的
  }
}
```
[執行](https://play.golang.org/p/xdzGid-9l1x)

``` shell

1
2
3
5
8
13
21
34
55
```

## 包裝函式，以便創建中間件（插件）

Go中的功能是一級變量。 這意味著您不僅可以動態創建匿名函式，還可以將函式作為參數傳遞給函式。 例如，在創建Web服務器時，通常會提供處理到特定路由的Web請求的功能。  

使用閉包創建中間件之前，先了解一下什麼叫中間件。  
中間件基本上是可重用函數的一個高級術語，它可以在設計用於處理Web請求的代碼之前和之後運行代碼。 在Go中，這些通常是通過閉包來完成的，但是在不同的編程語言中，它們可以通過其他方式來實現。

在編寫Web應用程序時，通常使用中間件，它們不僅對 `即時計時器` (如下)有用。 例如，中間件可用於編寫代碼以驗證用戶是否已登錄一次，然後將其應用於您的所有僅限會員頁面。

``` go
package main

import (
  "fmt"
  "net/http"
  "time"
)

func main() {
  http.HandleFunc("/hello", timed(hello)) // 這邊也可以直接調用 hello, 而不經過 timed() 包裝
  http.ListenAndServe(":3000", nil)
}

// 底下是中間件，可實作在別的套件中
//
// timed() 將參數傳進來的函式包裝起來，使得可以計算整個運算期間所花時間
// 因為 timed() 是包住 api 要求，因此返回值型態就必須符合規範
// 仔細跟 hello() 定義比對就可以清楚
func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()	// 先記住開始時間
    f(w, r)		// 調用參數傳進來的函式
    end := time.Now()	// 再記住結束時間
    fmt.Println("The request took", end.Sub(start)) // 就可以算出整個執行時間
  }
}

// hello() 也是中間件，也是可以被定義在別的套件中
func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "<h1>Hello!</h1>")
}
```

## 訪問不可見的數據

這個議題前面有提到，這邊強調它的實用性。  

``` go
package main

import (
  "fmt"
  "net/http"
)

type Database struct {
  Url string
}

func NewDatabase(url string) Database {
  return Database{url}
}

func main() {
  db := NewDatabase("localhost:5432")

  http.HandleFunc("/hello", hello(db))
  http.ListenAndServe(":3000", nil)
}

// 底下是中間件，可實作在別的套件中
func hello(db Database) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, db.Url)
  }
}
```

Database 結構有可能更複雜的處理而被包裝起來，這示例單純的返回用於示意。  
在 web api "/hello" 中，我們傳進去的是 hello(db), 而不是直接的 db  
也就是說，利用 hello() 裡面的閉包(返回無名函式)將 db 包起來，形成閉包  
此處無名函式可以做更複雜的資料處理。當然此處只用來取得 db.Url

怕有人還是像看無字天書般看不懂，這邊再囉嗦一點點。  

假設 hello() 被實作在別的套件中.....  
那麼，它怎麼處理 Database 我們就看不到，或是說，它可以處理 sqlite3, mysql, postgresql, 對我們(main())而言也不是重要議題，可以讓我們專心開發自己的主程式。當然中間件也可以專心開發它的功能. 

## 搜尋與排序

這邊主要要講的不是[搜尋與排序](https://golang.org/pkg/sort/)本身，而是要說明利用閉包時的優點，以應用在搜尋與排序示意  

``` go
package main

import (
  "fmt"
  "sort"
)

func main() {
  numbers := []int{1, 11, -5, 8, 2, 0, 12}
  sort.Ints(numbers)		// 利用 sort 套件來排序，這沒問題
  fmt.Println("Sorted:", numbers)

  // 但是，如果，想找的值並不在裡面怎麼辦？
  // 例如，我想找第一筆比 7 大的數字怎麼辦？
  // 因為資料(numbers)是在外部，利用閉包就很有用
  // 底下的 func(i int) bool { ... } 就是閉包
  index := sort.Search(len(numbers), func(i int) bool {
    return numbers[i] >= 7
  })
  fmt.Println("The first number >= 7 is at index:", index)
  fmt.Println("The first number >= 7 is:", numbers[index])
}
```
[執行](https://play.golang.org/p/Hy2oFJ2jAT2)

``` shell
$ go run prog.go
Sorted: [-5 0 1 2 8 11 12]
The first number >= 7 is at index: 4
The first number >= 7 is: 8
```

## 延遲工作

先來看個示例:

``` go
package main

import "fmt"
// 疑问一
func f() (ret int) {
    defer func() {
        ret++
    }()
    return 1
}
func main() {
    fmt.Println(f())
}
```
[執行](https://play.golang.org/p/KwE0XJ0HCSZ)

``` shell
$ go run prog.go
2
```

噫？結果怎麼會是 2？  
主要是因為 defer 在是 return 之後發生，而閉包參考的變數則是在閉包『產生』時引用。
所以， f() 呼叫返回 1, 這在 go 的定義中，會指定給 ret, 意思是 ret 在 f() 內部視為局域變數。  
結果 defer 在 return 發生後會產生無名函式，它引用的 ret 值是 1，所以最後 ret 的值是 2


下一範例: [遞迴](recursion.md)
