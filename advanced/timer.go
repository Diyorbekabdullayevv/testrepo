package advanced // timer.go

import (
  "fmt"
  "time"
)

func LongRunningOperation() {
  for i := range 20 {
    fmt.Println("i:", i)
    time.Sleep(time.Second)
  }
}

func WorkWithTimer() {
  // WorkWithTimer1()
  // WorkWithTimer2()
  // WorkWithTimer3()
  WorkWithTimer4()
}

func WorkWithTimer1() {
  // time.Sleep(time.Second)
  fmt.Println("Starting point!")
  timer := time.NewTimer(2 * time.Second)
  fmt.Println("Before timer expiration!")
  if timer.Stop() {
    fmt.Println("Timer stopped!")
  }
  timer.Reset(time.Second)
  fmt.Println("Timer reseted!")
  <-timer.C
  fmt.Println("Timer expired!")
}

func WorkWithTimer2() {
  timeout := time.After(2 * time.Second)
  done := make(chan bool)

  go func() {
    LongRunningOperation()
    done <- true
  }()

  select {
  case <-timeout:
    fmt.Println("Operation timed out!")
  case <-done:
    fmt.Println("Operation completed!")
  }
}

func WorkWithTimer3() {
  timer := time.NewTimer(2 * time.Second)

  go func() {
    <-timer.C
    fmt.Println("Deleyed operation executed!")
  }()

  fmt.Println("Waiting...")
  time.Sleep(3 * time.Second)
  fmt.Println("End of the program!")
}

func WorkWithTimer4() {
  timer1 := time.NewTimer(1 * time.Second)
  timer2 := time.NewTimer(2 * time.Second)

  select {
  case <-timer1.C:
    fmt.Println("Timer1 expired!")
  case <-timer2.C:
    fmt.Println("Timer2 expired!")
  }
}