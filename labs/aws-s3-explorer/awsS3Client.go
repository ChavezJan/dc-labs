package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var proxy = flag.String("proxy", "localhost:8000", "Proxy url.")
	var bucketName = flag.String("bucket", "", "S3 bucket name.")
	var directory = flag.String("directory", "", "Directory name.")
	flag.Parse()

	// error if bucket parameter is missing, the server cannot process anything
	if *bucketName == "" {
		fmt.Println("ERROR - Missing parameters.")
		return
	}

	// request to server the bucket and directory through a url
	request := fmt.Sprintf("http://%v/example?bucket=%v&dir=%v", *proxy, *bucketName, *directory)
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// prints the response from server
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
