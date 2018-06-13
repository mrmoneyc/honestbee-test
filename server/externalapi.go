package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	apiEndPoint = "http://pcc.g0v.ronny.tw/api/searchbytitle"
)

func requestExternalAPI(qryStr string) {
	var searchByTitle SearchByTitle

	log.Printf("[%v] Query: %s\n", time.Now(), qryStr)

	req, err := http.NewRequest("GET", apiEndPoint, nil)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	params := req.URL.Query()
	params.Add("query", qryStr)
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s: %s\n", resp.Status, req.URL)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	if err := json.Unmarshal([]byte(body), &searchByTitle); err != nil {
		log.Printf("%v\n", err)
		return
	}

	fmt.Printf("%v: %s\n", searchByTitle.Took, req.URL)
}
