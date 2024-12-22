package bshe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	GtName string
	GtYear int
}

type DataPoint struct {
	LineCode    string
	GaugeNumber string
	//MoonPhase     int - the value is empty in a lot of data points
	TidePhase rune
	DayOfWeek string
	Timestamp time.Time
	Height    float64
	// QualityFlag   int - the value is empty in a lot of data points
	NumberOfDay   int
	UTCTimeDiff   time.Time
	TransitNumber int
	Transit       int
	JulianDate    float64
}

type TideReport struct {
	Header Header
	Data   []DataPoint
}

func parseDataLine(l string) (DataPoint, error) {
	var err error
	dp := DataPoint{}

	parts := strings.Split(l, "#")
	if len(parts) < 2 {
		return DataPoint{}, fmt.Errorf("invalid data line: %v", l)
	}
	dp.LineCode = parts[0]
	dp.GaugeNumber = parts[1]
	//dp.MoonPhase, _ = strconv.Atoi(parts[2])
	dp.TidePhase = rune(parts[3][0])
	dp.DayOfWeek = parts[4]

	dp.UTCTimeDiff, err = time.Parse("+15:04", strings.ReplaceAll(parts[10], " ", "0"))
	if err != nil {
		return DataPoint{}, fmt.Errorf("invalid UTC time difference: %w", err)
	}

	dateTimeStr := strings.ReplaceAll(parts[5], " ", "0") + "-" + strings.ReplaceAll(parts[6], " ", "0")
	dp.Timestamp, err = time.Parse("02.01.2006-15:04", dateTimeStr)
	if err == nil {
		dp.Timestamp = dp.Timestamp.Add(-dp.UTCTimeDiff.Sub(time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)))
	} else {
		return DataPoint{}, fmt.Errorf("invalid date time: %w", err)
	}

	dp.Height, err = strconv.ParseFloat(strings.ReplaceAll(parts[7], " ", ""), 64)
	if err != nil {
		return DataPoint{}, fmt.Errorf("invalid height: %w", err)
	}

	//dp.QualityFlag, _ = strconv.Atoi(parts[8])
	dp.NumberOfDay, err = strconv.Atoi(strings.TrimSpace(parts[9]))
	if err != nil {
		return DataPoint{}, fmt.Errorf("invalid number of day: %w", err)
	}

	dp.TransitNumber, err = strconv.Atoi(strings.TrimSpace(parts[11]))
	if err != nil {
		return DataPoint{}, fmt.Errorf("invalid transit number: %w", err)
	}

	dp.Transit, err = strconv.Atoi(strings.TrimSpace(parts[12]))
	if err != nil {
		return DataPoint{}, fmt.Errorf("invalid transit: %w", err)
	}

	dp.JulianDate, _ = strconv.ParseFloat(strings.TrimSpace(parts[13]), 64)

	return dp, nil
}

func parseData(dataLines []string) ([]DataPoint, error) {

	var dataPoints []DataPoint

	for _, line := range dataLines {
		dp, err := parseDataLine(line)
		if err != nil {
			return nil, err
		}
		dataPoints = append(dataPoints, dp)
	}

	return dataPoints, nil
}

func parseHeader(headerLines []string) (Header, error) {

	header := Header{}

	for _, line := range headerLines {
		switch {
		case strings.HasPrefix(line, "A04"):
			re := regexp.MustCompile(`A04#GT-Name\s*:\s*#(.*)#`)
			matches := re.FindStringSubmatch(line)
			if matches != nil {
				header.GtName = matches[1]
			} else {
				return Header{}, fmt.Errorf("Error parsing GT-Name: %v", line)
			}
		case strings.HasPrefix(line, "A06"):
			re := regexp.MustCompile(`A06#GT-Jahr\s*:\s*#\s*(\d*)#`)
			matches := re.FindStringSubmatch(line)
			if matches != nil {
				year, err := strconv.Atoi(matches[1])
				if err != nil {
					return Header{}, fmt.Errorf("Error converting GT-Jahr: %w", err)
				}
				header.GtYear = year
			} else {
				return Header{}, fmt.Errorf("Error parsing GT-Jahr: %v", line)
			}
		}
	}

	return header, nil
}

func Load(fileName string) (TideReport, error) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return TideReport{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var headerLines []string
	var dataLines []string

	isHeader := true

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "LLL#") {
			isHeader = false
			continue
		} else if strings.HasPrefix(line, "EEE#") {
			break
		}

		if isHeader {
			headerLines = append(headerLines, line)
		} else {
			dataLines = append(dataLines, line)
		}
	}

	header, err := parseHeader(headerLines)
	if err != nil {
		fmt.Printf("Error parsing header: %v", err)
	}

	data, err := parseData(dataLines)
	if err != nil {
		fmt.Printf("Error parsing data: %v", err)
	}

	tideReport := TideReport{
		Header: header,
		Data:   data,
	}

	return tideReport, nil
}
