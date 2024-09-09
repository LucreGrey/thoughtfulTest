package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Function to check if a string represents a valid positive integer
func isValidPositiveInt(input string) (int, error) {
	// Check if the string contains a decimal point
	if strings.Contains(input, ".") {
		return 0, errors.New("invalid input: floating-point numbers are not allowed")
	}

	// Try to convert the string to an integer
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("invalid input: must be a valid integer")
	}

	// Check for positive non-zero integers
	if value <= 0 {
		return 0, errors.New("invalid input: dimensions and mass must be positive non-zero integers")
	}

	return value, nil
}

// sort function to determine which stack the package should go into
func sort(widthStr, heightStr, lengthStr, massStr string) (string, error) {
	// Validate and convert each input from string to int
	width, err := isValidPositiveInt(widthStr)
	if err != nil {
		return "", err
	}

	height, err := isValidPositiveInt(heightStr)
	if err != nil {
		return "", err
	}

	length, err := isValidPositiveInt(lengthStr)
	if err != nil {
		return "", err
	}

	mass, err := isValidPositiveInt(massStr)
	if err != nil {
		return "", err
	}

	// Calculate the volume of the package
	volume := width * height * length

	// Check for both bulky and heavy cases first
	if (volume >= 1000000 || width >= 150 || height >= 150 || length >= 150) && mass >= 20 {
		return "REJECTED", nil
	}

	// Check if it's bulky or heavy
	if volume >= 1000000 || width >= 150 || height >= 150 || length >= 150 || mass >= 20 {
		return "SPECIAL", nil
	}

	// Otherwise, it's a standard package
	return "STANDARD", nil
}

func main() {
	// Test cases
	tests := []struct {
		width, height, length, mass string
	}{
		{"100", "100", "100", "10"},   // STANDARD
		{"200", "100", "100", "10"},   // SPECIAL (Bulky)
		{"100", "100", "100", "25"},   // SPECIAL (Heavy)
		{"200", "200", "200", "25"},   // REJECTED (Bulky and Heavy)
		{"-100", "100", "100", "10"},  // Invalid input (Negative width)
		{"100", "100", "100", "-25"},  // Invalid input (Negative mass)
		{"0", "100", "100", "25"},     // Invalid input (Zero width)
		{"100", "0", "100", "25"},     // Invalid input (Zero height)
		{"100.5", "100", "100", "25"}, // Invalid input (Floating-point number)
		{"abc", "100", "100", "25"},   // Invalid input (String input)
	}

	for _, test := range tests {
		result, err := sort(test.width, test.height, test.length, test.mass)
		if err != nil {
			fmt.Printf("Input: (%s, %s, %s, %s) -> Error: %s\n", test.width, test.height, test.length, test.mass, err)
		} else {
			fmt.Printf("Input: (%s, %s, %s, %s) -> Output: %s\n", test.width, test.height, test.length, test.mass, result)
		}
	}
}
