package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%v\n", time.Now().Format("2006-01-02T15:04:05+08:00"))
	return
}
