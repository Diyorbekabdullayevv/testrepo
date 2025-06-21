package advanced // worker_pools.go

import (
  "fmt"
  "time"
)

type ticketRequest struct {
  personID   int
  numTickets int
  cost       int
}

func TicketProcessor(requests <-chan ticketRequest, results chan<- int) {
  for req := range requests {
    fmt.Printf("Processing %d ticket(s) of personID %d with total cost of 💲%d dollars\n", req.numTickets, req.personID, req.cost)
    time.Sleep(time.Second)
    results <- req.personID
  }
}

func Worker(id int, tasks <-chan int, results chan<- int) {
  for task := range tasks {
    fmt.Printf("Worker %d is processing task %d\n", id, task)
    time.Sleep(time.Second)
    results <- task * 2
  }
}

func WorkWithWorkerPools() {
  //   WorkWithWorkerPools1()
  WorkWithWorkerPools2()
  // WorkWithWorkerPools3()
  // WorkWithWorkerPools4()
}

func WorkWithWorkerPools1() {
  numWorkers := 3
  numJobs := 10
  tasks := make(chan int, numJobs)
  results := make(chan int, numJobs)

  for i := range numWorkers {
    go Worker(i, tasks, results)
  }

  for i := range numJobs {
    tasks <- i
  }

  close(tasks)

  for range numJobs {
    result := <-results
    fmt.Println("Result:", result)
  }
}

func WorkWithWorkerPools2() {
  numRequests := 5
  price := 5
  ticketRequests := make(chan ticketRequest, numRequests)
  ticketResults := make(chan int)

  for range 3 {
    go TicketProcessor(ticketRequests, ticketResults)
  }

  for i := range numRequests {
    ticketRequests <- ticketRequest{personID: i + 1, numTickets: (i + 1) * 2, cost: ((i + 1) * 2) * price}
  }
  close(ticketRequests)

  for range numRequests {
    fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
  }
}
