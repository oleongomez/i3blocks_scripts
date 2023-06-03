package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	fmt.Fprintf(os.Stdout, "  %02d/%02d/%d   %02d:%02d", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute())
	os.Exit(0)
}
