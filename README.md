# GoTraining
The repository to submit the exercises of the Go training in 2016

### License
<a rel="license" href="http://creativecommons.org/licenses/by-nc-sa/4.0/">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License</a>.

### Reference
#### Install
To install go env by Homebrew:
```sh
$ brew install go
```
To set GOPATH:
```sh
$ export GOPATH=$HOME/Go/GoTraining
```

#### Usage
To run a test:
```sh
$ go test ch1/ex1/main.go
```
To run a benchmark test:
```sh
$ go test ch1/ex1/main.go -bench=.
```

#### Go tool
To get a go tool:
```sh
$ go get URLofTool
```
Recommended tools:
- [__goimports__](https://godoc.org/golang.org/x/tools/cmd/goimports) imports automatically package.
