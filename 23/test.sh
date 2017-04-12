go build
go install
go test -cpu=4 -bench=. gopl.io/ch2/popcount
