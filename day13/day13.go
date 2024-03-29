package main

import (
	"fmt"
	"github.com/koenaad/Advent-of-Code-2017/day13/firewall"
)

func main() {
	fmt.Println("Advent of Code 2017 - day 13")

	f := firewall.Init(input)
	fmt.Printf("Severity accrued after stepping through firewall is %v\n", f.SeverityAccruedAfterSteppingThrough())
	fmt.Printf("Delay needed to avoid getting caught by the firewall is %v\n", delayNeededToAvoidGettingCaught(&f))
}

func delayNeededToAvoidGettingCaught(f *firewall.Firewall) int {
	delay := 0

	for {
		if f.CanStepThroughWithoutGettingCaughtAfter(delay) {
			return delay
		}
		delay += 1
	}
}

var input = `0: 3
1: 2
2: 4
4: 8
6: 5
8: 6
10: 6
12: 4
14: 6
16: 6
18: 9
20: 8
22: 8
24: 8
26: 8
28: 10
30: 8
32: 12
34: 10
36: 14
38: 12
40: 12
42: 12
44: 12
46: 12
48: 12
50: 14
52: 12
54: 14
56: 12
58: 12
60: 14
62: 18
64: 14
68: 14
70: 14
72: 14
74: 14
78: 14
80: 20
82: 14
84: 14
90: 17`
