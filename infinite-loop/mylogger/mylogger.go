package mylogger

import "log"

func ListenerForLog(ch chan string) {
	for {
		msg := <-ch
		log.Println(msg)
	}
}
