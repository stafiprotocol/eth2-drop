// Copyright 2021 stafiprotocol
// SPDX-License-Identifier: LGPL-3.0-only

package main

import "testing"

func TestGetDropFLowLatest(t *testing.T){
	l,err:=getDropFLowLatest("http://127.0.0.1:8082")
	t.Fatal(err,l)
}