# Go programming for dummies
Made this repo while reading through the Book along with things i learnt from other places like video tutorials etc


## Chapter 1
#### - Wagwan Go!
- To run
```sh
go build
```
- then
```sh
./chapter_1
# on Windows chapter_1.exe
```

To see the values of Go environment variables
```sh
go env
```

To compile your program for another os your need to set `GOOS` and `GOARCH`

OS      | GOOS     | GOARCH
--------|----------|------
macOS   | darwin   | amd64
linux   | linux    | amd64
windows | windows  | amd64

- For linux
```sh
GOOS=linux GOARCH=amd64 go build -0 Chapter_1-linux
```
for Raspberry Pi `GOARCH` will be `arm64`


## Chapter 2
#### - Datatypes
- To run
```sh
go run main.go
```

```go
// Short variable declaration
first_name := "Radiance"
last_name, username, streaks := "Babajide", "@tcitrogg", 20
```

```go
// Multiple variable declaration
var (
    first_name := "Radiance"
    username string = "@tcitrogg"
    streaks int = 20
)
```