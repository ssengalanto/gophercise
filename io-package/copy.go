package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	s := "some io.Reader stream to be read\n"
	r := strings.NewReader(s)

	if a, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(a, len(s))
	}
}
