package main

import (
	"fmt"
)

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(s string) error {
		fmt.Println("do something before calling the next")
		r := s + " should be green"
		return next(r)
	}
}
func home(s string) error {
	fmt.Println("home", s)
	return nil
}
func main() {
	warped := Use(home)

	w := warped("world")
	fmt.Println("end result", w)

}
