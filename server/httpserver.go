package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
)

const (
	apiServeHost = "localhost"
	apiServePort = "9528"
)

type Stat struct {
	CurrConnCount  int      `json:"current_connection_count"`
	CurrReqRate    float64  `json:"current_request_rate"`
	ProcessedReq   int      `json:"processed_request_count"`
	CurrReqCount   int      `json:"current_request_count"`
	RemainingJobs  int      `json:"remaining_jobs"`
	CurrGoRoutine  int      `json:"current_goroutine_count"`
	CurrConnClient []string `json:"current_connected_client"`
}

func startAPIServer(qryStr chan<- string) {
	http.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		stat := &Stat{
			CurrConnCount:  len(currClient),
			CurrReqRate:    float64(currReqCount) / float64(reqLimitPerSec),
			ProcessedReq:   processedReq,
			CurrReqCount:   currReqCount,
			RemainingJobs:  len(qryStr),
			CurrGoRoutine:  runtime.NumGoroutine(),
			CurrConnClient: currClient,
		}
		mu.RUnlock()

		b, err := json.Marshal(stat)
		if err != nil {
			fmt.Printf("json marshal failed: %v", err)
		}

		fmt.Fprintf(w, "%s", b)
	})

	fmt.Printf("HTTP Server listening on %s:%s\n", apiServeHost, apiServePort)
	err := http.ListenAndServe(apiServeHost+":"+apiServePort, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
