package main

import (
	"access-modifier/library"
	"fmt"
)

func main() {
	fmt.Println("HourInADay:", library.HourInADay)
	fmt.Println("Name:", library.Name)
	fmt.Println(library.DaysToMinutes(3))
}
