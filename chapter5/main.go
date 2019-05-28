package main

//Control Structures

import "fmt"

func main() {

	// chapter 1 Write a program that
	// displays numbers from 1 to 100,
	// which are divisible by 3

	/*	for i := 1; i <= 100; i++ {
				if i%3 == 0 {
					fmt.Println(i)
				}
			}
		}
	*/

	// chapter 2 Write a program that displays numbers
	//from 1 to 100. But for multiples of three,
	//you need to output “Fizz” instead of a number,
	//for multiples of five - “Buzz”, and for multiples
	//of three, and five - “FizzBuzz”.

	for i := 1; i <= 100; i++ {
		switch {

		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}
	}
}

//Examples from the book
/*	for i := 1; i <= 10; i++ {
	/*		if i % 2 == 0 {
						fmt.Println(i,"even")
						} else {
							fmt.Println(i,"odd")
						}
					}
				}

			 if i == 0 {
			    fmt.Println("Zero")
			} else if i == 1 {
			    fmt.Println("One")
			} else if i == 2 {
			    fmt.Println("Two")
			} else if i == 3 {
			    fmt.Println("Three")
			} else if i == 4 {
			    fmt.Println("Four")
			} else if i == 5 {
			    fmt.Println("Five")
			}
			}
			}
*/
/*		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}
}
*/
/*	i := 10

	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}
*/
