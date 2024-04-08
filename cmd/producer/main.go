package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Running producer...")
		time.Sleep(10 * time.Second)
	}

}
