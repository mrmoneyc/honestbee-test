package main

import (
	"bufio"
	"log"
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
)

func main() {
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	log.Printf("TCP Server listening on %s:%s\n", connHost, connPort)

	defer func() {
		log.Println("Listener Closed")
		l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}

		go requestHandler(conn)
	}
}

func requestHandler(conn net.Conn) {
	log.Printf("Handling new connection: %s...\n", conn.RemoteAddr())

	defer func() {
		log.Printf("Closing connetion: %s...\n", conn.RemoteAddr())
		conn.Close()
	}()

	timeoutDuration := connTimeoutDuration * time.Second
	bufReader := bufio.NewReader(conn)

	for {
		conn.SetDeadline(time.Now().Add(timeoutDuration))

		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			log.Printf("Reading buffer failed: %v\n", err)
			return
		}

		readLine := strings.TrimSuffix(string(bytes), "\n")

		switch readLine {
		case "quit":
			conn.Write([]byte("QUIT\n"))
			log.Println("QUIT")
			return
		default:
			conn.Write([]byte(readLine + "\n"))
		}
	}
}
