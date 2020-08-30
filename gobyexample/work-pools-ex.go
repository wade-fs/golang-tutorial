package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
	start := time.Now()
        fmt.Println("worker", id, "started  job", j, "at", start)
        time.Sleep(time.Duration(id) * time.Second)
        fmt.Println("worker", id, "finished job", j, "at", time.Now(), "spent", time.Now().Sub(start))
        results <- j * 2
    }
}

func main() {
    start := time.Now()
    fmt.Println("main Start at", start)
    const numJobs = 15
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= numJobs; a++ {
        <-results
    }
    fmt.Println("main finish at", time.Now(), "spent", time.Now().Sub(start))
}
