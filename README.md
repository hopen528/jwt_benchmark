# Run benchmark test for JWT signing.

$ go test -bench=. -count 5 | tee new.txt
$ go1.19 test -bench=. -count 5 | tee old.txt
$ benchstat old.txt new.txt 

## Misc

If go1.19 is not downloaded, 

$ go install golang.org/dl/go1.19@latest
$ go1.19 download

If benchstat is not installed,

$ go get golang.org/x/perf/cmd/benchstat
$ go install golang.org/x/perf/cmd/benchstat