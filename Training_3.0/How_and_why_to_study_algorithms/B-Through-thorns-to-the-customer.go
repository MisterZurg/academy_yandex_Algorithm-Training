package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	eventsById := make(map[int][]int)

	var day, hour, minute, rid int
	var event string
	var eventType int
	for i := 0; i < n; i++ {
		fmt.Scan(&day, &hour, &minute, &rid, &event)

		switch event {
		case "src":
			eventType = 0
		case "C":
			eventType = 1
		case "S":
			eventType = 1
		case "B":
			continue
		}
		hour = day*24 + hour
		minute = hour*60 + minute

		if _, ok := eventsById[rid]; !ok {
			eventsById[rid] = make([]int, 0)
		}
		eventsById[rid] = append(eventsById[rid], minute, eventType)
	}

	// sort keys
	eventsByIdkeys := make([]int, 0, len(eventsById))
	for k := range eventsByIdkeys {
		eventsByIdkeys = append(eventsByIdkeys, k)
	}
	sort.Ints(eventsByIdkeys)

	for id := range eventsByIdkeys {
		// eventsById[id].sort()
		ans := 0
		evt0, evt1 := eventsById[id][0], eventsById[id][1]
		var prevTime int
		if evt1 == 0 {
			prevTime = evt0
		} else if evt1 == 1 {
			ans += evt0 - prevTime
		}
		fmt.Println("SEX")
		fmt.Printf("%d ", ans)
	}
}
