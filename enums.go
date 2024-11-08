package main

import "fmt"

type responseStatus int

const (
	ok = iota
	noContent
	notFound
)

var statusName = map[responseStatus]string{
	ok:        "Ok",
	noContent: "No content",
	notFound:  "Not found",
}

func task() responseStatus {
	// logika
	return ok
}

type response struct {
	body   string
	status responseStatus
}

func enums() {
	status := notFound

	switch status {
	case ok, noContent:
		fmt.Println("Success")
	case notFound:
		fmt.Printf("Failure: %v", statusName[responseStatus(status)])
	}
}
