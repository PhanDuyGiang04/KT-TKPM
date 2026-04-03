package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World! This is a simple Go application running in a Docker container.")
	fmt.Println("This application will run indefinitely until it is stopped.")

	// Giữ container chạy để chúng ta có thể kiểm tra
	for {
		time.Sleep(1 * time.Hour)
	}
}