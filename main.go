package main

import (
	"fmt"
	"os"
	"os/signal"
	"simple-redis-go/cmd"
	"simple-redis-go/config"
	"syscall"
)

func main() {
	fmt.Print("Start...")
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		fmt.Println("Stop...")
		os.Exit(0)
	}()
	config.LoadConfig(".")
	cmd.Execute()
}
