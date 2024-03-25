### Run test

```
$ go test -test.v > out.txt
```

Output:
```
 === RUN   TestFtToMet
    lenconv_test.go:13: Convert random: 0 Ft. to Meter == 0 Meter
    lenconv_test.go:16: TestFtToMet Passed!!!
--- PASS: TestFtToMet (0.00s)
=== RUN   TestMetToFt
    lenconv_test.go:27: Convert random: 0 Met to Ft. == 0 Ft.
    lenconv_test.go:30: TestMetToFt Passed!!!
--- PASS: TestMetToFt (0.00s)
PASS
ok  	github.com/unixlinuxgeek/lenconv	0.002s
```
