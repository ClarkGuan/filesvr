package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments. For example, filesvr \":8899\" ~/Downloads")
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.Dir(os.Args[2])))
	fmt.Println(http.ListenAndServe(os.Args[1], nil))
}
