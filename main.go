package main

import (
    "fmt"
    //"time"
)

func main() {
  fmt.Println("Choose and option:")
  fmt.Println("\t(1) for Mutext ")
  fmt.Println("\t(2) for Producer Consumer")
  fmt.Println("\t(3) for Redis ")
  fmt.Println("\t(4) for Twitter API ")
  fmt.Println("\t(5) for Redis Chat ")

  var opt string
  fmt.Scanln(&opt)
  switch opt {
    case "1":
      fmt.Println("******* Selected Mutex Example *******")
      start_mutex()
    case "2":
      fmt.Println("******* Selected Producer-Consumer Example *******")
      start_producer_consumer()
    case "3":
      fmt.Println("******* Selected Redis Example *******")
      start_redis()
    case "4":
      fmt.Println("******* Selected Twitter API Example *******")
      start_twitter_api()
    case "5":
      fmt.Println("******* Selected Redis Chat Example *******")
      start_chat()
  }


}