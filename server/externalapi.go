package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiEndPoint = "http://pcc.g0v.ronny.tw/api/searchbytitle"
)

var (
	currReqCount int
)

func requestExternalAPI(qryStr string) {
	var searchByTitle SearchByTitle

	fmt.Printf("[%v] Query: %s\n", time.Now(), qryStr)

	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	params := req.URL.Query()
	params.Add("query", qryStr)
	req.URL.RawQuery = params.Encode()

	mu.Lock()
	currReqCount++
	mu.Unlock()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("%s: %s\n", resp.Status, req.URL)
		return
	}

	mu.Lock()
	currReqCount--
	mu.Unlock()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	if err := json.Unmarshal([]byte(body), &searchByTitle); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("%v: %s\n", searchByTitle.Took, req.URL)
}
