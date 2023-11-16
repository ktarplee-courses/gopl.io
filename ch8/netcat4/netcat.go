// See page 227.

// Netcat is a simple read/write client for TCP servers.
// Symmetric example
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn) //nolint
		log.Println("done STDOUT -> TCP")
		conn.Close()
		done <- struct{}{} // signal the main goroutine
	}()

	go func() {
		io.Copy(conn, os.Stdin) //nolint
		log.Println("done STDIN -> TCP")
		conn.Close()
		done <- struct{}{} // signal the main goroutine
	}()

	<-done // wait for background goroutine to finish
	<-done // wait for background goroutine to finish
}
