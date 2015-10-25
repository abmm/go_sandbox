package main 

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"redis"
	"strconv"
	"time"
	"sync/atomic"
	"runtime"
)


func addSingleKey(key string, new_value string) {
	spec 			:= redis.DefaultSpec()
	client, e := redis.NewSynchClientWithSpec (spec);
	if e != nil { log.Println ("failed to create the client", e); return }

	value, e := client.Get(key);
	if e!= nil { log.Println ("error on Get", e); return }
	if value == nil {
		value2set := []byte(new_value)
		client.Set(key, value2set);
		runtime.Gosched()
	}else{
		fmt.Println("Already found " + string(key))
	}
	//time.Sleep(time.Millisecond * 10)

}


func addKey(key string, new_value string, i *int32) {
	spec 			:= redis.DefaultSpec()
	client, e := redis.NewSynchClientWithSpec (spec);
	if e != nil { log.Println ("failed to create the client", e); return }
	for{
		atomic.AddInt32(i, 1)
		if *i > MAXKEYS{
			done <- true
		}else{
			if *i % 100000 == 0{
				fmt.Printf("Inserted %d \n" ,*i)
			}
		}
		key := key + strconv.Itoa(int(*i))
		value, e := client.Get(key);
		if e!= nil { log.Println ("error on Get", e); return }
		if value == nil {
			value2set := []byte(new_value)
			client.Set(key, value2set);
			runtime.Gosched()
		}else{
			//fmt.Println("Already found " + string(key))
		}
		//time.Sleep(time.Millisecond * 10)
	}
}

func delKey(key string, i *int32){
	spec 			:= redis.DefaultSpec()
	client, e := redis.NewSynchClientWithSpec (spec);
	if e != nil { log.Println ("failed to create the client", e); return }
	for{
		
		if *i > MAXKEYS/2{	
			s := strconv.Itoa(int(*i))
			key := key + s
			value, e := client.Get(key);
			if e!= nil { log.Println ("error on Get", e); return }
			if value != nil {
				client.Del(key)
				fmt.Println("Deleted " + key)
				atomic.AddInt32(i, -1)
				runtime.Gosched()
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}



const MAXKEYS int32 = 1000000
var totalkeys int32 = 0

func start_redis () {

	fmt.Printf("\nIntroduce the key to save in Redis:");
	reader := bufio.NewReader(os.Stdin);
	key, _ := reader.ReadString(byte('\n'));
	if len(key) > 1 {
		key = key[0:len(key)-1];
	}

	t0 := time.Now()
	go addKey(string(key), "valorx",&totalkeys)
	go addKey(string(key), "valory",&totalkeys)
	//go delKey(string(key),&totalkeys)

	//wait until done
	<- done
	t1 := time.Now()
	fmt.Printf("Finish one million!! in %v \n", t1.Sub(t0))

}
