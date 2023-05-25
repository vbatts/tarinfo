package main

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	fh, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	tr := tar.NewReader(fh)
	for {
		hdr, err := tr.Next()
		if err != nil {
			break
		}
		buf, err := json.MarshalIndent(hdr, "", "  ")
		if err != nil {
			log.Println("failed to marshal")
			continue
		}
		fmt.Println(string(buf))
		//fmt.Printf("%#v\n", hdr)
	}
}
