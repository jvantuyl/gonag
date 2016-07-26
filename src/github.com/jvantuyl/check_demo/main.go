package main

import "github.com/jvantuyl/gonag"

func main() {
    gonag.ReturnNagiosCheck(gonag.StatWarn, "OK - foo", "1,2,3\n4\n5")
}
