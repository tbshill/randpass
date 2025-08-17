package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, rand.Text())
}
