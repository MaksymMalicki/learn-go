package main

import (
	"golang/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
