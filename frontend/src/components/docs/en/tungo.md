---
title: Strings
titleTemplate: Strings are a sequence of characters that can represent text in the Nuru programming language
---

# Strings in Nuru

Strings are a sequence of characters that can represent text in the Nuru programming language. This page covers the basics of strings, their manipulation, and some built-in methods.

## Basic Syntax

Strings can be enclosed in either single quotes '' or double quotes "":

```go
andika("mambo") // mambo

fanya a = 'niaje'

andika("mambo", a) // mambo niaje
```

## Concatenating Strings

Strings can be concatenated using the + operator:

```go
fanya a = "habari" + " " + "yako"

andika(a) // habari yako

fanya b = "habari"

b += " yako"

// habari yako
```

You can also repeat a string n number of times using the \* operator:

```go
andika("mambo " * 4)

// mambo mambo mambo mambo

fanya a = "habari"

a *= 4

// habarihabarihabarihabari
```

## Looping over a String

You can loop through a string using the kwa keyword:

```go
fanya jina = "avicenna"

kwa i ktk jina {andika(i)}
```

Output

```go
a
v
i
c
e
n
n
a
```

And for key-value pairs:

```go
kwa i, v ktk jina {
	andika(i, "=>", v)
}
```

Output

```go
0 => a
1 => v
2 => i
3 => c
4 => e
5 => n
6 => n
7 => a
```

## Comparing Strings

You can compare two strings using the == operator:

```go
fanya a = "nuru"

andika(a == "nuru") // kweli

andika(a == "mambo") // sikweli
```

## String Methods

### idadi()

You can find the length of a string using the idadi method. It does not accept any parameters.

```go
fanya a = "mambo"
a.idadi() // 5
```

### herufikubwa()

This method converts a string to uppercase. It does not accept any parameters.

```go
fanya a = "nuru"
a.herufikubwa() // NURU
```

### herufindogo()

This method converts a string to lowercase. It does not accept any parameters.

```go
fanya a = "NURU"
a.herufindogo() // nuru
```

### gawa()

The gawa method splits a string into an array based on a specified delimiter. If no argument is provided, it will split the string according to whitespace.

Example without a parameter:

```go
fanya a = "nuru mambo habari"
fanya b = a.gawa()
andika(b) // ["nuru", "mambo", "habari"]
```

Example with a parameter:

```go
fanya a = "nuru,mambo,habari"
fanya b = a.gawa(",")
andika(b) // ["nuru", "mambo", "habari"]
```

By understanding strings and their manipulation in Nuru, you can effectively work with text data in your programs.
