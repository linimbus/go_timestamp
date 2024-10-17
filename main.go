package main

import (
	"flag"
	"io"
	"os"
	"time"
)

var (
	Option_help        bool
	Option_layout      string
	Option_buffer_size int
)

func init() {
	flag.StringVar(&Option_layout, "layout", "", `Timestamp output format layout, see https://golang.org/pkg/time/#Time.Format, some examples below:"
	Default     : "2006-01-02 15:04:05.000"
	Layout      : "01/02 03:04:05PM '06 -0700"
	ANSIC       : "Mon Jan _2 15:04:05 2006"
	UnixDate    : "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    : "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      : "02 Jan 06 15:04 MST"
	RFC822Z     : "02 Jan 06 15:04 -0700"
	RFC850      : "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     : "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    : "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339     : "2006-01-02T15:04:05Z07:00"
	RFC3339Nano : "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     : "3:04PM"
	Stamp       : "Jan _2 15:04:05"
	StampMilli  : "Jan _2 15:04:05.000"
	StampMicro  : "Jan _2 15:04:05.000000"
	StampNano   : "Jan _2 15:04:05.000000000"
	DateTime    : "2006-01-02 15:04:05"
	DateOnly    : "2006-01-02"
	TimeOnly    : "15:04:05"`)
	flag.IntVar(&Option_buffer_size, "buffer", 8192, "Maximum bytes to read from the pipe per loop")
	flag.BoolVar(&Option_help, "help", false, "Usage help")
}

func main() {
	flag.Parse()
	if Option_help {
		flag.Usage()
		return
	}

	if Option_layout == "" {
		Option_layout = "2006-01-02 15:04:05.000"
	}

	if Option_buffer_size < 1024 {
		Option_buffer_size = 1024
	}

	var latestBuffer []byte
	for {
		tempBuffer := make([]byte, Option_buffer_size)
		cnt, err := os.Stdin.Read(tempBuffer[:])
		if err != nil && err == io.EOF {
			return
		}

		latestBuffer = append(latestBuffer, tempBuffer[:cnt]...)
		beginIndex := 0
		for i, char := range latestBuffer[:] {
			if char == '\n' || char == '\r' {
				os.Stdout.WriteString(time.Now().Format(Option_layout) + " ")
				os.Stdout.Write(latestBuffer[beginIndex : i+1])
				beginIndex = i + 1
			}
		}
		if beginIndex != 0 {
			latestBuffer = latestBuffer[beginIndex:]
		}
	}
}
