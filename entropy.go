package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func calculatePasswordEntropy(password string) float64 {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	symbols := "!@#$%^&*()_+-=[]{}|;:,.<>?"

	usedLowercase := strings.ContainsAny(password, lowercase)
	usedUppercase := strings.ContainsAny(password, uppercase)
	usedDigits := strings.ContainsAny(password, digits)
	usedSymbols := strings.ContainsAny(password, symbols)

	charsetSize := 0
	if usedLowercase {
		charsetSize += 26
	}
	if usedUppercase {
		charsetSize += 26
	}
	if usedDigits {
		charsetSize += 10
	}
	if usedSymbols {
		charsetSize += 32
	}

	if charsetSize == 0 {
		return 0
	}
	return math.Log2(float64(charsetSize)) * float64(len(password))
}

func printUsage() {
	fmt.Println("Usage: go run entropy.go <password>")
	fmt.Println("Example: go run entropy.go MyPassword123!")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	password := os.Args[1]
	entropy := calculatePasswordEntropy(password)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Entropy: %.2f bits\n", entropy)
}
