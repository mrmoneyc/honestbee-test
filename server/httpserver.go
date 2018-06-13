package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiServeHost = "localhost"
	apiServePort = "9528"
)

type Stat struct {
	CurrConnCTR   int `json:"current_connection_count"`
	CurrReqRate   int `json:"current_request_rate"`
	ProcessedReq  int `json:"processed_request_count"`
	RemainingJobs int `json:"remaining_jobs"`
}

func startAPIServer(qryStr chan<- string) {
	http.HandleFunc("/stat", func(w http.ResponseWriter, r *http.Request) {
		stat := &Stat{RemainingJobs: len(qryStr)}
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
