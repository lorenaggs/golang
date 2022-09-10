package main

import (
	"io"
	"log"
	"net"
	"os"
)

/*

x = <-ch 	// a receive expression in an assignment statement
	job <- chan int  	(Param fn )
ch <- x 	// a send statement
	result chan <- int 	(Param fn)
<-ch 		// a receive statement; result is discarded
*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("conecction closed :::: ")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
