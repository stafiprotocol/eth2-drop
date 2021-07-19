package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRootHash(ledgerApi string, round int64) (string, error) {
	realUrl := fmt.Sprintf("%s/api/v1/root_hash?round=%d", ledgerApi, round)
	rsp, err := http.Get(realUrl)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get root hash status: %d", rsp.StatusCode)
	}
	rspBodyBts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if len(rspBodyBts) <= 0 {
		return "", fmt.Errorf("get root hash err %s", fmt.Errorf("body zero err"))
	}
	rspRootHash := RspRootHash{}
	err = json.Unmarshal(rspBodyBts, &rspRootHash)
	if err != nil {
		return "", fmt.Errorf("get root hash err %s", err)
	}
	if rspRootHash.Status != "80000" {
		return "", fmt.Errorf("get root hash err %s", fmt.Errorf("status err:%s", rspRootHash.Status))
	}
	return rspRootHash.Data.RootHash, nil

}

func getDropFLowLatest(ledgerApi string) (string, error) {
	realUrl := fmt.Sprintf("%s/api/v1/drop_flow_latest", ledgerApi)
	rsp, err := http.Get(realUrl)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get getDropFLow status: %d", rsp.StatusCode)
	}
	rspBodyBts, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if len(rspBodyBts) <= 0 {
		return "", fmt.Errorf("get getDropFLow err %s", fmt.Errorf("body zero err"))
	}
	rspDropFlow := RspDropFlow{}
	err = json.Unmarshal(rspBodyBts, &rspDropFlow)
	if err != nil {
		return "", fmt.Errorf("get getDropFLow err %s", err)
	}
	if rspDropFlow.Status != "80000" {
		return "", fmt.Errorf("get getDropFLow status err %s", rspDropFlow.Status)
	}
	return rspDropFlow.Data.DropFlowLatestDate, nil
}

type RspRootHash struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		RootHash string `json:"root_hash"`
	} `json:"data"`
}

type RspDropFlow struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		DropFlowLatestDate string `json:"drop_flow_latest_date"`
	} `json:"data"`
}
