package socket

import (
	// "fmt"
	"log"
	"net"
	// "os"
)

func ListenOn(port string) net.Listener {
	ln, err := net.Listen("tcp", ":8070")
	checkErr(err)

	return ln
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
