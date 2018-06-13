package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const (
	connHost            = "localhost"
	connPort            = "9527"
	connType            = "tcp"
	connTimeoutDuration = 10
	reqRate             = time.Second / 30
)

func main() {
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("TCP Server listening on %s:%s\n", connHost, connPort)

	defer func() {
		fmt.Println("Listener Closed")
		l.Close()
	}()

	var currConnection int
	var processedReq int
	qryStr := make(chan string, 100)
	throttle := time.Tick(reqRate)

	go func() {
		for q := range qryStr {
			<-throttle
			go requestExternalAPI(q)
			processedReq++
		}
	}()

	go startAPIServer(qryStr, &currConnection, &processedReq)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		go requestHandler(conn, qryStr, &currConnection)
	}
}

func requestHandler(conn net.Conn, qryStr chan<- string, currConnection *int) {
	fmt.Printf("Handling new connection: %s...\n", conn.RemoteAddr())
	*currConnection++

	defer func() {
		fmt.Printf("Closing connetion: %s...\n", conn.RemoteAddr())
		*currConnection--
		conn.Close()
	}()

	timeoutDuration := connTimeoutDuration * time.Second
	bufReader := bufio.NewReader(conn)

	for {
		conn.SetDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Reading buffer failed: %v\n", err)
			return
		}

		readLine := strings.TrimSuffix(string(bytes), "\n")

		switch readLine {
		case "quit":
			conn.Write([]byte("QUIT\n"))
			fmt.Println("QUIT")
			return
		default:
			conn.Write([]byte(readLine + "\n"))
			qryStr <- readLine
		}
	}
}
