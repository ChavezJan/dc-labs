// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, timeZone)
		t, err := TimeIn(time.Now(), timeZone)
		if err == nil {
			_, error := io.WriteString(c, t.Location().String()+t.Format("15:04")+"\n")
			if error != nil {
				return
			}
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
	ip := flag.String("port", "m", "the port")
	timeZone := flag.String("TZ", "y", "time zone")

	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*ip)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection: " + *ip)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn, *timeZone) // handle connections concurrently
	}
}
