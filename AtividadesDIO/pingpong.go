package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan struct{})
	pongChan := make(chan struct{})

	go func() {
		for {
			<-pingChan
			fmt.Println("ping")
			time.Sleep(1 * time.Second)
			pongChan <- struct{}{}
		}
	}()

	go func() {
		for {
			<-pongChan
			fmt.Println("pong")
			time.Sleep(1 * time.Second)
			pingChan <- struct{}{}
		}
	}()

	pingChan <- struct{}{} // Inicie o ciclo enviando um sinal para pingChan

	select {}
}
