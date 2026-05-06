package main

import "strings"

import "fmt"

func add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}
	return a / b, nil
}

func Subtract(a, b int) int {
	return a + b
}

func IsEven(n int) bool {
	return n%2 == 0
}

func ReverseString(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteString(string(s[i]))
	}
	return result.String()
}
