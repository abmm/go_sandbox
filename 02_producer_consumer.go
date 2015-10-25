package main

/* producer-consumer problem in Go */

import ("fmt"
				"time")

var done = make(chan bool)
var msgs = make(chan int)

func produce () {
		i := 0
    for {
    		i++
        msgs <- i
        fmt.Printf("Produced element : %d \n" , i)
        time.Sleep(time.Second)
    }
    done <- true
}

func consume () {
    for {
      msg := <-msgs
      fmt.Printf("Consumed element : %d \n" , msg)
    }
}

func start_producer_consumer () {
   go produce()
   go consume()
   <- done
}