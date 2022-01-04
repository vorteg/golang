package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	// read input from the user 5 time and write it to a log
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	//
	go mylogger.ListenerForLog(ch)

	fmt.Println("Enter something")

	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}
