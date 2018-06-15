package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	connHost            = "localhost"
	connPort            = "9527"
	connType            = "tcp"
	connTimeoutDuration = 10
	reqLimitPerSec      = 30
	reqRate             = time.Second / reqLimitPerSec
)

var (
	mu           sync.RWMutex
	processedReq int
	currClient   []string
	qryStr       = make(chan string, 100)
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

	throttle := time.Tick(reqRate)

	go func() {
		for q := range qryStr {
			<-throttle
			go requestExternalAPI(q)

			mu.Lock()
			processedReq++
			mu.Unlock()
		}
	}()

	go startAPIServer()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		go requestHandler(conn)
	}
}

func requestHandler(conn net.Conn) {
	fmt.Printf("Handling new connection: %s...\n", conn.RemoteAddr())

	mu.Lock()
	currClient = append(currClient, conn.RemoteAddr().String())
	mu.Unlock()

	defer func() {
		fmt.Printf("Closing connetion: %s...\n", conn.RemoteAddr())

		mu.Lock()
		for k, v := range currClient {
			if v == conn.RemoteAddr().String() {
				currClient = currClient[:k+copy(currClient[k:], currClient[k+1:])]
			}
		}
		mu.Unlock()

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
