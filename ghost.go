package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "8.8.8.8:53" // this is googles DNS
	fmt.Println("Targeting:", target)

	for i := 1; i <= 5; i++ {
		start := time.Now()
		conn, err := net.DialTimeout("tcp", target, 2*time.Second)
		duration := time.Since(start)

		if err != nil {
			fmt.Printf("Shot %d: MISSED (Error: %v)\n", i, err)
		} else {
			fmt.Printf("Shot %d: HIT! Latency: %v\n", i, duration)
			conn.Close()
		}
		time.Sleep(1 * time.Second)
	}
}