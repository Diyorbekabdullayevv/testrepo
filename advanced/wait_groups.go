package advanced // wait_groups.go

import (
  "fmt"
  "sync"
  "time"
)

type Worker5 struct {
  ID   int
  Task string
}

func (w *Worker5) PerformTask(wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
  time.Sleep(time.Second)
  fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func Worker2(id int, wg *sync.WaitGroup) {

  defer wg.Done()
  fmt.Printf("Worker %d starting...\n", id)
  time.Sleep(time.Second)
  fmt.Printf("Worker %d finished!\n", id)
}

func Worker3(id int, results chan<- int, wg *sync.WaitGroup) {

  defer wg.Done()
  fmt.Printf("Worker %d starting...\n", id)
  time.Sleep(time.Second)
  results <- id * 2
  fmt.Printf("Worker %d finished!\n", id)
}

func Worker4(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {

  defer wg.Done()
  fmt.Printf("Worker %d starting...\n", id)
  time.Sleep(time.Second)
  for task := range tasks {

    results <- task * 2
  }
  fmt.Printf("Worker %d finished!\n", id)
}

func WorkWithWaitGroups() {
  // WorkWithWaitGroups1()
  // WorkWithWaitGroups2()
  // WorkWithWaitGroups3()
  WorkWithWaitGroups4()
}

func WorkWithWaitGroups1() {
  var wg sync.WaitGroup
  numWorkers := 3
  wg.Add(numWorkers)

  for i := range numWorkers {
    go Worker2(i+1, &wg)
    // time.Sleep(time.Millisecond)
  }

  //   fmt.Println("Waiting for goroutines to finish...")
  wg.Wait() // Block until all Done() are called
  fmt.Println("All goroutines finished!")
}

func WorkWithWaitGroups2() {
  var wg sync.WaitGroup
  numWorkers := 3
  numJobs := 3
  results := make(chan int, numJobs)
  wg.Add(numWorkers)

  for i := range numWorkers {
    go Worker3(i+1, results, &wg)
    // time.Sleep(time.Millisecond)

  }
  go func() {
    wg.Wait()
    close(results)
  }()

  for result := range results {
    fmt.Println("Result:", result)
  }
}

func WorkWithWaitGroups3() {
  var wg sync.WaitGroup
  numWorkers := 3
  numJobs := 5
  results := make(chan int, numJobs)
  tasks := make(chan int, numJobs)
  wg.Add(numWorkers)

  for i := range numWorkers {
    go Worker4(i+1, tasks, results, &wg)
    // time.Sleep(time.Millisecond)
  }

  for i := range numJobs {
    tasks <- i + 1
  }
  close(tasks)

  go func() {
    wg.Wait()
    close(results)
  }()

  for result := range results {
    fmt.Println("Result:", result)
  }
}

func WorkWithWaitGroups4() {
  var wg sync.WaitGroup
  tasks := []string{"digging", "laying bricks", "painting"}

  for i, task := range tasks {
    worker := Worker5{ID: i + 1, Task: task}
    wg.Add(1)
    go worker.PerformTask(&wg)
    time.Sleep(time.Millisecond)
  }

  wg.Wait()

  fmt.Println("Construction finished!")
}