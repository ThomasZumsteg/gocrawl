package main

import (
    "./gocrawl"
)

func main() {
    dev := gocrawl.NewDevice("name", "ip")
    dev.GetUserAndPassword()
}
