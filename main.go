package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var thing string
var mutex = &sync.Mutex{}

func updateThing(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	thing = fmt.Sprintf("%v\n", time.Now())
}

func readThing(w http.ResponseWriter, r *http.Request) {
	v := thingProvider()
	fmt.Fprintf(w, v)
}

func thingProvider() string {
	mutex.Lock()
	defer mutex.Unlock()
	c := thing
	return c
}

func main() {
	thing = fmt.Sprintf("%v\n", time.Now())

	http.HandleFunc("/read", readThing)
	http.HandleFunc("/update", updateThing)

	http.ListenAndServe(":8080", nil)
}
