package game

import "fmt"

// Prompter gets a slice of anything and prompts the user for a choice

func Prompt[T fmt.Stringer](choices []T) int {
	for i, choice := range choices {
		fmt.Printf("%d) %v\n", i+1, choice)
	}

	fmt.Printf("Enter your choice 1-%d\n.", len(choices))

	var input int
	_, scanErr := fmt.Scanln(&input)
	if scanErr != nil {
		fmt.Println("Couldn't read input, try again...")
		return Prompt(choices)
	}
	if input > len(choices) || input < 1 {
		fmt.Println("Invalid card - try again")
		return Prompt(choices)
	}

	return input - 1
}
