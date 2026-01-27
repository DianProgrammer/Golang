package main

// Data-Oriented vs Object-Oriented (Conceptual Explanation)
//
// Data-Oriented design focuses on DATA first, not objects.
// The primary concern is how data is structured, stored, and processed,
// rather than how many objects or class hierarchies exist.
//
// In a data-oriented approach:
// - Data structures (structs / records) come first
// - Functions operate on data
// - Behavior is separated from data
// - Simplicity and predictability are prioritized
//
// Object-Oriented design focuses on OBJECTS and their behavior.
// Data and behavior are tightly coupled inside classes.
//
// In an object-oriented approach:
// - Objects are created from classes
// - Methods belong to objects
// - Inheritance and polymorphism are commonly used
// - The focus is on "what this object does"
//
// Key idea:
// Data-Oriented design emphasizes value assignment and data flow
// over object instantiation and class-based identity.
//
// Why Go is data-oriented:
// - Structs are simple data containers
// - No class inheritance
// - Value semantics by default
// - Functions and methods are kept explicit and minimal
//
// Summary:
// Data-Oriented design answers:
// "What does the data look like, and how do we process it?"
//
// Object-Oriented design answers:
// "What is this object, and what can it do?"

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//Pascal Case
	//Structs, interface, enums

	//Snake Case
	//Exp. user_id, first_name, http_request

	//UPPERCASE
	//use case is Constants

	//mixCase
	//Exp. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5
	var employeeID = 100
	fmt.Println("EmployeeID:", employeeID)

}

// In Go, struct plays a role similar to a class in C# or Java,
// but Go is NOT a classical object-oriented language.
//
// Similarities:
// - A struct groups related data (fields) together
// - A struct can have methods attached to it
// - You can create instances of a struct
//
// Key Differences from C#/Java classes:
//
// 1. No inheritance
// Go does NOT support class inheritance (extends / super).
// Instead, it uses composition and interfaces.
//
// 2. No access modifiers (public/private/protected)
// Visibility is controlled by naming:
// - Capitalized names are exported (public)
// - Lowercase names are unexported (private)
//
// 3. No real constructors
// Go does not have constructors.
// Factory functions (e.g. NewEmployee) are commonly used instead.
//
// 4. Methods are defined outside the struct
// Methods are declared separately and attached to the struct via receivers.
//
// 5. Structs are value types
// Assigning a struct copies its data by default.
// This is different from classes in C#/Java, which are reference types.
//
// Summary:
// A Go struct can be thought of as a "class-like" structure,
// but Go favors simplicity, composition, and explicit behavior
// over traditional object-oriented patterns.
