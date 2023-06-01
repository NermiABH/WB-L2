package main

import (
	"WB-L2/develop/dev10/telnet"
	"flag"
	"log"
	"os"
	"time"
)

var flagTimeout = flag.String("timeout", "10s", "Timeout connecting to server")

func init() {
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("should be host and port")
	}
	timeout, err := time.ParseDuration(*flagTimeout)
	if err != nil {
		log.Fatal(err)
	}
	config := &telnet.Config{Timeout: timeout, Host: args[0], Port: args[1]}
	conn, err := telnet.NewConn(os.Stdin, os.Stdout, config)
	if err != nil {
		log.Fatal(err)
	}
	conn.Run()
}
