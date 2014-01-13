package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"io/ioutil"
)

const (
	empty_filename = "<filename>"
	mllp_start = 0x0b
	mllp_end1  = 0x1c
	mllp_end2  = 0x0d
)

func main() {
	host := flag.String("host", "localhost", "hostname of MLLP server, default value is localhost")
	port := flag.Int("port", 2575, "portnumber of MLLP server, default value is 2575")
	file := flag.String("file", empty_filename, "path to file which contents will be send to the MLLP server")
	flag.Parse()

	f := *file
	if f == empty_filename {
		fmt.Println("'file' argument is required")
		os.Exit(1)
	}
	Send(file, host, port)
}

func Send(file *string, host *string, port *int) {
	fmt.Printf("Sending message in file %s over MLLP to %s:%d\n", *file, *host, *port)

	// constuct connection
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		fmt.Println("Connection failed:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// write the actual message
	conn.Write([]byte { mllp_start })
	fmt.Fprintf(conn, readfile(file))
	conn.Write([]byte { mllp_end1 })
	conn.Write([]byte { mllp_end2 })

	// read response
	reply := make([]byte, 1024)
	_, err = conn.Read(reply)
	if err != nil {
		fmt.Println("Handling server reply failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func readfile(file *string) string {
	content, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println("Reading file failed:", err.Error())
		os.Exit(1)
	}
	return string(content)
}
