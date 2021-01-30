package main

import (
	"io"
	"log"
	"net"
	"os"
)

/**
* @Author: super
* @Date: 2020-09-23 08:40
* @Description:
**/

func main() {
	conn, err := net.Dial("tcp", ":2020")
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
