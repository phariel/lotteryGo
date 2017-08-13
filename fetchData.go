package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jeffail/gabs"
	"github.com/phariel/couchGo"
)

type loResult struct {
	Expect   string `json:"expect"`
	Opencode string `json:"opencode"`
	Opentime string `json:"opentime"`
}

func fetchDataToDb() {
	req, _ := http.NewRequest("GET", "http://f.apiplus.net/ssq-20.json", bytes.NewBufferString(""))
	client := &http.Client{}
	resp, _ := client.Do(req)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	lo, _ := gabs.ParseJSON(buf.Bytes())
	loData, _ := lo.S("data").Children()
	var requests []couchGo.RequestData
	for _, loDataItem := range loData {
		expect, _ := loDataItem.S("expect").Data().(string)
		opencode, _ := loDataItem.S("opencode").Data().(string)
		opentime, _ := loDataItem.S("opentime").Data().(string)

		jsonStruct := &loResult{
			Expect:   expect,
			Opencode: opencode,
			Opentime: opentime}
		jsonBody, _ := json.Marshal(jsonStruct)
		requests = append(requests, couchGo.RequestData{ID: expect, JSONBody: jsonBody})
	}
	couchGo.Update(requests, func(res []couchGo.ResponseData) {
		fmt.Println(res)
	})
}
