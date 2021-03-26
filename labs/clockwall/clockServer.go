package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Input struct {
	port     int
	timezone string
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	loc, _ := time.LoadLocation(timeZone)
	for {
		time_now := time.Now().In(loc).Format("15:04:05\n")
		response := timeZone + " " + time_now
		_, err := io.WriteString(c, response)

		if err != nil {
			return
		} else {
			_, error := io.WriteString(c, timeZone+" TIMEZONE NOT AVAILABLE")
			if error != nil {
				return
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {

	var input Input
	var host string

	input = manageInput()
	host = "0.0.0.0:" + fmt.Sprintf("%d", input.port)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(host)
	for {
		conn, err := listener.Accept()
		log.Print(conn)
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, input.timezone)
	}
}

func manageInput() Input {

	var input Input
	var tmpPort *int

	input.timezone = os.Getenv("TZ")

	tmpPort = flag.Int("port", 9000, "port number.")
	flag.Parse()
	input.port = *tmpPort

	return input

}
