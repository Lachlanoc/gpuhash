package main

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// generateCombinations generates all possible combinations of characters
func generateCombinations(charset string, length int, prefix string, target string, knownPrefix string) (string, int64) {
	var count int64 = 0
	startTime := time.Now()

	// Create a function to recursively generate combinations
	var generate func(current string, depth int) string
	generate = func(current string, depth int) string {
		// Base case: if we've reached the desired length
		if depth == length {
			count++

			// Generate the potential flag
			potentialFlag := knownPrefix + current

			// Calculate SHA1 hash
			hash := sha1.Sum([]byte(potentialFlag))
			hashStr := fmt.Sprintf("%x", hash)

			// Check if hash matches target
			if hashStr == target {
				return potentialFlag
			}

			// Print progress every million attempts
			if count%1000000 == 0 {
				now := time.Now()
				elapsed := now.Sub(startTime).Seconds()
				rate := float64(count) / elapsed
				fmt.Printf("Tried %d combinations... (%.2f attempts/second)\n", count, rate)
			}

			return ""
		}

		// Try each character in the charset
		for i := 0; i < len(charset); i++ {
			result := generate(current+string(charset[i]), depth+1)
			if result != "" {
				return result
			}
		}

		return ""
	}

	// Start generation with empty string
	result := generate("", 0)

	return result, count
}

func main() {
	fmt.Println("SHA1 Hash Brute Force Tool")
	fmt.Println("==========================")

	// Hash options
	fmt.Println("\nSelect the hash type:")
	fmt.Println("1. SHA1")
	fmt.Println("2. SHA2")

	// Get target hash
	fmt.Print("Enter the SHA1 hash: ")
	var targetHash string
	fmt.Scanln(&targetHash)
	targetHash = strings.ToLower(targetHash)

	// Get known prefix
	fmt.Print("Enter the known prefix: ")
	var knownPrefix string
	fmt.Scanln(&knownPrefix)

	// Character set options
	fmt.Println("\nSelect character set for brute force:")
	fmt.Println("1. Numbers only (0-9)")
	fmt.Println("2. Lowercase letters only (a-z)")
	fmt.Println("3. Uppercase letters only (A-Z)")
	fmt.Println("4. Numbers and lowercase letters")
	fmt.Println("5. Numbers and uppercase letters")
	fmt.Println("6. All letters (a-z, A-Z)")
	fmt.Println("7. All alphanumeric characters (0-9, a-z, A-Z)")
	fmt.Println("8. Numbers and lowercase letters only (0-9, a-z)")

	// Get character set choice
	fmt.Print("Enter your choice (1-8): ")
	var choice string
	fmt.Scanln(&choice)

	// Define character sets
	digits := "0123456789"
	lowercaseLetters := "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var charset string

	// Set charset based on user choice
	switch choice {
	case "1":
		charset = digits
	case "2":
		charset = lowercaseLetters
	case "3":
		charset = uppercaseLetters
	case "4":
		charset = digits + lowercaseLetters
	case "5":
		charset = digits + uppercaseLetters
	case "6":
		charset = lowercaseLetters + uppercaseLetters
	case "7":
		charset = digits + lowercaseLetters + uppercaseLetters
	case "8":
		charset = digits + lowercaseLetters
	default:
		fmt.Println("Invalid choice. Using all alphanumeric characters.")
		charset = digits + lowercaseLetters + uppercaseLetters
	}

	// Get maximum length
	fmt.Print("Enter maximum length of characters to try (default: 8): ")
	var maxLengthStr string
	fmt.Scanln(&maxLengthStr)

	maxLength := 8
	if maxLengthStr != "" {
		var err error
		maxLength, err = strconv.Atoi(maxLengthStr)
		if err != nil {
			fmt.Println("Invalid input. Using default value of 8 characters.")
			maxLength = 8
		}
	}

	fmt.Printf("Starting brute force with prefix: %s\n", knownPrefix)
	fmt.Printf("Character set size: %d (%s)\n", len(charset), charset)
	fmt.Printf("Trying combinations up to %d characters\n", maxLength)

	startTime := time.Now()
	var totalAttempts int64 = 0

	// Try different lengths
	for length := 1; length <= maxLength; length++ {
		fmt.Printf("Trying %d character combinations...\n", length)

		result, attempts := generateCombinations(charset, length, "", targetHash, knownPrefix)
		totalAttempts += attempts

		if result != "" {
			elapsed := time.Since(startTime).Seconds()
			fmt.Printf("Flag found: %s\n", result)
			fmt.Printf("Completed in %.2f seconds after %d attempts\n", elapsed, totalAttempts)
			os.Exit(0)
		}
	}

	elapsed := time.Since(startTime).Seconds()
	fmt.Printf("Search completed in %.2f seconds after %d attempts\n", elapsed, totalAttempts)
	fmt.Println("Flag not found. Try increasing max_length or check your inputs.")
}
