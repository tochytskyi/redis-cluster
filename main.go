package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	redisInit "github.com/tochytskyi/redis-cluster/src/redis"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	log.Println("Running app")

	redisInit.Init()

	//currentValue, err := redis.GetInstance().Get("users").Result()
	//if err != nil {
	//	log.Println("No data in cache", err)
	//}

	<-done
	log.Println("exiting app")
}
