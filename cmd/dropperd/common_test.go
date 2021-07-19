package main

import "testing"

func TestGetDropFLowLatest(t *testing.T){
	l,err:=getDropFLowLatest("http://127.0.0.1:8082")
	t.Fatal(err,l)
}