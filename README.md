<p align="center">
    <img src="icon.png" width="280" height="170">
</p>

# Go testing examples
Some test writing examples in **Golang**.Fuzz testing will only work starting with Go 1.18.

☝️ Every folder contains a different testing method using standard libraries, convey and fuzz testing.

👉 We tried to keep it as simple as possible to make the logic behind the test as clear as possible.

## How to run test:
- make sure you are in selected folder e.g. `cd compare`
- run the test by `$ go test` ✅
## For fuzz test :
- run `go test` or `go test -v`
- run `go test -run=TestName`(for specific test)


## Troubleshooting:
- if you have an outdated dependencies in `go.mod` file
  - run `$ go mod tidy`
- if you have conflicts between your local Go 1.xx version and version in `go.mod` file
  - create new `go.mod` file by `$ go mod init nameoftheproject`
 
 
## Check test coverage:
- to check that your code is 100% covered

**Print result in console:**
- run `$ go test -cover -v` 😎

**Open result in browser:**
- run test `$ go test -coverprofile=coverage.out`
- open in browser `$ go tool cover -html=coverage.out`


List of test examples:
-
> here is the list of folders with function implementation and corresponding test 
---
- [x] [compare array string](compare-array-string)  
- [x] [convey](convey)  
- [x] [fuzz](fuzz)  
- [x] [hello-world](hello-world)  
- [x] [math](math)  
- [x] [simple](simple)
- [ ] ...

# Tips & Tricks
- (using 'Goland' Jetbrains) generate test template in two clicks
  [video](https://www.jetbrains.com/go/guide/tips/generate-a-test-for-an-element/) 
