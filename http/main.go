package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootfunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not existed\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("It works!\n"))
}

func main() {

	http.HandleFunc("/", rootfunc)

	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
