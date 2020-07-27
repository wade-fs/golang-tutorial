# [Go by Example](../gobyexample.md): 指標

Go 跟很多新式的程式語言不同，它支援指標，允許指定參考給值或是函式，這讓 Go 更有效率。最漂亮的是，Go 消除了 C/C++ 指標的複雜性。

``` go
package main

import "fmt"

// 底下兩個函式比較了傳值與傳參考(指標的意函就是參考到某東西)的差別，
// 如果你熟 C/C++ 的話就沒太大問題。
// 如果你是 Python/Java 的愛用者，或許會被困擾，簡單講就是:
// 傳值會複製一份資料，所以在函式中不會更動到函式外面的值
// 傳參考(指標)的話，只是該資料的參考，所以實際上會更動到外面的值。
// 通常用來直接修改外部的物件，效率相當高。
func zeroval(ival int) {	// 傳值
    ival = 0
}

func zeroptr(iptr *int) {	// 傳參考
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)			// 這邊就是指標的用法，利用 & 來取得 i 的參考指標
    fmt.Println("zeroptr:", i)

    fmt.Println("pointer:", &i)	// 指標也是一種資料，就是變數 i 存放的pointers位址, uint64

    nPtr := new(int)
    *nPtr = 2
    fmt.Println("new pointer:", nPtr, ", and value is ", *nPtr)
}
```
[執行](http://play.golang.org/p/oimmXypnAcs)

``` shell
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0xc00002c008
new pointer: 0xc00002c068 , and value is  2
```

仔細比對程式與結果，會發現 zeroval() 函式內的更動不影響外部的 i  
而 zeroptr() 函式內的更動同時影響外部的 i, 實際上是同一位置的資料。  

指標也可以搭配 new() 這種 C/C++ 標準用法。  
若變數本身是指標(就像 nPtr)的話，取值就變成 *nPtr 這種寫法。

常常會讓人搞不懂的是使用 & 與 * 的時機。暫時請自行比對上面的用法，先來看看更複雜物件的指標。

``` go
package main

import (
	"fmt"
)

type Member struct {
	Age  int
	Name string
}

func (m *Member) GetAge() int {
	return m.Age
}
func (m *Member) GetName() string {
	return m.Name
}

func main() {
	mem := &Member{
		Age:  18,
		Name: "syhlion",
	}
	fmt.Printf("Member Name:%s, Age:%d", mem.GetName(), mem.GetAge())
}
```
[執行](https://play.golang.org/p/NSqtIRbtY2l)

``` shell
$ go run prog.md
Member Name:syhlion, Age:18
```

上例說明結構指標，[結構](structs.md)我們會另外說明，其初始化也是正常，在其前面加上 & 就是指標。  
這邊比較特別的是 Go 的指標不像 C/C++ 用 `->`, 而是一樣使用結構在用的 `.`, 這樣就很方便不必特意去記住變數本身到底是指標或是普通值。  
同樣回到前面的問題，何時使用 `&`, 何時使用 `*` 就變的很簡單。就是需要取指標時用 &，需要取值是用 *。同樣適用在結構。

下一範例: [結構](structs.md)
