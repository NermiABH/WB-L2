package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

func PrintCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(time)
	}
	fmt.Println(time)
}

func main() {
	PrintCurrentTime()
}
