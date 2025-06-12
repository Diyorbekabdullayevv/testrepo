package advanced // channel synchronisation

import (
  "fmt"
  "time"
)

func RegularGoroutines() {

  fmt.Println("Befor goroutine!")

  go func() {
    time.Sleep(2 * time.Second)
    fmt.Println("Something processing...")
  }()

  fmt.Println("After goroutine!")
  time.Sleep(1 * time.Second)
  fmt.Println("The end!")
}

func WorkWithChanSync() {
  //   WorkWithChanSync1()
  // WorkWithChanSync2()
  // WorkWithChanSync3()
  WorkWithChanSync4()

  //   RegularGoroutines()

}

func WorkWithChanSync1() {
  done := make(chan struct{})

  go func() {
    fmt.Println("Working...")
    time.Sleep(1 * time.Second)
    done <- struct{}{}
  }()

  <-done
  fmt.Println("Finished!")
}

func WorkWithChanSync2() {
  ch := make(chan int)

  go func() {
    ch <- 9
    fmt.Println("Value sent!")

  }()

  value := <-ch
  fmt.Println("Value recieved!")
  fmt.Println("Value:", value)
}

func WorkWithChanSync3() {
  numGoroutines := 3
  done := make(chan int)

  for i := range numGoroutines {
    go func(id int) {
      fmt.Printf("Goroutine %d working...\n", id)
      done <- id
    }(i)
  }
  for range numGoroutines {
    <-done
  }
  fmt.Println("All goroutines are finished!")
}

func WorkWithChanSync4() {
  data := make(chan string)

  go func() {
    for i := range 5 {
      data <- "Hello " + string(i+'1')
      time.Sleep(100 + time.Millisecond)
      // time.Sleep(2 * time.Second)
    }
    close(data)
  }()
  
  for value := range data {
    fmt.Println("Recieved value:", value, ":", time.Now())
  }

}
