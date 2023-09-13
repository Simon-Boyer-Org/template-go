package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		for _, e := range os.Environ() {
			fmt.Println(e)
		}
		time.Sleep(15 * time.Second)
	}
}
