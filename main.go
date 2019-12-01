package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

// Parse cli flags to override configuration defaults
func parseFlags(conf *Config) {
	var inFile string
	flag.StringVar(&inFile, "in-file", conf.InFile, "input file name")

	var fromTime string
	flag.StringVar(&fromTime, "from", conf.FromTime, "from 2006-02-01 15:04:05")

	var toTime string
	flag.StringVar(&toTime, "to", conf.ToTime, "from 2006-02-01 15:04:05")

	var bufSize int
	flag.IntVar(&bufSize, "buf", conf.BufferSize, "buffer size in KB")

	var diffFormat bool
	flag.BoolVar(&diffFormat, "diff", conf.DiffFormat, "enable diff format - remove timestamps and request IDs")

	// parse flags
	flag.Parse()

	conf.InFile = inFile
	conf.FromTime = fromTime
	conf.ToTime = toTime
	conf.BufferSize = bufSize
	conf.DiffFormat = diffFormat
}

func urldecode(myURL string) string {
	decodedURL, err := url.QueryUnescape(myURL)
	if err != nil {
		return myURL
	}

	return decodedURL
}

func main() {
	// set defaults for Config
	conf := &Config{
		"2006-02-01 15:04:05",
		"",
		"",
		"",
		64,
		false,
	}

	parseFlags(conf)

	// check if file was given
	if strings.TrimSpace(conf.InFile) == "" {
		log.Print("no file given")
		os.Exit(0)
	}

	// open file handler for reading
	file, err := os.Open(conf.InFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// text scanner
	scanner := bufio.NewScanner(file)

	// read 64KB of the message by default but
	// user can increas if it is not sufficient
	scanner.Buffer([]byte{}, conf.BufferSize*1024)

	// iterate over log rows
	var logEntry Entry
	for scanner.Scan() {
		json.Unmarshal([]byte(scanner.Text()), &logEntry)

		// skip forward messages
		if logEntry.Type == "forward" {
			continue
		}

		// skip log entry before start date
		if strings.TrimSpace(conf.FromTime) != "" && conf.GetTime(logEntry.Time, false).Before(conf.GetTime(conf.FromTime, true)) {
			continue
		}

		// skip log entry after end date
		if strings.TrimSpace(conf.ToTime) != "" && conf.GetTime(logEntry.Time, false).After(conf.GetTime(conf.ToTime, false)) {
			continue
		}

		if conf.DiffFormat {
			fmt.Printf("----------- %s - (%s/%s) -------------------\r\n", "*TIMESTAMP*", logEntry.Type, "*REQUESTID*")
		} else {
			fmt.Printf("----------- %s - (%s/%s) -------------------\r\n", logEntry.Time, logEntry.Type, logEntry.RequestID)
		}

		if logEntry.Type == "request" {
			fmt.Printf("%s", urldecode(strings.TrimSpace(logEntry.Message)))
		} else {
			fmt.Printf("%s", strings.TrimSpace(logEntry.Message))
		}

		fmt.Print("\r\n\r\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
