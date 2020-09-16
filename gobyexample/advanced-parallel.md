# [Go by Example](../gobyexample.md): Stateful Goroutines

並發是個複雜的課題，所以....  
首先，我們列出一個清單，以便對並發做進階研究的基礎。

* 先讀讀 [有效的 Go: 並發](https://golang.org/doc/effective_go.html#concurrency)
* 看影片教學: [用 Go 模擬真實世界系統](https://www.dotconferences.com/2017/11/sameer-ajmani-simulating-a-real-world-system-in-go)
* 研讀[Go 程式語言規範書](https://golang.org/ref/spec), 尤其是  
  - [Go 敘述](https://golang.org/ref/spec#Go_statements)
  - [通道型態](https://golang.org/ref/spec#Channel_types)
  - [傳送敘述](https://golang.org/ref/spec#Send_statements)
  - [接收運算子](https://golang.org/ref/spec#Receive_operator)
  - [Select 敘述](https://golang.org/ref/spec#Select_statements)
* 讀讀程式碼: [並發教學](http://tour.golang.org/concurrency/1)
* 讀讀[常見的問與答](http://golang.org/doc/faq), 尤其是  
  - [為什麼要在CSP的思想上建立並發性？](http://golang.org/doc/faq#csp)  
    CSP 原文是 循序程序間的通訊  
  - [為什麼使用goroutines而不是線程？](http://golang.org/doc/faq#goroutines)
  - [為什麼未將映射操作定義為原子操作？](http://golang.org/doc/faq#atomic_maps)
  - [什麼是原子操作？ 互斥鎖呢？](http://golang.org/doc/faq#What_operations_are_atomic_What_about_mutexes)
  - [為什麼我的程序在使用更多CPU時不能更快地運行？](https://golang.org/doc/faq#parallel_slow)
  - [如何控制CPU的數量？](https://golang.org/doc/faq#number_cpus)
  - [作為goroutines運行的閉包會發生什麼？](http://golang.org/doc/faq#closures_and_goroutines)

請立即閱讀:  
* 研讀 [Go by Example](../gobyexample.md) from [goroutines](goroutines.md) through [stateful goroutines](stateful-goroutines.md)
* 觀看 [Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1)
* 觀看 [A Practical Guide to Preventing Deadlocks and Leaks in Go](https://www.youtube.com/watch?v=3EW1hZ8DVyw)
* 研讀 [Share Memory By Communicating](http://blog.golang.org/share-memory-by-communicating) 與實作 [codewalk](http://golang.org/doc/codewalk/sharemem/)
* 閱讀 [Go Concurrency Patterns: Timing out, moving on](http://blog.golang.org/go-concurrency-patterns-timing-out-and)
* 觀看 [Concurrency is not Parallelism](http://talks.golang.org/2012/waza.slide#1)
* 閱讀 [Go Concurrency Patterns: Pipelines and Cancellation](http://blog.golang.org/pipelines)
* 閱讀 [Rethinking Classical Concurrency Patterns](https://github.com/golang/go/wiki/Go-Community-Slides#rethinking-classical-concurrency-patterns)
* 探討 [Package sync](https://golang.org/pkg/sync/)
* 研讀 [Introducing the Go Race Detector](http://blog.golang.org/race-detector)
* 觀看 [Go: code that grows with grace](http://talks.golang.org/2012/chat.slide#1)
* 研讀 [Mutexes and Semaphores Demystified](http://www.barrgroup.com/Embedded-Systems/How-To/RTOS-Mutex-Semaphore)

進階:
* 觀看 [Advanced Go Concurrency Patterns](http://blog.golang.org/advanced-go-concurrency-patterns)
* 研讀 [Advanced Go Concurrency Patterns](http://talks.golang.org/2013/advconc.slide#1)
* 研讀 [Go Concurrency Patterns: Context](http://blog.golang.org/context)
* 探討 [The Go Memory Model](https://golang.org/ref/mem)
* 探討 [Package atomic](https://golang.org/pkg/sync/atomic/)
* 研讀 [Principles of Designing Go APIs with Channels](https://inconshreveable.com/07-08-2014/principles-of-designing-go-apis-with-channels/)
* 研讀 [Advanced Go Concurrency Primitives](https://encore.dev/blog/advanced-go-concurrency)
* 觀看 [The Scheduler Saga](https://www.youtube.com/watch?v=YHRO5WQGh0k)
* 研讀 [The Scheduler Saga](https://speakerdeck.com/kavya719/the-scheduler-saga)
* 觀看 [Understanding Channels](https://www.youtube.com/watch?v=KBZlN0izeiY)
* 研讀 [Understanding Channels](https://speakerdeck.com/kavya719/understanding-channels)

下一範例: [排序](sorting.md)
