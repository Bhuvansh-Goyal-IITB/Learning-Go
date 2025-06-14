package main

import (
	"learning-go/17/svg"
	"os"
	"time"
)

func main() {
	t := time.Now().Local()
	svg.Write(os.Stdout, t)
}
