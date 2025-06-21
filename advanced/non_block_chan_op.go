package advanced // non_block_chan_op.go

import (
  "fmt"
  "time"
)

func WorkWithNonBlockingChannelOperations() {
  // WorkWithNonBlockingChannelOperations1()
  // WorkWithNonBlockingChannelOperations2()
  WorkWithNonBlockingChannelOperations3()

}

func WorkWithNonBlockingChannelOperations1() {

  ch := make(chan int)

  // go func() {
  // ch <- 18
  // }()
  // time.Sleep(1 * time.Second)

  select {
  case msg := <-ch:
    fmt.Println("Received:", msg)
  default:
    fmt.Println("No messages available!")
  }

}

func WorkWithNonBlockingChannelOperations2() {

  ch := make(chan int)

  select {
  case ch <- 1:
    fmt.Println("Value sent!")
  default:
    fmt.Println("Channel is not ready to receive!")
  }
}

func WorkWithNonBlockingChannelOperations3() {

  data := make(chan int)
  quit := make(chan bool)

  go func() {

    for {
      select {
      case d := <-data:
        fmt.Println("Data received:", d)
      case <-quit:
        fmt.Println("Stopping...")
      default:
        fmt.Println("Waiting for data...")
        time.Sleep(500 * time.Millisecond)
      }
    }
  }()

  for i := range 5 {
    data <- i
    time.Sleep(time.Second)
  }

  quit <- true

}