package advanced // channel directories // chan_directions.go

import "fmt"

func WorkWithChannelDirections() {
  ch := make(chan int)
  Producer(ch)
  // go func(ch chan<- int) {
  //   for i := range 5 {
  //     ch <- i
  //   }
  //   close(ch)
  // }(ch)

  // for value := range ch {
  //   fmt.Println("Recieved:", value)
  // }
  Reciever(ch)
}

func Producer(ch chan<- int) {
  go func() {
    for i := range 5 {
      ch <- i
    }
    close(ch)
  }()
}

func Reciever(ch <-chan int) {
  for value := range ch {
    fmt.Println("Recieved:", value+1)
  }
}