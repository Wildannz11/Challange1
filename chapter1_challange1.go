package main

import "fmt"

func main () {
	var i int = 21
	no2 := 10
	t := true
	no4 := t
	
	no5 := 1071
	no6 := 21
	no7 := 21
	no8 := 15
	
	no9 := 15
	no10 := 1071
	no11 := 123.456
	no12 := 123.456



	fmt.Printf(" no1 => %v", i)
	fmt.Printf("\n no2 => %T", no2)
	fmt.Printf("\n no3 => %%")
	fmt.Printf("\n no4 => %t", no4)

	// fmt.Printf("\n no5 => %b", no5)
	fmt.Printf("\n no5 => %c", no5)
	fmt.Printf("\n no6 => %d", no6)
	fmt.Printf("\n no7 => %o", no7)
	fmt.Printf("\n no8 => %x", no8)

	fmt.Printf("\n no9 => %X", no9)
	fmt.Printf("\n no10 => %U", no10)
	fmt.Printf("\n no11 => %f", no11)
	fmt.Printf("\n no12 => %E", no12)

}


