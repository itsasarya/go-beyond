# Functions and Scope
  
# Introduction

This is where Go starts feeling like real backend engineering.

Before this point, programming mostly looks like:

- variables
- loops
- conditions
- basic calculations

But backend systems are not built from isolated lines of code.
They are built from:

- reusable functions
- predictable error handling
- controlled execution flow
- safe cleanup
- properly scoped variables

This chapter teaches the foundation of all of that.

If you look at production Go services, most files are made of:

- functions calling functions
- values being returned
- errors being checked
- resources being cleaned up with `defer`
- variables carefully managed to avoid hidden bugs

That is why this chapter matters.

---

# Chapter 1 — Functions

A function is a reusable block of logic.

Instead of rewriting the same code repeatedly, you place it inside a function and call it whenever needed.

Functions are one of the main tools used to organize backend systems.

Without functions:

- code becomes repetitive
- logic becomes hard to test
- bugs become difficult to isolate
- programs become unreadable

---

# Basic Function Syntax

```go
func add(a int, b int) int {
    return a + b
}
```

This small function already contains several important ideas.

| Part             | Meaning             |
| ---------------- | ------------------- |
| `func`           | Declares a function |
| `add`            | Function name       |
| `(a int, b int)` | Parameters          |
| `int`            | Return type         |
| `return`         | Sends a value back  |

---

# Thinking About Functions Properly

A beginner usually thinks:

> “Functions are just reusable code.”

That is true, but incomplete.

A better way to think about functions is:

> Functions create boundaries.

Inside the boundary:

- work happens
- logic runs
- data changes

Outside the boundary:

- callers only care about inputs and outputs

This separation is what makes large backend systems manageable.

---

# Parameters

Parameters are inputs.

Example:

```go
func greet(name string) string {
    return "Hello " + name
}
```

Usage:

```go
message := greet("Aman")
fmt.Println(message)
```

Output:

```text
Hello Aman
```

The function does not care where the value came from.
It only cares about the input it received.

That separation is extremely important in backend systems.

---

# Return Values

Functions can return values.

Example:

```go
func square(n int) int {
    return n * n
}
```

Usage:

```go
result := square(5)
fmt.Println(result)
```

Output:

```text
25
```

---

# Chapter 2 — Multiple Return Values

Go allows functions to return multiple values.

This is one of the language’s most important design choices.

Example:

```go
func divide(a int, b int) (int, int) {
    return a / b, a % b
}
```

Usage:

```go
q, r := divide(10, 3)
```

Result:

```text
q = 3
r = 1
```

---

# Why Multiple Returns Matter

Most programming languages rely heavily on exceptions.

Go takes a different approach.

Instead of throwing exceptions for normal failures, Go returns:

```go
value, err
```

This design makes failures explicit.

You cannot ignore them accidentally.

---

# The Core Backend Pattern

This pattern exists everywhere in Go:

```go
value, err := doSomething()
if err != nil {
    return err
}
```

You will see this thousands of times.

Understanding this deeply is mandatory for backend development.

---

# Chapter 3 — Error Handling

Go treats errors as normal values.

Errors are not magical.
They are simply returned like other data.

Example:

```go
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }

    return a / b, nil
}
```

---

# Breaking Down the Error Return

When division fails:

```go
return 0, fmt.Errorf("divide by zero")
```

Two things are returned:

| Value   | Meaning            |
| ------- | ------------------ |
| `0`     | Zero value for int |
| `error` | Explains failure   |

---

# Why Return Zero?

The function signature requires:

```go
(int, error)
```

That means both values must always be returned.

Even during failure.

Since there is no valid result, the zero value is returned.

---

# What Does `nil` Mean?

Successful execution:

```go
return a / b, nil
```

`nil` means:

> “There is no error.”

---

# Correct Error Flow

Proper execution flow:

```go
value, err := Divide(10, 2)
if err != nil {
    return err
}

fmt.Println(value)
```

Notice the order carefully:

1. Call function
2. Receive values
3. Check error FIRST
4. Use value ONLY if error is nil

This order matters.

---

# A Dangerous Beginner Mistake

Bad:

```go
value, err := doThing()
fmt.Println(value)
```

Problem:

You used the value before checking whether it is valid.

Correct:

```go
value, err := doThing()
if err != nil {
    return err
}

fmt.Println(value)
```

---

# Chapter 4 — Scope

Scope controls where variables exist.

Example:

```go
func main() {
    x := 10

    if true {
        x := 20
        fmt.Println(x)
    }

    fmt.Println(x)
}
```

Output:

```text
20
10
```

---

# Understanding What Happened

Inside the `if` block:

```go
x := 20
```

created a NEW variable.

It did NOT update the outer variable.

This is called:

# Shadowing

---

# Shadowing

Shadowing happens when a new variable hides another variable with the same name.

Example:

```go
err := doThing()

if err != nil {
    err := fmt.Errorf("wrapped error")
}
```

This looks harmless.

But it creates TWO different variables.

That causes logic bugs.

---

# `:=` vs `=`

Understanding this distinction is critical.

## `:=`

Creates a new variable.

## `=`

Updates an existing variable.

---

# Silent Backend Bugs

This type of mistake is dangerous because:

- the program compiles
- the program runs
- the logic becomes wrong silently

Those are the worst bugs.

---

# Chapter 5 — Defer

`defer` delays execution until the surrounding function exits.

Example:

```go
func main() {
    defer fmt.Println("done")

    fmt.Println("working")
}
```

Output:

```text
working
done
```

---

# Understanding Defer Properly

This line:

```go
defer fmt.Println("done")
```

DOES NOT execute immediately.

It registers the function call.

Execution happens later.

Specifically:

> Right before the surrounding function exits.

---

# Why Defer Exists

Backend systems constantly open resources:

- files
- database connections
- HTTP response bodies
- mutex locks

Those resources must be cleaned up.

Without `defer`, cleanup is easy to forget.

---

# Real Example

```go
file, err := os.Open("data.txt")
if err != nil {
    return err
}

defer file.Close()
```

This guarantees cleanup.

Even if:

- errors happen
- returns happen early
- panic occurs

---

# The Mental Model

Think of `defer` as:

> “Before leaving this function, do this cleanup.”

---

# Multiple Defer Calls

Defer follows:

# LIFO (Last In, First Out)

Example:

```go
func test() {
    defer fmt.Println("A")
    defer fmt.Println("B")

    fmt.Println("C")
}
```

Output:

```text
C
B
A
```

Why?

The last deferred function runs first.

---

# Chapter 6 — Panic

`panic` immediately stops normal execution.

Example:

```go
panic("something broke")
```

Unlike normal errors:

- execution stops
- stack unwinds
- deferred functions run
- program may crash

---

# When Panic Should Be Used

Rarely.

Panic is NOT for normal backend failures.

Do NOT panic for:

- invalid user input
- missing records
- authentication failures
- normal database errors

Use returned errors instead.

---

# Appropriate Panic Cases

Panic is acceptable for:

- impossible states
- corrupted internal state
- fatal startup failures
- programmer mistakes

---

# Chapter 7 — Recover

`recover()` catches a panic.

It only works inside deferred functions.

Example:

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    panic("boom")
}
```

Output:

```text
Recovered: boom
```

---

# Important Recover Rule

This works:

```go
defer func() {
    recover()
}()
```

This does NOT:

```go
recover()
```

outside deferred execution.

---

# Why Recover Matters in Backends

Production HTTP servers often use recovery middleware.

If one request panics:

- server should not fully crash
- panic should be logged
- request should fail safely

Recover helps achieve that.

---

# Chapter 8 — Utility Functions Built During Practice

## Divide

```go
func Divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("divide by zero")
    }

    return a / b, nil
}
```

---

## IsEven

```go
func IsEven(n int) bool {
    return n%2 == 0
}
```

---

## IsOdd

```go
func IsOdd(n int) bool {
    return n%2 != 0
}
```

---

## ReverseString

```go
func ReverseString(s string) string {
    var result strings.Builder

    for i := len(s) - 1; i >= 0; i-- {
        result.WriteString(string(s[i]))
    }

    return result.String()
}
```

---

# Why `strings.Builder` Is Better

This:

```go
result += something
```

creates repeated string allocations.

Strings are immutable.

That means every concatenation creates another string.

`strings.Builder` avoids repeated allocations.

Benefits:

- faster execution
- lower memory usage
- better scalability

Important in backend systems.

---

# Unicode Warning

This:

```go
s[i]
```

works with bytes.

Not full Unicode characters.

So naive string reversal can fail for:

- emojis
- Hindi
- Japanese
- Unicode text

Proper Unicode handling uses runes.

---

# Chapter 9 — Common Beginner Mistakes

## Forgetting Error Checks

Bad:

```go
value, err := doThing()
fmt.Println(value)
```

Good:

```go
value, err := doThing()
if err != nil {
    return err
}
```

---

## Using Panic for Normal Failures

Bad:

```go
panic("user not found")
```

Good:

```go
return fmt.Errorf("user not found")
```

---

## Shadowing Accidentally

Bad:

```go
err := doThing()

if err != nil {
    err := fmt.Errorf("wrapped")
}
```

---

## Forgetting Cleanup

Bad:

```go
file, _ := os.Open("x.txt")
```

Good:

```go
file, err := os.Open("x.txt")
if err != nil {
    return err
}

defer file.Close()
```

---

# Chapter 10 — Mental Models

## Functions

Reusable logic boundaries.

---

## Errors

Expected failures.

---

## Panic

Program state is broken.

---

## Defer

Guaranteed cleanup.

---

## `:=`

May create new variables.

---

## `=`

Updates existing variables.

---

# What You Should Know After This Chapter

You should now be able to:

- write reusable functions
- return multiple values
- follow Go error handling flow
- understand scope properly
- avoid shadowing bugs
- use defer safely
- distinguish panic vs errors
- build reusable utility logic

---

# Final Thoughts

This is one of the most important foundations in Go.

Most backend services are primarily made from:

- functions
- returned errors
- deferred cleanup
- scoped variables

If these concepts become natural, later backend topics become dramatically easier.

Do not rush through this.

Master it properly.
