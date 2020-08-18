package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const (
	emptyFilename  = "<filename>"
	emptyDirectory = "<empty_dir>"
	mllpStart      = 0x0b
	mllpEnd        = 0x1c
	mllpEnd2       = 0x0d
)

func main() {
	host := flag.String("host", "localhost", "hostname of MLLP server, default value is localhost")
	port := flag.Int("port", 2575, "portnumber of MLLP server, default value is 2575")
	file := flag.String("file", emptyFilename, "path to file which contents will be send to the MLLP server")
	dir := flag.String("dir", emptyDirectory, "path to file which contents will be send to the MLLP server")
	flag.Parse()

	d := *dir
	f := *file

	if d == emptyDirectory && f == emptyFilename {
		fmt.Println("A file or directory must be provided. Exiting.")
		os.Exit(1)
	}

	if f != emptyFilename {
		Send(file, host, port)
	}

	if d != emptyDirectory {
		SendDir(dir, host, port)
	}

}

//Send sends a file over MLLP
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
	conn.Write([]byte{mllpStart})
	fmt.Fprintf(conn, readfile(file))
	conn.Write([]byte{mllpEnd})
	conn.Write([]byte{mllpEnd2})

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

// SendDir sends all files in a directory over MLLP
func SendDir(dir *string, host *string, port *int) {
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		file := fmt.Sprintf("%s/%s", *dir, f.Name())
		Send(&file, host, port)
	}

}
