package main

import (
    "fmt"
    "log"
    "github.com/AressS-Git/syspulse/internal/agent"
)

func main() {
    var stats, err = agent.GetMetrics()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(stats)
}