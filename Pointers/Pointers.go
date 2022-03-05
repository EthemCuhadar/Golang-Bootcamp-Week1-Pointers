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
	
	fmt.Println("----- Values -----")
	fmt.Println("foo1 value is: ", foo1)
	fmt.Println("foo2 value is: ", foo2)
	fmt.Println("foo3 value is: ", foo3)
	fmt.Println("foo4 value is: ", foo4)
	
	fmt.Println("----- Equalities -----")
	fmt.Println(foo1 == foo2)
	fmt.Println(foo3 == foo4)
	
	fmt.Println("----- Types -----")
	fmt.Println("The type of foo1 variable is: ", reflect.TypeOf(foo1))
	fmt.Println("The type of foo2 variable is: ", reflect.TypeOf(foo2))
	fmt.Println("The type of foo3 variable is: ", reflect.TypeOf(foo3))
	fmt.Println("The type of foo4 variable is: ", reflect.TypeOf(foo4))
	
	// ADRESSES OF VARIABLES
	
	fmt.Println("----- Adresses -----")
	fmt.Println("Adress of foo1 is: ", &foo1)
	fmt.Println("Adress of foo2 is: ", &foo2)
	fmt.Println("Adress of foo3 is: ", &foo3)
	fmt.Println("Adress of foo4 is: ", &foo4)
	
	// A variable can be used to refer adress of another variable
	var strVal string
	strVal = "cat"
	
	adress_strVal := &strVal
	
	fmt.Println("Adress of strVal is: ", adress_strVal)
	
	// *(star) concept
	var floatVal float32
	floatVal = 15.99
	
	adress_floatVal := &floatVal
	
	fmt.Println(adress_floatVal)
	fmt.Println()
	
	x1 := 10
	x2 := &x1
	// x2 refers to address(memory location) of x1
	fmt.Println(x1)
	fmt.Println(x2)
	
	// printing the value that is pointed out by x1
	fmt.Println(*x2)
	// printing the memory location again
	fmt.Println(&*x2)
	// printing the value of that memory location which refers to the value
	fmt.Println(*&*x2)
	
	fmt.Printf("%T, %v\n", x2, x2)
	
	// Complex examples
	y1 := 10
	y2 := y1
	fmt.Println(y1, y2)
	y1 = 5
	fmt.Println(y1, y2)
	
	y3 := 15
	y4 := &y3
	fmt.Println(y3, y4)
	fmt.Println(y3, *y4)
	
	y3 = 9
	fmt.Println(y3, *y4)

}
