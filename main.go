// Copyright © 2021 Rikesh Karmacharya

package main

import (
	"fmt"
	"time"

	placeholder "github.com/rikeshkarma/Retro-LED-Clock/components"
)

func main() {
	
	//screen.Clear()
	fmt.Print("\033[2J")

	for {
		//screen.MoveTopLeft()
		fmt.Print("\033[H")

		// get the current hour, minute and second
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()

		// extract the digits: 17 becomes, 1 and 7 respectively
		clock := [...]placeholder.Placeholder{
			placeholder.Digits[hour/10], placeholder.Digits[hour%10],
			placeholder.Colon,
			placeholder.Digits[min/10], placeholder.Digits[min%10],
			placeholder.Colon,
			placeholder.Digits[sec/10], placeholder.Digits[sec%10],
		}

		// Explanation: clock[0]
		// + Each element of clock has the same length.
		// + So: Getting the length of only one element is OK.
		// + This could be: "zero" or "one" and so on... Instead of: digits[0]
		//
		// The range clause below is ~equal to the following code:
		//   line := 0; line < len(clock[0]); line++
		for line := range clock[0] {
			// Print a line for each placeholder in clock
			for index, digit := range clock {
				// Colon blink on every two seconds.
				// + On each sec divisible by two, prints an empty line
				// + Otherwise: prints the current pixel
				next := clock[index][line]
				if digit == placeholder.Colon && sec%2 == 0 {
					next = "   "
				}
				// Print the next line and,
				// give it enough space for the next placeholder
				fmt.Print(next, "  ")
			}

			// After each line of a placeholder, print a newline
			fmt.Println()
		}

		// pause for 1 second
		time.Sleep(time.Second)
	}
}