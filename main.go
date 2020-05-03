/*package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "github.com/saipranav/getstarted-grpc-go/store"
)

func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(b)
	}
}

func storeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		/*
			en := &store.Entity{}
			err = protojson.Unmarshal(b, en)
			if err != nil {
				fmt.Println(err)
			}

			data, err := proto.Marshal(en)
			if err != nil {
				fmt.Println(err)
			}
		*/

		w.WriteHeader(http.StatusAccepted)
		w.Write(data)
	}
}

func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/store", storeHandler)
	http.ListenAndServe(":8080", nil)
}
*/