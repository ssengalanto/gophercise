package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "Golang is cool!")

}
