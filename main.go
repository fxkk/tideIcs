package main

import (
	"fmt"
	"os"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/fxkk/tideIcs/bshe"
	"github.com/google/uuid"
)

func tideReportToIcs(tideReport bshe.TideReport) *ics.Calendar {
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodRequest)

	for _, dp := range tideReport.Data {
		eventUUID := uuid.New()
		event := cal.AddEvent(eventUUID.String())

		event.SetStartAt(dp.Timestamp)
		event.SetEndAt(dp.Timestamp.Add(30 * time.Minute))
		event.SetCreatedTime(time.Now())
		event.SetDtStampTime(time.Now())

		if dp.TidePhase == 'H' {
			event.SetSummary("Hochwasser")
		} else {
			event.SetSummary("Niedrigwasser")
		}
	}

	return cal
}

func main() {
	if len(os.Args) < 2 {
		println("Usage: tideics <filePath>")
		os.Exit(1)
	}

	if os.Args[1] == "--help" {
		println("Usage: tideics <filePath>")
		println("Generate an iCalendar file from tide bsh txt data")
		os.Exit(0)
	}

	filePath := os.Args[1]

	tideReport, err := bshe.Load(filePath)
	if err != nil {
		panic(err)
	}

	cal := tideReportToIcs(tideReport)
	icsData := cal.Serialize()

	fmt.Print(icsData)
}
