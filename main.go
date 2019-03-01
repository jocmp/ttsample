package main

import (
	"github.com/jocmp/tt"
	"fmt"
)

// KEY where key is the CTA Developer License Agreement key
// Learn more at https://www.transitchicago.com/developers/traintracker/
func main()  {
	// Western Brown Line GTFS Stop ID found at https://www.transitchicago.com/developers/ttdocs/
	westernStopID := 41480
	trainTracker, err := tt.FetchArrivals(KEY, westernStopID)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for _, arrival := range trainTracker.Arrivals {
		fmt.Printf("arrival: run %s at %s\n", arrival.Run, arrival.ArrivalTime)
	}
}
