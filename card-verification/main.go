package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// hasConsecutiveRepeatedDigits checks if any digit repeats n or more times consecutively
func hasConsecutiveRepeatedDigits(cardNumber string) bool {
	// Step 1: Remove hyphens
	reRemoveHyphen := regexp.MustCompile("-")
	processedString := reRemoveHyphen.ReplaceAllString(cardNumber, "")

	// Step 2: Check for any digit repeating 4 or more times consecutively
	n := 4 // Number of consecutive strings
	for i := 0; i <= len(processedString)-n; i++ {
		// Check if the next n characters are the same as the current character
		if strings.Count(processedString[i:i+n], string(processedString[i])) == n {
			return false
		}
	}
	return true
}

func validateCreditCard(cardNumber string) string {
	// Regex pattern for validating the credit card number

	// 1. Starts with 4, 5, or 6
	// 2. Contains exactly 16 digits or follows the format with 4 groups of 4 digits separated by hyphens
	pattern := `^[456]\d{3}(-?\d{4}){3}$`

	// Compile the regex
	re := regexp.MustCompile(pattern)

	// Check if the card number matches the pattern
	if re.MatchString(cardNumber) && hasConsecutiveRepeatedDigits(cardNumber){
		return "Valid"
	}
	return "Invalid"
}

func main() {
	var cardNumbers []string
	// Input number of credit card numbers
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}
	
	// Read each credit card number and validate it
	for range n {
		if scanner.Scan() {
			cardNumber := scanner.Text()
			cardNumbers = append(cardNumbers, cardNumber)
		}
	}

	for _, card := range cardNumbers {
		fmt.Println(validateCreditCard(card))
	}
}