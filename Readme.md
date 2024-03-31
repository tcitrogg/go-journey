# Go journey
Made this repo while reading through the Book **Go programming for dummies** along with things i learnt from other places like video tutorials etc


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
or run like we did in ## Chapter 1

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

```go
// Type convertion
// - Bool
b, err := strconv.ParseBool("t")

// - Float                   value,   bit
f, err := strconv.ParseFloat("3.1415", 64)

// - Integer               value, base, bit
i, err := strconv.ParseInt("-18", 10, 64)

// - Unsigned Integer
f, err := strconv.ParseUint("20", 10, 64)

f, err := strconv.ParseUint("-53", 10, 64)   //  Will error, cause it cannot parse a signed int as an unsigned int
```


## Chapter 3
Comparison has to be between between same types
```go
// For Passing or failing
grade := "C"
switch grade {
// case "A":
// 	fallthrough
// case "B":
// 	fallthrough
// case "C":
// 	fallthrough
// case "D":
// 	fmt.Println("Passed")
// Or we can just do
case "A", "B", "C", "D":
    fmt.Println("Passed")
case "F":
    fmt.Println("Failed")
default:
    fmt.Println("Absent")
}
```


## Chapter 4