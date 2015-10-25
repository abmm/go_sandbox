package main

import (
				"fmt"
 				"strconv"
 				"time"
    		"sync"
    		"sync/atomic"

)



func start_mutex() {


	c1 := make(chan string)
	c2 := make(chan string)

	var mutex = &sync.Mutex{}

	var totalsum int32 = 0 

	go inc(c1, "hilo1",500, &totalsum, mutex)
	go dec(c2, "hilo2",1500, &totalsum, mutex)
	go dec(c2, "hilo3",100, &totalsum, mutex)
	//go inc(c3, "hilo3",7500)

	for{
		select{
			case msg1 := <- c1:
				fmt.Println("INC received", msg1)
			case msg2 := <- c2:
				fmt.Println("DEC received", msg2)
			case <- time.After(time.Second * 1050):
				fmt.Println("--------------")
		}
	}



}


func inc(channel chan string, hilo string, tiempo int , sum *int32, mutex *sync.Mutex) (int){
	for{

		mutex.Lock()
		if *sum < 10 {
    	atomic.AddInt32(sum, 1)
    	channel <- hilo + " -- "+ strconv.Itoa(int(*sum))
    }else{
    	channel <- hilo + " LLENO " + strconv.Itoa(int(*sum))
    }
    mutex.Unlock()
		
		time.Sleep(time.Duration(tiempo) * time.Millisecond)
		
	}
	return 0
}


func dec(channel chan string, hilo string, tiempo int , sum *int32, mutex *sync.Mutex) (int){
	for{

		mutex.Lock()
		if *sum > 0 {
    	atomic.AddInt32(sum, -1)
    }else{
    	channel <- hilo + " VACIO " + strconv.Itoa(int(*sum))
    }
    mutex.Unlock()


		time.Sleep(time.Duration(tiempo) * time.Millisecond)
		channel <- hilo + " -- "+ strconv.Itoa(int(*sum))
	}
	return 0
}



