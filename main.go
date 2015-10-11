package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type thing string

var setThing = make(chan thing)
var getThing = make(chan thing)

func thingUpdater() {
	var t thing
	log.Print("thingUpdater: running")
	for {
		select {
		case t = <-setThing:
		case getThing <- t:
		}
	}
}

func updateThing(w http.ResponseWriter, r *http.Request) {
	setThing <- thing(fmt.Sprintf("%v", time.Now()))
}

func readThing(w http.ResponseWriter, r *http.Request) {
	v := <-getThing
	fmt.Fprint(w, v)
}

func main() {
	go thingUpdater()
	setThing <- thing(fmt.Sprintf("%v", time.Now()))

	http.HandleFunc("/read", readThing)
	http.HandleFunc("/update", updateThing)

	http.ListenAndServe(":8080", nil)
}
