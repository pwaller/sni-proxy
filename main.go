package main

import (
	"io"
	"log"
	"net"

	"github.com/polvi/sni"
)

func main() {

	listener, err := net.Listen("tcp", ":8443")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection")
		}
		go serve(c)
	}
}

func serve(c net.Conn) {

	name, connIn, err := sni.ServerNameFromConn(c)
	if err != nil {
		panic(err)
	}

	log.Printf("Got connection for %q", name)

	connOut, err := net.Dial("tcp", "localhost:9443")
	if err != nil {
		panic(err)
	}

	go io.Copy(connIn, connOut)
	go io.Copy(connOut, connIn)

}
