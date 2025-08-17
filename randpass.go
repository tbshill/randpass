package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

const alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	const maxSize = 10
	chars := []byte(alpha)
	out := [maxSize]byte{}

	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i < maxSize; i++ {
		out[i] = chars[rand.Intn(len(alpha))]
	}
	fmt.Fprintf(os.Stdout, string(out[:]))
}