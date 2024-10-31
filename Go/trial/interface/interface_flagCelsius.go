/*
package flag

	type Value interface{
		String() string
		Set(string) error
	}

Value is the interface to the value stored in a flag
*/
package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

/*
	package main

	import "tempconv"

	var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
	func main(){
		flag.Parse()
		fmt.Println(*temp)
	}
*/
