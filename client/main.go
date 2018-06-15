package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	connType        = "tcp"
	produceDuration = 1000 * time.Millisecond
	workerBootDelay = 100 * time.Millisecond
)

var (
	connHost    string
	connPort    string
	numOfWorker int
	qryElements = []string{"淡江大學", "測試", "平台", "test", "quit"}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&connHost, "h", "localhost", "TCP server address")
	flag.StringVar(&connPort, "p", "9527", "TCP server port")
	flag.IntVar(&numOfWorker, "w", 60, "Concurrent client quantity")
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup

	for w := 1; w <= numOfWorker; w++ {
		wg.Add(1)
		go requestWorker(&wg, w)
		time.Sleep(workerBootDelay)
	}

	wg.Wait()
	log.Println("All worker finished")
}

func requestWorker(wg *sync.WaitGroup, w int) {
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		log.Printf("[Worker %d] %s\n", w, err.Error())
		return
	}

	for _, v := range qryElements {
		log.Printf("[Worker %d] %s\n", w, v)
		fmt.Fprintf(conn, v+"\n")

		respMessage, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("Reading string failed: %s\n", err.Error())
			continue
		}

		readLine := strings.TrimSuffix(respMessage, "\n")

		switch readLine {
		case "quit":
			fmt.Println("QUIT")
			return
		}

		time.Sleep(produceDuration)
	}

	wg.Done()
}
