### Run test

```
$ go test -test.v > out.txt
```

Example output:
```
=== RUN   TestFtToMet
    lenconv_test.go:13: Convert random: 6 Ft. to Meter == 1.828710758914965 Meter
    lenconv_test.go:16: TestFtToMet Passed!!!
--- PASS: TestFtToMet (0.00s)
=== RUN   TestMetToFt
    lenconv_test.go:27: Convert random: 8 Met to Ft. == 26.248 Ft.
    lenconv_test.go:30: TestMetToFt Passed!!!
--- PASS: TestMetToFt (0.00s)
PASS
ok  	github.com/unixlinuxgeek/lenconv	0.001s
```
