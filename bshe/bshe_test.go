package bshe

import (
	"fmt"
	"testing"
)

func TestParseData(t *testing.T) {
	dataLines := []string{
		"VB2#DE__506P# #H#Mi# 1. 1.2025# 1:35# 6.73 # #  1#+ 1:00#  26466#   1#2460676.524368#",
		"VB2#DE__506P# #N#Mi#11.12.2025# 8:30# 3.56 # #345#+ 1:00#  26466#   2#2460676.812754#",
	}
	dataPoints, err := parseData(dataLines)
	if err != nil {
		t.Errorf("parseData() error = %v, wantErr %v", err, nil)
	}
	if len(dataPoints) != 2 {
		t.Errorf("parseData() len = %v, want %v", len(dataPoints), 2)
	}
	if dataPoints[0].LineCode != "VB2" {
		t.Errorf("parseData() LineIdentifier = %v, want %v", dataPoints[0].LineCode, "VB2")
	}
	if dataPoints[0].GaugeNumber != "DE__506P" {
		t.Errorf("parseData() GaugeNumber = %v, want %v", dataPoints[0].GaugeNumber, "DE__506P")
	}
	if dataPoints[0].TidePhase != 'H' {
		t.Errorf("parseData() TidePhase = %v, want %v", dataPoints[0].TidePhase, 'H')
	}
	if dataPoints[1].TidePhase != 'N' {
		t.Errorf("parseData() TidePhase = %v, want %v", dataPoints[1].TidePhase, 'N')
	}
	if dataPoints[0].DayOfWeek != "Mi" {
		t.Errorf("parseData() DayOfWeek = %v, want %v", dataPoints[0].DayOfWeek, "Mi")
	}
	if dataPoints[0].Timestamp.Format("2.1.2006-15:04") != "1.1.2025-00:35" {
		t.Errorf("parseData() Date = %v, want %v", dataPoints[0].Timestamp.Format("2.1.2006-15:04"), "1.1.2025-00:35")
	}
	if dataPoints[1].Timestamp.Format("2.1.2006-15:04") != "11.12.2025-07:30" {
		t.Errorf("parseData() Date = %v, want %v", dataPoints[1].Timestamp.Format("2.1.2006-15:04"), "11.12.2025-07:30")
	}
	if fmt.Sprintf("%.2f", dataPoints[0].Height) != "6.73" {
		t.Errorf("parseData() Height = %v, want %v", fmt.Sprintf("%.2f", dataPoints[0].Height), "6.73")
	}
	if dataPoints[0].UTCTimeDiff.Format("+15:04") != "+01:00" {
		t.Errorf("parseData() UTCTimeDiff = %v, want %v", dataPoints[0].UTCTimeDiff.Format("-07:00"), "+01:00")
	}
	if dataPoints[0].NumberOfDay != 1 {
		t.Errorf("parseData() NumberOfDay = %v, want %v", dataPoints[0].NumberOfDay, 1)
	}
	if dataPoints[1].NumberOfDay != 345 {
		t.Errorf("parseData() NumberOfDay = %v, want %v", dataPoints[1].NumberOfDay, 11)
	}
	if dataPoints[0].TransitNumber != 26466 {
		t.Errorf("parseData() TransitNumber = %v, want %v", dataPoints[0].TransitNumber, 26466)
	}
	if dataPoints[1].TransitNumber != 26466 {
		t.Errorf("parseData() TransitNumber = %v, want %v", dataPoints[1].TransitNumber, 26466)
	}
	if dataPoints[0].Transit != 1 {
		t.Errorf("parseData() Transit = %v, want %v", dataPoints[0].Transit, 1)
	}
	if dataPoints[1].Transit != 2 {
		t.Errorf("parseData() Transit = %v, want %v", dataPoints[1].Transit, 2)
	}
	if fmt.Sprintf("%.6f", dataPoints[0].JulianDate) != "2460676.524368" {
		t.Errorf("parseData() JulianDate = %v, want %v", fmt.Sprintf("%.6f", dataPoints[0].JulianDate), "2460676.524368")
	}
	if fmt.Sprintf("%.6f", dataPoints[1].JulianDate) != "2460676.812754" {
		t.Errorf("parseData() JulianDate = %v, want %v", fmt.Sprintf("%.6f", dataPoints[1].JulianDate), "2460676.812754")
	}

}

func TestParseHeader(t *testing.T) {
	headerLines := []string{
		"A04#GT-Name    :#Cuxhaven, Steubenhöft, Elbe#",
		"A06#GT-Jahr    :#      2025#",
	}
	header, err := parseHeader(headerLines)
	if err != nil {
		t.Errorf("parseHeader() error = %v, wantErr %v", err, nil)
	}
	if header.GtName != "Cuxhaven, Steubenhöft, Elbe" {
		t.Errorf("parseHeader() GtName = %v, want %v", header.GtName, "Cuxhaven, Steubenhöft, Elbe")
	}
	if header.GtYear != 2025 {
		t.Errorf("parseHeader() GtYear = %v, want %v", header.GtYear, 2025)
	}
}

//func TestLoad(t *testing.T) {
//	// Assuming Load function takes a string argument and returns an error
//	err := Load("../data/DE__506P2025.txt")
//	if err != nil {
//		t.Errorf("Load() error = %v, wantErr %v", err, nil)
//	}
//}
