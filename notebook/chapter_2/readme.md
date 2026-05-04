# Chapter 2 — Variables, Types, and Control Flow

## Core Goal

Write basic logic in Go without hesitation: declare variables, use conditions, loops, and structure decisions clearly.

---

## 1. Variables

### Declaration styles

```go
var age int = 25
name := "Alex" // shorthand (inside functions only)
```

### Rules

- `:=` only inside functions
- `var` works everywhere
- Prefer `:=` inside functions for brevity

---

## 2. Basic Types

- `int`
- `float64`
- `string`
- `bool`

### Zero Values (default)

```go
var i int       // 0
var s string    // ""
var b bool      // false
```

---

## 3. Type Conversion

Go does **not** do implicit conversion.

```go
var i int = 10
var f float64 = float64(i)
```

---

## 4. Conditions (`if`)

- Condition must be **boolean**

```go
if x != 0 {
    // valid
}
```

❌ Not allowed:

```go
if x { }
```

---

## 5. Switch

```go
switch day {
case 1:
    fmt.Println("Mon")
case 2:
    fmt.Println("Tue")
default:
    fmt.Println("Unknown")
}
```

### Key Rules

- No automatic fallthrough
- Use `fallthrough` explicitly if needed

---

## 6. Loops (`for`)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### Keywords

- `continue` → skip rest of current iteration
- `break` → exit loop

---

## 7. Logic Thinking

- Always be explicit
- Break problems into steps
- Avoid guessing

---

## 8. Input Handling

```go
var x float64
fmt.Scanln(&x)
```

---

## 9. Validation Pattern

Use a map as a set:

```go
validUnits := map[string]bool{
    "c": true,
    "f": true,
    "k": true,
}

if !validUnits[input] {
    fmt.Println("Invalid input")
}
```

---

## 10. Control Flow Design Principles

- Validate early → return early
- Handle edge cases first (invalid input, same values)
- Keep logic readable (prefer `switch` over long `if-else` chains)

---

## 11. Common Mistakes

- Using wrong formulas (always sanity check)
- Mixing units or logic steps
- Overusing global variables
- Writing messy conditional chains

---

## 12. Mini Project

- **Temperature Converter**

  > Features implemented:
  >
  > - Input temperature
  > - Input source & target units
  > - Validation using map
  > - Same-unit check
  > - Conversion logic using functions

- **Number Guess**
  > Features implemented:
  >
  > - Let user guess number
  > - Guide them with too high or to Low
  > - Validate input and range
  > - Loop till guessed right

---

## Final Takeaway

- Be explicit
- Think in steps
- Structure logic cleanly
- Validate before processing

You should now be able to write small programs with clear control flow and correct
