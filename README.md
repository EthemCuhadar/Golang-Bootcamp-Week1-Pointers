# Golang

![Golang Image](golang.png)

--------------------------------------------

# Pointers

--------------------------------------------------

## Declerations and assignments

Pointers are one of the important concepts in Go programming language. There are several pointers and several decleration types to define a variable.

```go
// First decleration
var foo int
foo = 22
```

```go
// Second decleration
var foo int = 22
```

```go
// Third decleration
var foo = 22
```

```go
// Fourth decleration (shord-hand form)
foo := 22
```

* These are all same variables. Their names, values and types are same. 

* In short-hand form decleration and assignment are applied in background.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {

    // DECLERATIONS 

    // first decleration
    var foo1 int
    foo1 = 32

    // second decleration
    var foo2 int = 32

    // third decleration
    var foo3 = 32

    // fourth decleration (short-hand form)
    foo4 := 32

    // printing the values
    fmt.Println("----- Values -----")
    fmt.Println("foo1 value is: ", foo1)
    fmt.Println("foo2 value is: ", foo2)
    fmt.Println("foo3 value is: ", foo3)
    fmt.Println("foo4 value is: ", foo4)

    // equality checks
    fmt.Println("----- Equalities -----")
    fmt.Println(foo1 == foo2)
    fmt.Println(foo3 == foo4)

    // types of the variables
    fmt.Println("----- Types -----")
    fmt.Println("The type of foo1 variable is: ", reflect.TypeOf(foo1))
    fmt.Println("The type of foo2 variable is: ", reflect.TypeOf(foo2))
    fmt.Println("The type of foo3 variable is: ", reflect.TypeOf(foo3))
    fmt.Println("The type of foo4 variable is: ", reflect.TypeOf(foo4))

}
```

```console
---- Values -----
foo1 value is:  32
foo2 value is:  32
foo3 value is:  32
foo4 value is:  32
----- Equalities -----
true
true
----- Types -----
The type of foo1 variable is:  int
The type of foo2 variable is:  int
The type of foo3 variable is:  int
The type of foo4 variable is:  int
```

* Different variables are defined above which are **foo1**, **foo2**, **foo3** and **foo4**. Their names different. However, values and types of the variables are same as it can be seen console above.

* Note that, in third and fourth declerations, the types of the variables were not defined. In Go programming language default integer type is **int**. However, if we want to define **int32**, it should be specified explicitly.

------------------------------------------------------

## Address of a variable

* Address of a variable refers to memory location of the variable.

* In Go programming language, adress of a variable can be defined with **address of** symbol (**&**) using before the varible name.

```go
    // ADRESSES OF VARIABLES

    fmt.Println("----- Adresses -----")
    fmt.Println("Address of foo1 is: ", &foo1)
    fmt.Println("Address of foo2 is: ", &foo2)
    fmt.Println("Address of foo3 is: ", &foo3)
    fmt.Println("Address of foo4 is: ", &foo4)
```

```console
----- Adresses -----
Address of foo1 is:  0xc000014098
Address of foo2 is:  0xc0000140b0
Address of foo3 is:  0xc0000140b8
Address of foo4 is:  0xc0000140c0
```

* A variable can be point out adress of another variable.

```go
var strVal string
strVal = "cat"

adress_strVal := &strVal

fmt.Println("Address of strVal is: ", address_strVal)
```

```console
Address of strVal is:  0xc000044230
```

----------------------------------------------------

## Asteriks (*) concept

* Asteriks (*) concept is another important concept in Go programming language. The star symbol can be a little bit confusing. There used in two ways.
1. The star can be used in variable type.

2. The star can be used in variable name.
* If asteriks used before variable name, which refers to the value of the pointer that pointed out. 

* On the other hand, if it is used just before the type (*int, *string), that refers to the pointer which can be pointed out stated type of data. The examples for these 2 situations can be seen below.

```go
x1 := 10
x2 := &x1
// x2 refers to address(memory location) of x1
fmt.Println(x1)
fmt.Println(x2)
```

```console
10
0xc000014100
```

* Now we will use asteriks just before x2. Hence, that will give the value that is pointed out by x1.

```go
fmt.Println(*x2)
```

```console
10
```

* The operation above is names as **derefering**.

* If we want to print out memory location of ***x2**, we just need to use another **address of**.

```go
fmt.Println(&*x2)
```

```console
0xc00001410
```

* As you can see above, it gave the same memory location again. Furthermore, lets print out the value in that memory location again.

```go
fmt.Println(*&*x2)
```

```console
10
```

* It can be clearly seen that, it is like a loop that gives the same memory locations and values.

* The second one is using the asteriks before the type. That refers to the pointer which can be pointed out defined type of values only.

```go
fmt.Printf("%T, %v\n", x2, x2)
```

```console
*int, 0xc000014100
```

* The type of the x2 is **asteriks integer** and memory location of the x2 is **0xc000014100**. Asteriks integer type means that the pointer x2 can be only pointed out integer values.

* Now lets look at the more complex examples with these topics.

```go
// Complex examples
y1 := 10
y2 := y1
fmt.Println(y1, y2)
y1 = 5
fmt.Println(y1, y2)
```

```console
10 10
5 10
```

* In this example **y1** and **y2** shares the same value which is 10. After that **y1** value is changed as **5**. However, **y2** value didn't change. This is just because of they are not sharing same addresses. They just share same values.

* Let's look at another example.

```go
y3 := 15
y4 := &y3
fmt.Println(y3, y4)
fmt.Println(y3, *y4)
	
y3 = 9
fmt.Println(y3, *y4)
```

```console
15 15
9 9
```

* This time **y4** pointed out memory location of **y3**. Which means any changes in **y3** affects **y4** as well. Because they share same address.
