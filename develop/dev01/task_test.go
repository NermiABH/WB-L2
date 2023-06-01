package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintCurrentTime(t *testing.T) {
	PrintCurrentTime()
	fmt.Println(time.Now())
}
