package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	SECS_IN_MINUTE = 60
	SECS_IN_HOUR   = SECS_IN_MINUTE * 60
	SECS_IN_DAY    = SECS_IN_HOUR * 24
)

func parseTime(timeStr string) time.Time {
	tm, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	return tm
}

func getSeconds(ts time.Time) int {
	return ts.Hour()*SECS_IN_HOUR + ts.Minute()*SECS_IN_MINUTE + ts.Second()
}

func getTimestampFromSeconds(secs int) string {
	h := secs / SECS_IN_HOUR
	h %= 24
	secs %= SECS_IN_HOUR
	m := secs / SECS_IN_MINUTE
	secs %= SECS_IN_MINUTE
	s := secs

	hs, ms, ss := strconv.Itoa(h), strconv.Itoa(m), strconv.Itoa(s)

	if len(hs) == 1 {
		hs = "0" + hs
	}

	if len(ms) == 1 {
		ms = "0" + ms
	}

	if len(ss) == 1 {
		ss = "0" + ss
	}

	return fmt.Sprintf("%s:%s:%s", hs, ms, ss)
}

func main() {
	var A, B, C string
	fmt.Scan(&A, &B, &C)

	aTS := getSeconds(parseTime(A))
	bTS := getSeconds(parseTime(B))
	cTS := getSeconds(parseTime(C))

	tm2 := 2*bTS + mod(cTS-aTS, SECS_IN_DAY)
	tm := tm2/2 + tm2%2

	fmt.Println(getTimestampFromSeconds(tm))
}

// mod is a helper func that works with negative numbers
func mod(a, b int) int {
	return (a%b + b) % b
}

/*
Tc 2
11:37:00
23:51:00
23:59:59

Output
06:02:30

Tc 9
21:00:00
22:00:00
06:00:00

Output
02:30:00
*/
