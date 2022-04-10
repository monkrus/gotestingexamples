<p align="center">
    <img src="icon.png" width="280" height="170">
</p>

# Go testing examples
Some test writing examples in **Golang**.

☝️ Every folder contains a different testing methods using standard libraries, convey and fuzz testing.

👉 We tried to keep it as simple as possible to make the logic behind the test as clear as possible.

## How to run test:
- make sure you are in selected folder e.g. `cd compare`
- run the test by `$ go test` ✅

### Troubleshooting:
- in case you have outdated dependencies in `go.mod` file
  - run `$ go mod tidy`
- if you have conflicts between your local Go 1.xx version and version in `go.mod` file
  - create new `go.mod` file by `$ go mod init main`
  

## Check test coverage:
> if you want to check that your code cover is 100%

**Print result in console:**
- run `$ go test -cover -v` 😎

**Open result in browser:**
- run test `$ go test -coverprofile=coverage.out`
- open in browser `$ go tool cover -html=coverage.out`