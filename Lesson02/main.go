// Calling Package
package main

//Import more than libraries
import (
	"fmt"

	//making Alias
	foo "net/http"
)

// Function
func main() {

	//Printing
	fmt.Println("Hello World")

	//Get response
	// := [It shows if the variables does not exist make them]
	// We should use the name of Alias of the library
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1/")

	//if Condition => Error Handling
	if err != nil {

		//Print Error
		fmt.Println("Error: ", err)

		//Return & Show value
		return
	}
	//defer : It should execute anyway
	defer resp.Body.Close()

	//Printing the Response from Get
	fmt.Println("HTPP Response Status: ", resp.Status)

}
