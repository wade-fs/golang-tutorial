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

老練的調用者可以使用類型斷言來檢查NegativeSqrtError並進行特殊處理，而僅將錯誤傳遞給fmt.Println或log.Fatal的調用者不需要更改對錯誤處理的慣用法。  

再舉一個例子，json包指定了SyntaxError類型，當json.Decode函數在解析JSON blob時遇到語法錯誤時，該類型將返回。它的結構定義如下:

``` go  
type SyntaxError struct {
    msg    string // 錯誤訊息內容
    Offset int64  // 錯誤發生的位置
}

func (e *SyntaxError) Error() string { return e.msg }
```

錯誤的默認格式甚至都沒有顯示“偏移”字段，但是調用者可以使用它將文件和行信息添加到錯誤消息中：  

``` go  
if err := dec.Decode(&val); err != nil {
    if serr, ok := err.(*json.SyntaxError); ok {
        line, col := findLine(f, serr.Offset)
        return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
    }
    return err
}
```

上例利用型別轉換(err.(*json.SyntaxError)), 就可以使用 serr.Offset 找到錯誤發生的內容。  

（這是[Camlistore](http://camlistore.org/)項目中某些[實際代碼](https://github.com/camlistore/go4/blob/03efcb870d84809319ea509714dd6d19a1498483/jsonconfig/eval.go#L123-L135)的略微簡化的版本。）

錯誤接口只需要一個Error方法； 特定的錯誤實現可能還有其他方法。 例如，[net](https://golang.org/pkg/net/)包遵循通常的約定返回錯誤類型為error的錯誤，但是某些錯誤實現具有由net.Error接口定義的其他方法：   
``` go
package net

type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```

客戶端代碼可以使用類型斷言來測試net.Error，然後將瞬態網絡錯誤與永久性網絡錯誤區分開。 例如，網絡搜尋器在遇到臨時錯誤時可能會休眠並重試，否則會放棄。  

``` go
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)
    continue
}
if err != nil {
    log.Fatal(err)
}
```

## 簡化重複性錯誤處理

在Go中，錯誤處理很重要。 該語言的設計和約定鼓勵您明確檢查錯誤發生的位置（不同於其他語言中的引發異常並有時捕獲它們的約定）。 在某些情況下，這會使Go代碼變得冗長，但是幸運的是，可以使用一些技術來減少重複性錯誤處理。

考慮具有HTTP處理程序的[App Engine](https://cloud.google.com/appengine/docs/go/)應用程序，該處理程序從數據存儲區檢索記錄並使用模板對其進行格式化。

``` go
func init() {
    http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```

此函數處理datastore.Get()函數和viewTemplate.Execute()方法返回的錯誤。 在兩種情況下，它都會向用戶顯示一條簡單的錯誤消息，並帶有HTTP狀態代碼500（“內部服務器錯誤”）。 這看起來像是可管理的代碼量，但是添加更多的HTTP處理程序，您很快就會得到許多相同錯誤處理代碼的副本。

為了減少重複，我們可以定義自己的HTTP appHandler類型，其中包括錯誤返回值：  
``` go
type appHandler func(http.ResponseWriter, *http.Request) error
```
然後，我們可以更改viewRecord函數以返回錯誤：  

``` go  
func viewRecord(w http.ResponseWriter, r *http.Request) error {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return err
    }
    return viewTemplate.Execute(w, record)
}
```

這比原始版本更簡單，但是http包不理解返回錯誤的函數。 為了解決這個問題，我們可以在appHandler上實現http.Handler接口的ServeHTTP方法：  

``` go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```
ServeHTTP方法調用appHandler函數，並向用戶顯示返回的錯誤（如果有）。 注意，該方法的接收者fn是一個函數。 （Go可以做到！）該方法通過在表達式fn（w，r）中調用接收方來調用函數。

現在，當使用http包註冊viewRecord時，我們使用Handle函數（而不是HandleFunc），因為appHandler是http.Handler（而不是http.HandlerFunc）。  

``` go
type appError struct {
    Error   error
    Message string
    Code    int
}
```

接下來，我們修改appHandler類型以返回* appError值：  

``` go
type appHandler func(http.ResponseWriter, *http.Request) *appError
```

（出於[Go FAQ](https://golang.org/doc/go_faq.html#nil_error)中討論的原因，傳回錯誤而不是錯誤的具體類型通常是一個錯誤，但這是正確的做法，因為ServeHTTP是唯一看到值並使用其內容的地方。） 

使appHandler的ServeHTTP方法使用正確的HTTP狀態代碼向用戶顯示appError的消息，並將完整的Error記錄到開發者控制台：  

``` go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if e := fn(w, r); e != nil { // e is *appError, not os.Error.
        c := appengine.NewContext(r)
        c.Errorf("%v", e.Error)
        http.Error(w, e.Message, e.Code)
    }
}
```
最後，我們將viewRecord更新為新的函數簽名，並使其在遇到錯誤時返回更多上下文：  
``` go
func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return &appError{err, "Record not found", 404}
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        return &appError{err, "Can't display record", 500}
    }
    return nil
}
```

此版本的viewRecord與原始版本的長度相同，但是現在這些行中的每一行都有特定含義，我們將提供更友好的用戶體驗。

它並沒有就此結束； 我們可以進一步改善應用程序中的錯誤處理。 一些想法：

- 給錯誤處理程序一個漂亮的HTML模板，
- 當用戶是管理員時，通過將堆棧跟踪寫入HTTP響應，可以簡化調試工作，
- 為appError編寫一個構造函數，該函數存儲堆棧跟踪信息以便於調試，
- 從appHandler內的緊急情況中恢復，將錯誤記錄為“嚴重”到控制台，同時告訴用戶“發生了嚴重錯誤”。 這是避免用戶暴露於由編程錯誤引起的難以理解的錯誤消息的好方法。 有關更多詳細信息，請參見[延緩，緊急情況和恢復](https://golang.org/doc/articles/defer_panic_recover.html)文章。
