package advanced // closing channels // close_chan.go

import "fmt"

func Producer1(ch chan<- int) {

  for i := range 5 {
    ch <- i
  }
  close(ch)
}

func Filter(in <-chan int, out chan<- int) {
  for val := range in {
    if val%2 == 0 {
      out <- val
    }
  }
  close(out)
}

func WorkWithClosingChannels() {
  // WorkWithClosingChannels1()
  // WorkWithClosingChannels2()
  WorkWithClosingChannels3()
}

func WorkWithClosingChannels1() {

  ch := make(chan int)

  go func() {
    for i := range 5 {
      ch <- i
    }
    close(ch)
  }()

  for val := range ch {
    fmt.Printf("%d recieved: %d\n", val+1, val)
  }

}

func WorkWithClosingChannels2() {

  ch := make(chan int)
  close(ch)

  val, ok := <-ch
  if !ok {
    fmt.Println("Channel is closed!")
    return
  }

  fmt.Println("Value:", val+1)
}

func WorkWithClosingChannels3() {
  ch1 := make(chan int)
  ch2 := make(chan int)

  go Producer1(ch1)
  go Filter(ch1, ch2)

  for val := range ch2 {
    fmt.Println("Value:", val)
  }
}