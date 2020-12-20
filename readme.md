## Quick start
Download `go-fuzz`
```
$ go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
```

Generate the testing zip for the package strfuncs
```
$ cd strfuncs
$ go-fuzz-build
```

Run the fuzz test against the testing zip with no corpus
```
$ go-fuzz -bin=./strfuncs-fuzz.zip
```