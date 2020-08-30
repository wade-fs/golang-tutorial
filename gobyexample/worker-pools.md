# [Go by Example](../gobyexample.md): Worker Pools

在此示例中，我們將研究如何使用 *goroutine* 和通道實現 *Worker Pool* (工作池)。  

``` go
package main

import (
    "fmt"
    "time"
)

// 這是工作程序，我們將運行其中的多個並發實例。
// 這些 Workers 將在工作渠道上接收工作，
// 並根據(接收工作的)結果發送相應的結果。
// 我們將為每項工作休息一秒鐘，以模擬工作負載。
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    const numJobs = 5			// 為了使用我們的工作池
    jobs := make(chan int, numJobs)	// 需要向他們發送工作並收集他們的結果
    results := make(chan int, numJobs)	// 為此，我們製作了2個渠道

    // 這將啟動3個工作, 初始時因為沒有丟入工作所以會阻斷
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 在這裡，我們發送了5個工作......
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // 然後關閉該渠道以表明我們已經完成了所有工作。

    // 最後，我們收集所有工作結果。
    // 這也確保了工作程序已完成。
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```
[執行](http://play.golang.org/p/hiSJJsYZJKL)

``` shell
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real    0m2.358s
```
我們的運行程序顯示了由不同工作執行的5個作業。  
儘管有3名工人同時工作，但該程序僅花費大約2秒鐘（儘管完成了大約5秒鐘的總工作）。  
等待多個 *goroutine* 的另一種方法是使用*[WaitGroup]*(waitgroups.md)。

下一範例: [WaitGroups](wait-groups.md)
