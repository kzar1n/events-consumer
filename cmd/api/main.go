package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Running API server...")
		time.Sleep(10 * time.Second)
	}

}
