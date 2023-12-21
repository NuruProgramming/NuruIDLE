---
title: Arrays
titleTemplate: Arrays in Nuru are versatile data structures that can hold multiple items
---

# Arrays in Nuru

Arrays in Nuru are versatile data structures that can hold multiple items, including different types such as numbers, strings, booleans, functions, and null values. This page covers various aspects of arrays, including how to create, manipulate, and iterate over them using Nuru's built-in keywords and methods.

## Creating Arrays

To create an array, use square brackets [] and separate items with commas:

```go
orodha = [1, "pili", kweli]
```

## Accessing and Modifying Array Elements

Arrays in Nuru are zero-indexed. To access an element, use the element's index in square brackets:

```go
namba = [10, 20, 30]
jina = namba[1]  // jina is 20
```

You can reassign an element in an array using its index:

```go
namba[1] = 25
```

## Concatenating Arrays

To concatenate two or more arrays, use the + operator:

```go
a = [1, 2, 3]
b = [4, 5, 6]
c = a + b
// c is now [1, 2, 3, 4, 5, 6]
```

## Checking for Array Membership

Use the `ktk` keyword to check if an item exists in an array:

```go
namba = [10, 20, 30]
andika(20 ktk namba)  // will print kweli
```

## Looping Over Arrays

You can use the kwa and ktk keywords to loop over array elements. To loop over just the values, use the following syntax:

```
namba = [1, 2, 3, 4, 5]

kwa thamani ktk namba {
    andika(thamani)
}
```

To loop over both index and value pairs, use this syntax:

```go
majina = ["Juma", "Asha", "Haruna"]

kwa idx, jina ktk majina {
    andika(idx, "-", jina)
}
```

## Array Methods

Arrays in Nuru have several built-in methods:

### idadi()

`idadi()` returns the length of an array:

```go
a = [1, 2, 3]
urefu = a.idadi()
andika(urefu)  // will print 3
```

### sukuma()

`sukuma()` adds one or more items to the end of an array:

```go
a = [1, 2, 3]
a.sukuma("s", "g")
andika(a)  // will print [1, 2, 3, "s", "g"]
```

### yamwisho()

`yamwisho()` returns the last item in an array, or tupu if the array is empty:

```go
a = [1, 2, 3]
mwisho = a.yamwisho()
andika(mwisho)  // will print 3

b = []
mwisho = b.yamwisho()
andika(mwisho)  // will print tupu
```

With this information, you can now effectively work with arrays in Nuru, making it easy to manipulate collections of data in your programs.
