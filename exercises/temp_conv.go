package main

import (
	"fmt"
	"strings"
)

func ConvertTemprature() {
	var temperature float64
	var from_unit string
	var to_unit string
	fmt.Println("Enter Temperature: ")
	fmt.Scanln(&temperature)
	fmt.Println("Enter unit(c, k, f): ")
	fmt.Scanln(&from_unit)
	fmt.Println("Enter converstion unit(c, k, f): ")
	fmt.Scanln((&to_unit))

	lowerFromUnit := strings.ToLower(from_unit)
	lowerToUnit := strings.ToLower(to_unit)

	validUnits := map[string]bool{
		"c": true,
		"f": true,
		"k": true,
	}
	// validate input unit
	if !validUnits[lowerFromUnit] || !validUnits[lowerToUnit] {
		fmt.Println("Invalid unit")
		return
	}
	// check for same unit for conversion
	if lowerFromUnit == lowerToUnit {
		fmt.Println(temperature)
		return
	}

	switch lowerFromUnit {
	case "c":
		switch lowerToUnit {
		case "f":
			result := c2f(temperature)
			fmt.Println(result)
		case "k":
			result := c2k(temperature)
			fmt.Println(result)
		}
	case "f":
		switch lowerToUnit {
		case "k":
			c := f2c(temperature)
			result := c2k(c)
			fmt.Println(result)
		case "c":
			result := f2c(temperature)
			fmt.Println(result)
		}
	case "k":
		switch lowerToUnit {
		case "c":
			result := k2c(temperature)
			fmt.Println(result)
		case "f":
			c := k2c(temperature)
			result := c2f(c)
			fmt.Println(result)
		}
	}
}

func c2f(temp float64) float64 {
	// converts Celsius to Fahrenheit
	return ((9.0/5.0)*temp + 32)
}

func f2c(temp float64) float64 {
	// converts Fahrenheit to Celsius
	return ((5.0 / 9.0) * (temp - 32))
}

func c2k(temp float64) float64 {
	// converts celsius to kelvin
	return temp + 273.15
}

func k2c(temp float64) float64 {
	// converts kelvin to celsius
	return temp - 273.15
}
