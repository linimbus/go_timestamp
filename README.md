# go_timestamp

## Functional Description
This is a tool that uses pipes to output timestamps. It is written in golang. If your script or software does not have a log format that outputs timestamps, this may be helpful.

## How to use
If you are familiar with Linux pipe notation, it is very simple. Here is an example test.sh.

Before:
```
$ ./test.sh 
run test sleep 6
run test sleep 5
run test sleep 6
run test sleep 8
run test sleep 6
run test sleep 7
run test sleep 6
run test sleep 3
run test sleep 9
run test sleep 1
```

After:
```
./test.sh | ./go_timestamp 
2024-10-17 07:58:03.808 run test sleep 8
2024-10-17 07:58:11.810 run test sleep 6
2024-10-17 07:58:17.812 run test sleep 1
2024-10-17 07:58:18.814 run test sleep 2
2024-10-17 07:58:20.815 run test sleep 4
2024-10-17 07:58:24.817 run test sleep 7
2024-10-17 07:58:31.819 run test sleep 0
2024-10-17 07:58:31.820 run test sleep 4
2024-10-17 07:58:35.822 run test sleep 3
2024-10-17 07:58:38.823 run test sleep 9
```


## Option
```
Usage of ./go_timestamp:
  -buffer int
        Maximum bytes to read from the pipe per loop (default 8192)
  -help
        Usage help
  -layout string
        Timestamp output format layout, see https://golang.org/pkg/time/#Time.Format, some examples below:"
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
                TimeOnly    : "15:04:05"
```

Example:
```
./test.sh | ./go_timestamp -layout "Monday, 02-Jan-06 15:04:05 MST"
Thursday, 17-Oct-24 08:03:19 UTC run test sleep 1
Thursday, 17-Oct-24 08:03:20 UTC run test sleep 1
Thursday, 17-Oct-24 08:03:21 UTC run test sleep 7
```
