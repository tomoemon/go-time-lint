# What is timelint?

timelint is a lint tool to avoid problems related to timezone when writing Go programs.

When calling functions that depend on a timezone, such as the Date() and Hour() functions, you must be aware of the timezone of the time.Time variable, or you may encounter a fatal bug. Since these bugs are generally hard to find in reviews and tests, they should be checked with lint tools.

# Examples

The following code will output reports in three lines where timezone is not explicitly specified.

```go
func myTimeFunc() time.Time {
	return time.Now()
}

func f() {
	a := time.Now()
	fmt.Printf("%+v", a.Day()+20) // report "time.Day() called without In(timezone)"
	fmt.Printf("%+v", a.In(time.UTC).Day())
	if h := a.Hour(); h > 10 { // report "time.Hour() called without In(timezone)"
		fmt.Println(h)
	}
	defer func() {
		fmt.Println(myTimeFunc().Date()) // report "time.Date() called without In(timezone)"
	}()
}

```

# How to install

```
go get -u github.com/tomoemon/go-time-lint
```

# How to use

```
timelint ./...
```
