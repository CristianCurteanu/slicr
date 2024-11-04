

# Golang's Slicr 

## Description

Golang Slicr (pronounced as `slicer`) is a collection of generic functions, that makes is easier to operate Golang's slices. You can read all the functions reference down below in this description.

In order to be able to use it, make sure, that you are importing for an Go version `>= 1.18`, as starting with this version, Go started to support generics, and as mentioned above, these are generic functions.

## Installation

In  order to install this package, you need to use `go get`:

```sh
go get github.com/CristianCurteanu/slicr
```

## Reference

Table of contents:
- Find
- Contains
- Find
- Contains
- Any
- All
- Filter
- ForEach
- Map
- Reduce
- Pop
- Prepend
- Shift
- Delete
- Extend
- ExtendAt
- ExtendCap
- Deduplicate
- Count
- Partition
- GroupBy
- Distinct/Unique
- Reverse
- Sort
- Zip
- Unzip
- Flatten
- Chunk
- Take
- Drop
- TakeWhile
- DropWhile
- Max
- Min
- Sum
- Product
- Average/Mean
- Copy/Clone
- Equal

### Find

Locate the first element that satisfies a predicate.

```go

fruits := []string{"apple", "orange", "banana"}

fruit, found := slicr.Find(fruits, func (f string) bool { return f == "orange" })
fmt.Printf("found: %s", fruit)
```

The generic type of `Find` will be `string`, as it will automatically detect by slice type reference. In this case it will return the `orange` element, and will return the `found` value as `true`

In case the predicate will run through each element and will not find the element according to the predicate, it will return zero value of the slice element type, and `false` to the found flag.

### Contains

Check if a slice contains a specific element. 

```go

fruits := []string{"apple", "orange", "banana"}

found := slicr.Contains(fruits, "orange")
fmt.Printf("found: %s", found)
```

Bear in mind that the generic type of the slice should be `comparable`, otherwise you need to consider `Any`, to use a predicate function

### Any

Determine if any elements satisfy a predicate

It works the same way as the return found flag of the `Find` function, but it only returns the `bool` value, that signals that there is a value that satisfies the predicate function.

```go

fruits := []string{"apple", "orange", "banana"}

found := slicr.Any(fruits, func (f string) bool { return f == "orange" })
fmt.Printf("found: %v", fruit)
```

### All

Check if all elements satisfy a predicate.

```go
fruits := []string{"apple", "orange", "banana"}

noEmptyStrings := slicr.All(fruits, func(f string) bool { return f != "" })
fmt.Printf("any empty string?: %v", !noEmptyStrings)
```

### Filter

Select elements from a slice that satisfy a given predicate function.

If the predicate returns `true`, the value will be added to the new slice.

```go

fruits := []string{"apple", "orange", "banana"}

fruits = slicr.Filter(fruits, func (f string) bool { return f == "orange" })
fmt.Printf("selected fruits: %s", fruits)
```

### ForEach

Execute a function for each element in a slice, typically for side effects.

```go

fruits := []string{"apple", "orange", "banana"}

fruits = slicr.ForEach(fruits, func (f string) { 
    fmt.Println(fruit)
})
```

This could be useful, when an existing predicate function could be applied for each element, and not expecting a result; same as range, but it's a bit cleaner.

### Map

Apply a function to each element of a slice, producing a new slice with the results.

```go

type Fruit struct {
    Name string
}

fruits := []string{"apple", "orange", "banana"}

fruitsEntities := slicr.Map(fruits, func (f string) Fruit { return Fruit{Name: f} })
fmt.Printf("mapped fruits: %+v", fruitsEntities)
```

The predicate acts as mapping function.

### Reduce

Aggregate elements of a slice into a single value using an accumulator function.

```go

fruits := []string{"apple", "orange", "banana"}

list = slicr.Reduce(fruits, func (res, f string) string { 
    return fmt.Sprintf("%s, %s", res, f)
})
fmt.Printf("reduced fruits: %s", list) // reduced fruits: apple, orange, banana
```

This will fold the initial string slice, into a single string, by joining all elements according to the reducer function.

### Pop

Pops the last element of the slice, and reduces the slice by last popped element

```go
fruits := []string{"apple", "orange", "banana"}

fruit, fruits := slicr.Pop(fruits)
fmt.Printf("popped: %q\n", fruit) // popped: "banana"
fmt.Printf("fruits: %+v", fruits) // fruits: [apple orange]
```

### Prepend

Adds a new element at the beginning of the slice

```go
fruits := []string{"apple", "orange", "banana"}

fruits = slicr.Prepend(fruits, "mango", "pineapple")
fmt.Printf("fruits: %+v", fruits)
```

### Shift

Pops an element at the beginning of the slice

```go
fruits := []string{"apple", "orange", "banana"}

fruit, fruits := slicr.Shift(fruits)
fmt.Printf("fruit: %+v\n", fruit) // fruit: apple
fmt.Printf("fruits: %+v", fruits) // fruits: [orange banana]
```

### Delete

Delete element at the given index

```go
fruits := []string{"apple", "orange", "banana"}

fruits = slicr.Delete(fruits, 1)

fmt.Printf("fruits: %+v", fruits) // fruits: [apple banana]
```

### Extend

Extends the size of a given slice.

```go
fruits := []string{"apple", "orange", "banana"}

fmt.Printf("fruits.size: %+v\n", len(fruits)) // fruits.size: 3
fruits = slicr.Extend(fruits, 1)
fmt.Printf("fruits.size: %+v", len(fruits)) // fruits.size: 4
```


### ExtendAt

Extends the size of slice by n elements at position i

```go
fruits := []string{"apple", "orange", "banana"}

fmt.Printf("fruits.size: %+v\n", len(fruits)) // fruits.size: 3

fruits = slicr.ExtendAt(fruits, 2, 1)
fmt.Printf("fruits.size: %+v\n", len(fruits)) // fruits.size: 5
fmt.Printf("fruits: %+v", fruits) // fruits: [apple   orange banana]
```

### ExtendCap

Extends the capacity of a given slice by n elements

```go
fruits := []string{"apple", "orange", "banana"}

fmt.Printf("fruits.cap: %+v\n", cap(fruits)) // fruits.cap: 3

fruits = slicr.ExtendCap(fruits, 5)
fmt.Printf("fruits.cap: %+v\n", cap(fruits)) // fruits.cap: 8
fmt.Printf("fruits: %+v", fruits) // fruits: [apple orange banana]
```

### Deduplicate

Removes all duplicate elements in a slice

```go
fruits := []string{"apple", "orange", "banana", "orange"}

fruits = Deduplicate(fruits)

fmt.Printf("fruits: %+v", fruits) // fruits: [apple banana orange]
```

### Count

Count the number of elements that satisfy a predicate

```go
fruits := []string{"apple", "orange", "banana", "orange"}

count := Count(fruits, func(f string) bool {
    return f == "orange"
})

fmt.Printf("oranges: %+v", count) // oranges: 2
```

### Partition

Split a slice into two slices based on a predicate.

```go
fruits := []string{"apple", "orange", "banana", "orange"}

oranges, fruits := Partition(fruits, func(f string) bool {
    return f == "orange"
})

fmt.Printf("oranges: %+v\n", oranges) // oranges: [orange orange]
fmt.Printf("fruits: %+v", fruits)     // fruits: [apple banana]
```

### GroupBy

Group elements of a slice based on a key selector function.

```go
fruits := []string{"apple", "orange", "banana", "orange"}

groupedFruits := GroupBy(fruits, func(f string) (string, string) {
    return string(f[len([]byte(f))-1]), f // This will group by last char of strings, which will be `e` and `a` groups
})

fmt.Printf("fruits: %+v", groupedFruits) // fruits: [[apple orange orange] [banana]]
```

### Chunk

Split a slice into multiple slices of a specified size.


```go
fruits := []string{"apple", "orange", "banana", "orange"}

groupedFruits := Chunk(fruits, 3)
fmt.Printf("fruits: %+v", groupedFruits) // fruits: [[apple orange banana] [orange]]
```

### Zip

Combine two slices into one slice of tuples.

### Unzip

Separate a slice of tuples into two slices.

### Flatten

Flatten a slice of slices into a single slice.

### Take

Take the first N elements from a slice.

### Drop

Skip the first N elements of a slice.

### TakeWhile

Take elements from the slice while a predicate is true.

### DropWhile

Skip elements in the slice while a predicate is true.

### Max

Find the maximum element in a slice (requires ordering).

### Min

Find the minimum element in a slice (requires ordering).

### Sum

Calculate the sum of elements (numeric slices).

### Product

Calculate the product of elements (numeric slices).

### Average/Mean

Compute the average value of elements (numeric slices).

### Copy/Clone

Create a shallow copy of a slice.

### Equal

Check if two slices are equal in content and order.

