package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	redisCache "github.com/tochytskyi/redis-cluster/src/redis"
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

	redisCache.Init()

	redisCache.GetOrRefresh("testKey2", int(time.Second)*10)

	<-done
	log.Println("exiting app")
}
