package main // Package declaration

import "fmt" // Import package `fmt`

/*
This function calculates the square of an ocean
*/
func ocean(length int, width int) (int, string) { // Function which takes two int and returns two values
	/*
	   Returns the square of an ocean and a message.

	       Parameters
	       ----------
	       length: int
	           Length of the ocean.
	       width: int
	           Width of the ocean.

	       Returns
	       -------
	       int
	           Square of the ocean.
	       string
	           Message of the square of the ocean.

	*/
	square := length * width // Multiply length and width and store in square
	switch {
	case square <= 10: // Check if square is less than or equal to 10
		return square, "To staw" // Return square and message
	case square <= 20: // Else body
		return square, "To więcej niż staw"
	default:
		return square, "To ocean" // Return square and message
	}
}

func main() { // main function
	var length int                                                    // Initialize length variable
	var width int                                                     // Initialize width variable
	fmt.Println("length: ")                                           // Print message on console
	fmt.Scanln(&length)                                               // Take input from user
	fmt.Println("width: ")                                            // Print message on console
	fmt.Scanln(&width)                                                // Take input from user
	square, message := ocean(length, width)                           // Function call
	fmt.Println("Your water area has a square of:", square, "meters") // Print message on console
	fmt.Println(message)                                              // Print message on console
}
