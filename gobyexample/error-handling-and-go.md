# [Go 語言的錯誤處理](https://blog.golang.org/error-handling-and-go)

## 簡介  

如果您編寫了任何Go代碼，會很常遇到內置錯誤類型。 Go代碼使用錯誤值指示異常狀態。 例如，當os.Open函數無法打開文件時，它返回一個非nil錯誤值。

``` go
func Open(name string) (file *File, err error)

f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
..... 其餘程式碼放這邊，不必把程式碼放在 else 中, 因為 log.Fatal() 會離開程式 .....  
```

幾乎只要看到 Go 程式碼就會遇到錯誤處理，也只有在累積足夠多的錯誤處理知識才有辦法完成更多的工作。  
本文將仔細研究錯誤處理，並探討 Go 中錯誤處理的良好慣用法。  

## error 型別  

錯誤型別是個界面型別，這讓程式師很容易自訂錯誤型別。一個錯誤變數的內函是『任何可以將自己表示為字串的值』。還記得前面提到結構的 String() 是個特別的多型用法，我們來看看關於 error 的[內建定義](https://golang.org/pkg/errors/):  

``` go
type error interface {
    Error() string
}
```

再回頭看[錯誤處理](errors.md)時，有示範 `type argError struct` 自訂錯誤型態，這邊再舉個例:  

``` go
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

.... 略 ....

func New(text string) error {
    return &errorString{text}	// 使用自訂錯誤型態返回 error
}
```

還有一種比較常見的是使用 errors.New() 來產生錯誤內容:  

``` go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 繼續實作....
}

if f, err := Sqrt(-1); err != nil {
    fmt.Println(err)	// 此處的 err 是叫用自訂錯誤型態中的 Error() 函式
}
```

關於 Sqrt() 的錯誤處理訊息，仔細看你會發現並不是很清楚，就像常見的 os.Open() 會告訴你錯誤發生的原因，這邊可以透過 fmt.Errorf() 來進一步顯示訊息。  

``` go  
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, fmt.Errorf("math: square root of negative number %g", f)
    }
    // implementation
}
```

使用 fmt.Errorf() 已經相當夠用(而且常見), 但是既然 error 型別是個界面型別，可以進一步使用任意資料結構為錯誤值，讓呼叫方檢視錯誤內容。底下舉個不錯的簡單範例，假設調用者可能想要恢復傳遞給Sqrt的無效參數，我們可以通過定義一個新的錯誤實現而不是使用errors.errorString來實現:  

``` go
type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
    return fmt.Sprintf("math: square root of negative number %g", float64(f))
}
```


