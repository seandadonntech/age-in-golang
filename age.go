package main

import "fmt"

func main() {
	var age int
	fmt.Println("enter your age:")
	fmt.Scan(&age)
	fmt.Println("you entered", age, "as you're age")

	if age >= 21 {
		fmt.Println("nice")
	} else {
		fmt.Println("woah you cant drink its not legal")
	}
       time.Sleep(8 * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")
}

