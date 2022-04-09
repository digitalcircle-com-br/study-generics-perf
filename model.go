package main

import "time"

type SomeReq1 struct {
	A   string    `json:"a"`
	B   time.Time `json:"b"`
	C   int64     `json:"c"`
	V1  string    `json:"v1"`
	Src string    `json:"src"`
}
type SomeRes1 struct {
	A  string    `json:"a"`
	B  time.Time `json:"b"`
	C  int64     `json:"c"`
	V1 string    `json:"v1"`
}

type SomeReq2 struct {
	A  string    `json:"a"`
	B  time.Time `json:"b"`
	C  int64     `json:"c"`
	V2 string    `json:"v2"`
}
type SomeRes2 struct {
	A  string    `json:"a"`
	B  time.Time `json:"b"`
	C  int64     `json:"c"`
	V2 string    `json:"v2"`
}
