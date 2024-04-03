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
- Datatypes

```sh
# To run
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
- If statement
- Switch-case statement

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
- For loop
- Infinite loop

```go
count++ // count + 1
count-- // count - 1
```

Fibonacci in short lines
```go
fmt.Println("Fibonacci")
max := 100
//    init;     condition;     post
for a, b := 0, 1; b <= max; a, b = b, a+b {
    println(b)
}
```

Iterating over an array with index
```go
var OS [3]string
OS[0] = "iOS"
OS[1] = "Android"
OS[2] = "SailOS"

for index, value := range OS {
    println(index, value)
}
```


## Chapter 5
- Functions
- Anonymous Functions
- Closure

```go
func add_num(num1, num2 int) (sum int) {
    result = num1 + num2
    return
    // We can still use `return result`
}
```

Closure
```go
func fib() func() int {
	f1 := 0
	f2 := 1
	// this closure
	return func() int {
		f1, f2 = f2, f1+f2
		return f2
	}
}

// to run
get := fib()
fmt.Println(get())
```


## Chapter 6
- Array
- Slice

Array is like tuple in Python, Slice is like list