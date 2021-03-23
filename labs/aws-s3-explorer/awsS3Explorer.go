package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Bucket struct {
	fileName xml.Name   `xml:"Bucket"`
	Contents []Contents `xml:"Contents"`
	Name     string     `xml:"Name"`
}

type Contents struct {
	XMLName      xml.Name `xml:"Contents"`
	Key          string   `xml:"Key"`
	LastModified string   `xml:"LastModified"`
	ETag         string   `xml:"ETag"`
	Size         string   `xml:"Size"`
	StorageClass string   `xml:"StorageClass"`
}

type Ext struct {
	extension string
	num       int
}

type Info struct {
	nameBucket      string
	numObjs, numDir int
	extensions      map[string]int
}

func main() {
	url := flag.String("bucket", "m", "the bucket url")
	flag.Parse()
	resp, err := http.Get("https://" + *url + ".s3.amazonaws.com")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	var contents Bucket
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(body, &contents)
	info := Info{*url, 0, 0, make(map[string]int)}
	for _, current := range contents.Contents {
		file := strings.Split(current.Key, ".")
		if len(file) > 1 {
			info.extensions[file[len(file)-1]]++
			info.numObjs++
		} else {
			info.numDir++
		}
	}
	keys := make([]string, 0, len(info.extensions))

	for key := range info.extensions {
		keys = append(keys, key)
	}
	fmt.Printf("AWS S3 Explorer\n")
	fmt.Printf("Bucket Name            : %v\n", info.nameBucket)
	fmt.Printf("Number of objects      : %v\n", info.numObjs)
	fmt.Printf("Number of directories  : %v\n", info.numDir)
	fmt.Printf("Extensions             : ")
	for idx, obj := range keys {
		fmt.Printf("%s(%v)", obj, info.extensions[obj])
		if idx+1 < len(keys) {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}
