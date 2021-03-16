package main

import (
	"io"
	"log"
	"net"
	"os"
)

func createChannel(conn io.Reader) chan int {
	done := make(chan int)
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- 1 // signal the main goroutine
	}()
	return done
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8010")
	if err != nil {
		log.Fatal(err)
	}

	conn1, err := net.Dial("tcp", "localhost:8020")
	if err != nil {
		log.Fatal(err)
	}

	conn2, err := net.Dial("tcp", "localhost:8030")
	if err != nil {
		log.Fatal(err)
	}
	var servers [3]chan int
	for i := range servers {
		servers[i] = make(chan int)
	}
	servers[0] = createChannel(conn)
	servers[1] = createChannel(conn1)
	servers[2] = createChannel(conn2)
	for i := range servers {
		x := 1
		x = <-servers[i] // wait for background goroutine to finish
		log.Println("Channel Closed with value: ", x)
		close(servers[i])
	}
}
