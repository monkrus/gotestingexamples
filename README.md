<p align="center">
    <img src="icon.png" width="280" height="170">
</p>

# Go testing examples
Some test writing examples in **Golang**.

☝️ Every folder contains a different testing methods using standard libraries, convey and fuzz testing.

👉 We tried to keep it as simple as possible to make the logic behind the test as clear as possible.

-----

## How to run test:
- make sure you are in root folder
- run the test by `$ go test` ✅

### Troubleshooting:
- in case you have outdated dependencies in `go.mod` file 
  - run `$ go mod tidy`
- if you have conflicts between your local Go 1.xx version and version in `go.mod` file 
  - create new `go.mod` file by `$ go mod init main`

## Check test coverage:
if you want to double-check if your cover 100% or not
- run `$ go test -cover` 😎
 



