package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments. For example, filesvr \":8899\" ~/Downloads")
		os.Exit(1)
	}
	findLocalIpAddress()
	http.Handle("/", http.FileServer(http.Dir(os.Args[2])))
	fmt.Println(http.ListenAndServe(os.Args[1], nil))
}

func findLocalIpAddress() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("--------------------------")
	for _, inter := range interfaces {
		if inter.Flags&net.FlagUp != 0 && inter.Flags&(net.FlagLoopback|net.FlagPointToPoint) == 0 {
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}

			for i, addr := range addrs {
				if i == 0 {
					fmt.Println(inter.Name, inter.Flags)
				}
				fmt.Println(" ", addr)
			}
		}
	}
	fmt.Println("--------------------------")
}
