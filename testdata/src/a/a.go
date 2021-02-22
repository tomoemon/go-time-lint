package a

import (
	"fmt"
	"time"
)

func myTimeFunc() time.Time {
	return time.Now()
}

func f() {
	a := time.Now()
	fmt.Printf("%+v", a.Day()+20) // want "time.Day\\(\\) called without In\\(timezone\\)"
	fmt.Printf("%+v", a.In(time.UTC).Day())
	if h := a.Hour(); h > 10 { // want "time.Hour\\(\\) called without In\\(timezone\\)"
		fmt.Println(h)
	}
	defer func() {
		fmt.Println(myTimeFunc().Date()) // want "time.Date\\(\\) called without In\\(timezone\\)"
	}()
}
