<p align="center">
    <img src="icon.png" width="280" height="170">
</p>

# Go testing examples
Some test writing examples in **Golang**.

☝️ Every folder contains a different testing methods using standard libraries, convey and fuzz testing.

👉 We tried to keep it as simple as possible to make the logic behind the test as clear as possible.

-----

## How to run test:
- checkout to selected folder e.g. `$ cd compare`
- 💡 to create `go.mod` follow 
    - `$ go mod init compare/m` to initialize a v0 or v1 module
    - `$ go mod tidy` to create go.mod file
- run the test by `$ go test` ✅

-----

## Check test coverage:
> if you want to check that your code cover is 100%

**Print result in console:**
- run `$ go test -cover -v` 😎

**Open result in browser:**
- run test `$ go test -coverprofile=coverage.out`
- open in browser `$ go tool cover -html=coverage.out`
 




