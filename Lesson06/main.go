package main

import (
	"errors"
	"fmt"
)

/*
========================
1. Simple Function
========================
*/

// Simple function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, Golang!")
}

/*
========================
2. Function with Parameters
========================
*/

// Function that takes parameters but returns no value
func greetUser(name string, age int) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

/*
========================
3. Function with Return Value
========================
*/

// Function that returns a single value
func add(a int, b int) int {
	return a + b
}

/*
========================
4. Function with Multiple Return Values
========================
*/

// Function that returns multiple values (result + success flag)
func divide(a int, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

/*
========================
5. Named Return Values
========================
*/

// Function with named return value
func rectangleArea(width int, height int) (area int) {
	area = width * height
	return // returns "area" automatically
}

/*
========================
6. Function with Error Return
========================
*/

// Function that returns a value and an error
func withdraw(balance float64, amount float64) (float64, error) {
	if amount > balance {
		return balance, errors.New("insufficient funds")
	}
	return balance - amount, nil
}

/*
========================
7. Variadic Function
========================
*/

// Variadic function that accepts unlimited integers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

/*
========================
8. Anonymous Function
========================
*/

// Anonymous function assigned to a variable
var multiply = func(a int, b int) int {
	return a * b
}

/*
========================
9. Function as Parameter
========================
*/

// Function that accepts another function as a parameter
func operate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

/*
========================
10. Function Returning a Function
========================
*/

// Function that returns another function
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

/*
========================
11. Struct and Method
========================
*/

// Struct definition
type User struct {
	Name string
	Age  int
}

// Method attached to User struct
func (u User) greet() {
	fmt.Println("Hello,", u.Name)
}

/*
========================
12. Pointer Receiver Method
========================
*/

// Method with pointer receiver (can modify struct data)
func (u *User) birthday() {
	u.Age++
}

/*
========================
13. Recursive Function
========================
*/

// Recursive function example (factorial)
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*
========================
14. Defer Example
========================
*/

// Function demonstrating defer keyword
func process() {
	defer fmt.Println("This runs at the end")
	fmt.Println("Processing...")
}

/*
========================
15. Closure
========================
*/

// Closure function that remembers state
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

/*
========================
Main Function
========================
*/

func main() {

	// Basic function
	sayHello()

	// Function with parameters
	greetUser("Kiana", 22)

	// Function with return value
	fmt.Println("Add:", add(3, 5))

	// Multiple return values
	result, ok := divide(10, 2)
	if ok {
		fmt.Println("Divide result:", result)
	}

	// Named return value
	fmt.Println("Area:", rectangleArea(4, 5))

	// Error handling function
	newBalance, err := withdraw(100, 30)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Balance:", newBalance)
	}

	// Variadic function
	fmt.Println("Sum:", sum(1, 2, 3, 4))

	// Anonymous function
	fmt.Println("Multiply:", multiply(4, 6))

	// Function as parameter
	fmt.Println("Operate:", operate(5, 3, add))

	// Function returning function
	double := multiplier(2)
	fmt.Println("Double:", double(10))

	// Struct methods
	user := User{Name: "Alex", Age: 25}
	user.greet()
	user.birthday()
	fmt.Println("New Age:", user.Age)

	// Recursive function
	fmt.Println("Factorial:", factorial(5))

	// Defer
	process()

	// Closure
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
}
