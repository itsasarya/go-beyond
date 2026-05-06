# Variables, Types, and Control Flow in Go

# Table of Contents

1. Introduction
2. Variables in Go
3. Basic Data Types
4. Zero Values
5. Constants
6. Type Conversion
7. Conditional Statements
8. Switch Statements
9. Loops in Go
10. Break and Continue
11. Input Handling
12. Validation Patterns
13. Building Better Logic
14. Mini Project — Temperature Converter
15. Common Mistakes
16. Key Takeaways

---

# 1. Introduction

This is where Go programming starts becoming practical.

This chapter focuses on:

- Variables
- Types
- Conditions
- Loops
- Decision making
- Input validation
- Basic program structure

The goal is not memorizing syntax.

The real goal is:

> Write logic cleanly without hesitation.

---

# 2. Variables in Go

Go provides multiple ways to declare variables.

## Using `var`

```go
var age int = 25
```

Explicit declaration:

- variable name
- type
- value

---

## Short Declaration (`:=`)

```go
name := "Alex"
```

Go automatically infers the type.

### Important Rule

`:=` only works inside functions.

---

## Best Practice

Inside functions:

```go
count := 10
```

Outside functions:

```go
var version string
```

---

# 3. Basic Data Types

## Integer

```go
var age int = 21
```

Used for whole numbers.

---

## Float

```go
var price float64 = 19.99
```

Used for decimal numbers.

---

## String

```go
var name string = "Go"
```

Used for text.

---

## Boolean

```go
var active bool = true
```

Used for true/false logic.

---

# 4. Zero Values

Variables in Go always have a default value.

| Type    | Zero Value |
| ------- | ---------- |
| int     | 0          |
| float64 | 0          |
| string  | ""         |
| bool    | false      |

Example:

```go
var x int
fmt.Println(x)
```

Output:

```text
0
```

---

# 5. Constants

Constants cannot change after declaration.

```go
const pi = 3.14159
```

Useful for:

- fixed values
- configuration
- formulas

---

# 6. Type Conversion

Go does not allow implicit conversion.

This is invalid:

```go
var x int = 10
var y float64 = x
```

Correct:

```go
var y float64 = float64(x)
```

Go forces explicit conversion to avoid hidden bugs.

---

# 7. Conditional Statements

## Basic If Statement

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

---

## Important Go Rule

Conditions must return a boolean.

Invalid:

```go
if x {
}
```

Correct:

```go
if x != 0 {
}
```

Go does not support truthy/falsy behavior like JavaScript or Python.

---

## If Else

```go
if score >= 90 {
    fmt.Println("A")
} else {
    fmt.Println("Not A")
}
```

---

# 8. Switch Statements

Switch is cleaner than long if-else chains.

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Unknown")
}
```

---

## Important Rule

Go switch statements do NOT fall through automatically.

---

## Fallthrough

```go
switch x {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

Output:

```text
One
Two
```

`fallthrough` forces execution of the next case.

---

# 9. Loops in Go

Go only has one loop keyword:

```go
for
```

---

## Standard Loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

---

## Infinite Loop

```go
for {
    fmt.Println("Running")
}
```

---

## Loop with Condition

```go
for x < 10 {
    x++
}
```

---

# 10. Break and Continue

## Continue

Skips remaining code in the current iteration.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }

    fmt.Println(i)
}
```

Output:

```text
0
1
3
4
```

---

## Break

Stops the loop completely.

```go
for {
    break
}
```

---

# 11. Input Handling

## Reading User Input

```go
var age int
fmt.Scanln(&age)
```

`&age` passes the memory address.

---

## Common Input Problem

If user enters invalid data:

```text
abc
```

Scanning may fail.

Always check errors.

---

# 12. Validation Patterns

## Using a Map as a Set

Instead of:

```go
if unit == "c" || unit == "f" || unit == "k"
```

Use:

```go
validUnits := map[string]bool{
    "c": true,
    "f": true,
    "k": true,
}
```

Validation:

```go
if !validUnits[input] {
    fmt.Println("Invalid input")
}
```

This scales much better.

---

# 13. Building Better Logic

## Early Return Pattern

Bad:

```go
if invalid {
    // logic
} else {
    // real code
}
```

Better:

```go
if invalid {
    return
}

// main logic
```

Cleaner and easier to read.

---

## Handle Edge Cases Early

Examples:

- invalid input
- empty values
- same-unit conversion
- divide by zero

---

# 14. Mini Project — Temperature Converter

## Features Built

- Celsius ↔ Fahrenheit
- Celsius ↔ Kelvin
- Fahrenheit ↔ Kelvin
- Input validation
- Same-unit detection
- Conversion chaining
- Structured switch logic

---

## Example Conversion Function

```go
func c2f(temp float64) float64 {
    return (9.0/5.0)*temp + 32
}
```

---

## Important Lesson

Correct logic matters more than memorized formulas.

Always sanity check:

- 0°C = 32°F
- 0K = -273.15°C

---

# 15. Common Mistakes

## 1. Wrong Formulas

Mixing:

- Celsius
- Fahrenheit
- Kelvin

---

## 2. Truthy/Falsy Thinking

This does NOT work:

```go
if x {
}
```

---

## 3. Infinite Loops Accidentally

Bad:

```go
for x >= 0 {
}
```

Without updating `x`.

---

## 4. Broken Counters

Using `continue` before incrementing counters.

---

## 5. Overusing Global Variables

Prefer local variables whenever possible.

---

# 16. Key Takeaways

By the end of this chapter, you should be able to:

- Declare variables correctly
- Understand Go types
- Use zero values confidently
- Write conditions properly
- Use loops naturally
- Handle input safely
- Validate data cleanly
- Structure logic clearly
- Think step-by-step instead of guessing

---

# Final Thought

Good Go code is:

- explicit
- readable
- predictable
- structured

Avoid clever code.

Write code another developer can understand immediately.
