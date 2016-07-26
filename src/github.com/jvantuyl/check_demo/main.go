package main

import "fmt"
import "os"
import "strings"

const (
    STAT_OK = iota
    STAT_WARN = iota
    STAT_CRIT = iota
    STAT_UNKN = iota
)

type Status int

func send_strings(s []string, c chan string) {
    defer close(c)
    for i, t := range s {
      if (i + 1 < len(s)) || (len(t) > 0) {
          c <- t
      }
    }
}

func return_check(stat Status, output string, perfdata string) {
    outputs := make(chan string)
    perfs := make(chan string)
    go send_strings(strings.Split(output, "\n"), outputs)
    go send_strings(strings.Split(perfdata, "\n"), perfs)

    text0, have_text0 := <- outputs
    if (!have_text0) { panic("no output") }
    perf0, have_perf0 := <- perfs
    if (have_text0) {
        if (have_perf0) {
            fmt.Printf("%s|%s", text0, perf0)
        } else {
            fmt.Printf("%s", text0)
        }
    } else {
        panic("no text?")
    }
    perf1, have_perf1 := <- perfs
    out := false
    for text := range outputs {
        out = true
        fmt.Printf("\n%s", text)
    }
    if (have_perf1) {
        if (!out) { fmt.Print("\n") }
        fmt.Printf("|%s\n", perf1)
        for perf := range perfs {
            fmt.Printf("%s\n", perf)
        }
    } else {
      fmt.Print("\n")
    }
    os.Exit(int(stat))
}

func main() {
    return_check(STAT_WARN, "OK - foo", "1,2,3\n4\n5")
}
