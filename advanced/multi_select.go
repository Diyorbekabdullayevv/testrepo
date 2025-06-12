package advanced // multiplexing select // multi_select.go 

import (
  "fmt"
  "time"
)

func WorkWithMultiplexingSelect() {
  // WorkWithMultiplexingSelect1()
  // WorkWithMultiplexingSelect2()
  WorkWithMultiplexingSelect3()
}

func WorkWithMultiplexingSelect1() {

  ch1 := make(chan int)
  ch2 := make(chan int)

  go func() {
    time.Sleep(time.Second)
    ch1 <- 1
  }()

  go func() {
    // time.Sleep(time.Second)
    ch2 <- 2
  }()

  time.Sleep(2 * time.Second)

  for range 2 {
    select {
    case msg := <-ch1:
      fmt.Println("Receieved from ch1:", msg)
    case msg := <-ch2:
      fmt.Println("Receieved from ch2:", msg)
    default:
      fmt.Println("No channels ready!")
    }
  }
  fmt.Println("The end!")

}

func WorkWithMultiplexingSelect2() {
  ch := make(chan int)

  go func() {
    time.Sleep(2 * time.Second)
    ch <- 1
    close(ch)
  }()

  select {
  case msg := <-ch:
    fmt.Println("Receieved:", msg)
  case <-time.After(3 * time.Second):
    fmt.Println("Timeout!")
    // default:
    //   fmt.Println("No channels ready!")
  }
}

func WorkWithMultiplexingSelect3() {

  ch := make(chan int)

  go func() {
    ch <- 1
    close(ch)
  }()

  for {
    select {
    case msg, ok := <-ch:
      if !ok {
        fmt.Println("Channel closed!")
        return
      }
      fmt.Println("Receieved:", msg)
    }
  }
}