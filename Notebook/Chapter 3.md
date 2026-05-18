# Chapter 3 — Collections in Go

## Introduction

Collections are one of the most important parts of backend development.

Most backend systems spend their time:

* storing data
* moving data
* transforming data
* grouping data
* searching data

In Go, the main collection types are:

* Arrays
* Slices
* Maps

Understanding how these work internally is critical because Go exposes memory behavior more directly than many higher-level languages.

This chapter covers:

* arrays
* slices
* slice internals
* append behavior
* copying slices
* maps
* range loops
* common pitfalls

---

# 1. Arrays

## What Is an Array?

An array is a fixed-size collection of elements of the same type.

Example:

```go
var nums [3]int
```

This creates:

* an array of integers
* length = 3
* default values = 0

---

## Array Initialization

```go
package main

import "fmt"

func main() {
	nums := [3]int{10, 20, 30}

	fmt.Println(nums)
}
```

Output:

```text
[10 20 30]
```

---

## Accessing Elements

```go
fmt.Println(nums[0])
```

Output:

```text
10
```

---

## Out of Bounds Access

This causes a runtime panic:

```go
nums[3]
```

Error:

```text
panic: runtime error: index out of range
```

---

## Important Notes About Arrays

Arrays:

* have fixed size
* are rarely used directly in backend development
* copy all elements when assigned

Example:

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

Arrays are copied completely.

---

# 2. Slices

## What Is a Slice?

A slice is a dynamic view over an array.

Slices are used constantly in Go backend development.

Example:

```go
nums := []int{10, 20, 30}
```

Unlike arrays:

* slices do not have fixed size
* slices can grow
* slices are lightweight structures

---

## Slice Structure

Internally, a slice contains:

```text
pointer
length
capacity
```

Visualization:

```text
Underlying array:
[10 20 30 0 0]

Slice:
ptr -> first element
len = 3
cap = 5
```

---

# 3. Creating Slices

## Slice Literal

```go
nums := []int{1, 2, 3}
```

---

## Using make

```go
nums := make([]int, 3, 5)
```

Format:

```go
make([]Type, length, capacity)
```

Result:

* length = 3
* capacity = 5
* values = [0 0 0]

---

## Length vs Capacity

### Length

How many elements are currently inside the slice.

### Capacity

How many elements fit before Go allocates a new array.

---

Example:

```go
nums := make([]int, 3, 5)

fmt.Println(len(nums))
fmt.Println(cap(nums))
```

Output:

```text
3
5
```

---

# 4. Append

## Adding Elements

```go
nums := []int{1, 2, 3}

nums = append(nums, 4)
```

Output:

```text
[1 2 3 4]
```

---

## Important Rule

Always store append result:

```go
nums = append(nums, 5)
```

NOT:

```go
append(nums, 5)
```

Reason:

* append may allocate a new array
* append returns the new slice

---

# 5. Slice Internals and Shared Memory

## Slice Assignment Does Not Copy Data

Example:

```go
a := []int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[100 2 3]
[100 2 3]
```

Why?

Because both slices share the same underlying array.

---

## Real Backend Danger

```go
cachedUsers := users
```

Changing `cachedUsers` may accidentally modify `users`.

This creates difficult bugs.

---

# 6. Copying Slices Properly

## Using copy

```go
a := []int{1, 2, 3}

b := make([]int, len(a))

copy(b, a)

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

Now both slices are independent.

---

# 7. Slice Expressions

General syntax:

```go
slice[start:end]
```

Rules:

* start inclusive
* end exclusive

---

## Examples

### First 3 Elements

```go
nums[:3]
```

---

### Middle Elements

```go
nums[1:4]
```

---

### Last 2 Elements

```go
nums[len(nums)-2:]
```

Go does not support negative indexes.

Invalid:

```go
nums[-2:]
```

---

# 8. Sub-Slices Share Memory

Example:

```go
nums := []int{1, 2, 3}

a := nums[:2]

a[0] = 100

fmt.Println(nums)
fmt.Println(a)
```

Output:

```text
[100 2 3]
[100 2]
```

Sub-slices share the same underlying array.

---

# 9. Range Loops

## Basic Syntax

```go
for index, value := range nums {
	fmt.Println(index, value)
}
```

---

## Ignore Index

```go
for _, value := range nums {
	fmt.Println(value)
}
```

---

## Index Only

```go
for i := range nums {
	fmt.Println(i)
}
```

---

# 10. Range Loop Pitfall

## Values Are Copies

Example:

```go
nums := []int{1, 2, 3}

for _, v := range nums {
	v = v * 2
}

fmt.Println(nums)
```

Output:

```text
[1 2 3]
```

Why?

Because `v` is a copy.

---

## Correct Way to Modify Slice

```go
for i := range nums {
	nums[i] = nums[i] * 2
}
```

Output:

```text
[2 4 6]
```

---

# 11. Structs Inside Slices

Example:

```go
type Task struct {
	ID   int
	Name string
	Done bool
}
```

---

## Slice of Structs

```go
tasks := []Task{
	{ID: 1, Name: "Task 1", Done: false},
	{ID: 2, Name: "Task 2", Done: true},
}
```

---

## Modifying Structs

Correct:

```go
tasks[0].Done = true
```

---

## Common Mistake

Wrong:

```go
for _, task := range tasks {
	task.Done = true
}
```

Reason:

* `task` is a copy

Correct:

```go
for i := range tasks {
	tasks[i].Done = true
}
```

---

# 12. Maps

## What Is a Map?

A map stores key-value pairs.

Example:

```go
ages := map[string]int{
	"Alice": 25,
	"Bob":   30,
}
```

---

# 13. Accessing Values

```go
fmt.Println(ages["Alice"])
```

---

# 14. Adding and Updating

```go
ages["Charlie"] = 40
```

---

# 15. Deleting Entries

```go
delete(ages, "Bob")
```

---

# 16. Zero Values in Maps

Example:

```go
scores := map[string]int{}

fmt.Println(scores["unknown"])
```

Output:

```text
0
```

Missing keys return zero value.

---

# 17. Comma-OK Idiom

Used to check if key exists.

Example:

```go
value, exists := scores["Bob"]
```

---

## Example

```go
scores := map[string]int{
	"Alice": 10,
}

value, exists := scores["Bob"]

fmt.Println(value)
fmt.Println(exists)
```

Output:

```text
0
false
```

---

# 18. Nil Maps

This causes panic:

```go
var users map[string]int

users["alice"] = 1
```

Reason:

* memory was never allocated

Correct:

```go
users := make(map[string]int)
```

OR:

```go
users := map[string]int{}
```

---

# 19. Word Frequency Example

```go
package main

import (
	"fmt"
	"strings"
)

func WordFrequency(s string) map[string]int {
	frequency := make(map[string]int)

	words := strings.Fields(s)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	result := WordFrequency("go is fun and go is fast")

	fmt.Println(result)
}
```

Possible Output:

```text
map[and:1 fast:1 fun:1 go:2 is:2]
```

---

# 20. In-Memory Task Tracker Example

```go
package main

import "fmt"

type Task struct {
	ID   int
	Name string
	Done bool
}

func main() {
	tasks := []Task{
		{ID: 1, Name: "Testing 1", Done: false},
		{ID: 2, Name: "Testing 2", Done: false},
		{ID: 3, Name: "Testing 3", Done: false},
	}

	tasks[1].Done = true

	for _, task := range tasks {
		fmt.Println(task)
	}
}
```

---

# 21. Common Backend Uses

Collections are everywhere in backend systems.

## Slices

Used for:

* API responses
* database results
* queues
* batching
* pagination

## Maps

Used for:

* caches
* lookup tables
* counters
* aggregations
* rate limiting

---

# 22. Common Pitfalls Summary

## Slice Assignment Shares Memory

```go
b := a
```

Does not copy underlying data.

---

## Range Values Are Copies

```go
for _, v := range nums
```

`v` is a copy.

---

## Sub-Slices Share Memory

```go
a := nums[:2]
```

Still uses same underlying array.

---

## Nil Maps Panic on Writes

Always initialize maps before writing.

---

# 23. Key Takeaways

After this chapter, you should understand:

* arrays are fixed-size
* slices are dynamic views over arrays
* slices share underlying memory
* append may allocate new arrays
* copy creates independent slices
* range values are copies
* maps return zero values for missing keys
* comma-ok idiom checks existence
* nil maps panic on writes

These concepts are foundational for Go backend development.

Mastering slices and maps will make later topics much easier:

* concurrency
* APIs
* databases
* caching
* middleware
* worker pools
* data processing
