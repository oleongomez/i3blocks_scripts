package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	connected_fle, err := os.ReadFile("/sys/class/power_supply/AC/online")
	logError(err)
	status_file, err := os.ReadFile("/sys/class/power_supply/BAT0/status")
	logError(err)
	capacity_file, err := os.ReadFile("/sys/class/power_supply/BAT0/capacity")
	logError(err)
	power_is_connected := strings.TrimRight(string(connected_fle), "\n")
	status := strings.TrimRight(string(status_file), "\n")
	capacity := strings.TrimRight(string(capacity_file), "\n")

	int_capacity, err := strconv.Atoi(capacity)
    logError(err)
	int_connected, err := strconv.Atoi(power_is_connected)
    logError(err)

    connected_chars := map[int]string{0: "", 1: ""}

	logError(err)
	var battery_capacity = ""
	if int_capacity > 0 && int_capacity <= 25 {
		battery_capacity = ""
	}

	if int_capacity > 25 && int_capacity <= 50 {
		battery_capacity = ""
	}

	if int_capacity > 50 && int_capacity <= 75 {
		battery_capacity = ""
	}

	if int_capacity > 75 && int_capacity <= 100 {
		battery_capacity = ""
	}

	fmt.Fprintf(os.Stdout, "%s  %s %s", connected_chars[int_connected], status, battery_capacity)
	os.Exit(0)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
